package data

import (
	"context"
	"service/internal/biz"
	"service/internal/ent"
	"service/internal/ent/user"
)

type userRepo struct {
	data *Data //ORM操作方法，如ENT
}

func convertUser(user *ent.User) *biz.User {
	return &biz.User{
		Username:    user.Username,
		Phone:       user.Phone,
		Id:          user.ID,
		Gender:      user.Gender,
		CarPictures: user.CarPictures,
		Openid:      user.Openid,
	}
}

// FindAll (data层) 搜索所有用户，可设置页码current，尺寸size，关键词keyword
func (u userRepo) FindAll(current, size int32, keyWord string, ctx context.Context) ([]biz.User, error) {
	all, err := u.data.db.User.Query().
		Where(user.UsernameContains(keyWord)).
		Offset(int((current - 1) * size)).
		Limit(int(size)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]biz.User, len(all))
	for i, s := range all {
		res[i] = *convertUser(s)
	}
	return res, nil
}

func (u userRepo) FindByPhone(ctx context.Context, phone string) (*biz.User, error) {
	first, err := u.data.db.User.Query().Where(user.Phone(phone)).First(ctx)
	if err != nil {
		return nil, err
	}
	return convertUser(first), nil
}

func (u userRepo) FindById(ctx context.Context, i int32) (*biz.User, error) {
	first, err := u.data.db.User.Query().Where(user.ID(i)).First(ctx)
	if err != nil {
		return nil, err
	}
	return convertUser(first), nil
}

func (u userRepo) FindByUsername(ctx context.Context, s string) (*[]biz.User, error) {
	users, err := u.data.db.User.Query().Where(user.UsernameContains(s)).All(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]biz.User, len(users))
	for i, u := range users {
		res[i] = *convertUser(u) // 传递内容给res返回，convertUser返回的是首地址
	}
	return &res, nil
}

func (u userRepo) FindByOpenId(ctx context.Context, s string) (*biz.User, error) {
	first, err := u.data.db.User.Query().Where(user.Openid(s)).First(ctx)
	if err != nil {
		return nil, err
	}
	return convertUser(first), nil
}

func (u userRepo) Create(ctx context.Context, user *biz.User) error {
	return u.data.db.User.Create().
		SetOpenid(user.Openid).
		SetUsername(user.Username).
		SetGender(user.Gender).
		SetCarPictures(user.CarPictures).
		SetPhone(user.Phone).
		SetDescription(user.Description).
		Exec(ctx)
}

func (u userRepo) Update(ctx context.Context, user *biz.User) (*biz.User, error) {
	err := u.data.db.User.Update().
		SetPhone(user.Phone).
		SetDescription(user.Description).
		SetUsername(user.Username).
		SetGender(user.Gender).
		SetCarPictures(user.CarPictures).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	res, e := u.FindByOpenId(ctx, user.Openid)
	if e != nil {
		return nil, err
	}
	return res, nil
}

func (u userRepo) Delete(ctx context.Context, i int32) error {
	_, err := u.data.db.User.Delete().Where(user.ID(i)).Exec(ctx)
	return err
}

// NewUserRepo wire注入时执行，将biz层的接口实现，把接口套在userRepo结构体上。
func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}
