package service

import (
	"context"
	v1 "service/api/blackpig/v2"
	"service/internal/biz"
)

// convertUer: 需要设置一个转置函数将biz的user转化为v1的user,方便返回给前端
func convertUer(user *biz.User) *v1.User {
	return &v1.User{
		Id:          user.Id,
		Username:    user.Username,
		Phone:       user.Phone,
		Openid:      user.Openid,
		CarPictures: user.CarPictures,
		Gender:      user.Gender,
	}
}

type UserService struct {
	v1.UnimplementedUserServiceServer
	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}

// Login （service层）这里是实现api中所申明的方法，具体实现api中的login方法，api只需要通过rpc调用方法名而不会进入内部执行
func (u *UserService) Login(ctx context.Context, request *v1.LoginRequest) (*v1.UserReply, error) {
	code := request.Code
	user, err := u.uc.Login(ctx, code)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{User: convertUer(user)}, nil
}

// ListUser （获取用户列表方法）
func (u *UserService) ListUser(ctx context.Context, request *v1.ListRequest) (*v1.ListReply, error) {
	// 获取请求中的关键词当前页，页面尺寸，关键词
	current, size, keyword := request.Current, request.Size, request.KeyWord
	users, err := u.uc.UserList(ctx, current, size, keyword)
	if err != nil {
		return nil, err
	}
	reply := make([]*v1.User, len(users))
	for i, user := range users {
		reply[i] = convertUer(&user)
	}
	return &v1.ListReply{User: reply}, nil
}

// FindUserByPhone 通过用户手机号进行查找用户
func (u *UserService) FindUserByPhone(ctx context.Context, request *v1.UserByPhoneRequest) (*v1.UserReply, error) {
	phone := request.Phone
	user, err := u.uc.FindUserByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{User: convertUer(user)}, nil
}

// Register 用户注册接口方法实现
func (u *UserService) Register(ctx context.Context, request *v1.RegisterRequest) (*v1.RegisterReply, error) {
	user := &biz.User{
		Id:     request.User.Id,
		Phone:  request.User.Phone,
		Gender: request.User.Gender,
		Openid: request.User.Openid,
	}
	err := u.uc.Register(ctx, user)
	if err != nil {
		return &v1.RegisterReply{State: "注册失败"}, err
	}
	return &v1.RegisterReply{State: "注册成功"}, nil
}

func (u *UserService) UpdateUser(ctx context.Context, request *v1.RegisterRequest) (*v1.UserReply, error) {
	user := &biz.User{
		Phone:       request.User.Phone,
		Gender:      request.User.Gender,
		Openid:      request.User.Openid,
		CarPictures: request.User.CarPictures,
		Username:    request.User.Username,
		Description: request.User.Description,
	}
	res, err := u.uc.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{User: convertUer(res)}, nil
}

// DeleteUser 删除用户
func (u *UserService) DeleteUser(ctx context.Context, request *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	id := request.Id
	err := u.uc.RemoveUser(ctx, id)
	if err != nil {
		return &v1.DeleteUserReply{State: "删除用户失败"}, err
	}
	return &v1.DeleteUserReply{State: "删除用户成功"}, nil
}
