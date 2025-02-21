package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"net/http"
	v1 "service/api/blackpig/v1"
	"service/internal/conf"
)

type User struct {
	Id          int32
	Openid      string
	Phone       string
	Username    string
	Gender      string
	CarPictures []string
	Description string
}

func revertUser(user *v1.User) *User {
	return &User{
		Id:          user.Id,
		Openid:      user.Openid,
		Phone:       user.Phone,
		Username:    user.Username,
		Gender:      user.Gender,
		CarPictures: user.CarPictures,
	}
}

// UserRepo biz层定义User的基本调用方法，去data层实现
type UserRepo interface {
	FindAll(current, size int32, keyWord string, ctx context.Context) ([]User, error)
	FindByPhone(context.Context, string) (*User, error)
	FindById(context.Context, int32) (*User, error)
	FindByUsername(context.Context, string) (*[]User, error)
	FindByOpenId(context.Context, string) (*User, error)
	Create(context.Context, *User) error
	Update(context.Context, *User) (*User, error)
	Delete(context.Context, int32) error
}

// UserUseCase 实例化结构体的所有方法，具体方法实现在data中
type UserUseCase struct {
	wechat *conf.Wechat
	repo   UserRepo
}

// NewUserUseCase 实例化User这个结构体，相当于构造方法，其中的初始化UserRepo的方法会在data层实现
func NewUserUseCase(conf *conf.Wechat, repo UserRepo) *UserUseCase {
	return &UserUseCase{
		wechat: conf,
		repo:   repo,
	}
}

// 在 biz层使用data层实现的原子方法组成service所需要的方法

func (u *UserUseCase) Register(ctx context.Context, user *User) error {
	err := u.repo.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) Login(ctx context.Context, code string) (*User, error) {
	wechatURL := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		u.wechat.Appid,
		u.wechat.AppSecret,
		code,
	)
	resp, err := http.Get(wechatURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// TODO:这里应先查看resp.body的格式是怎样的
	fmt.Println(resp.Body)
	var result struct {
		OpenID     string `json:"openid"`
		SessionKey string `json:"session_key"`
		ErrCode    int    `json:"errcode"`
		ErrMsg     string `json:"errmsg"`
	}
	// 以防万一先重置err
	err = nil
	// 这里把resp.body中的内容转化存储到result中
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, errors.New(http.StatusBadRequest, "微信登录失败", result.ErrMsg)
	}
	// 去数据库查找这个用户的所有信息
	user, e := u.repo.FindByOpenId(ctx, result.OpenID)
	if e != nil {
		return nil, e
	}
	// 隐藏其OpenId不返回
	user.Openid = ""
	//Todo:生成token部分代码
	return user, nil
}

func (u *UserUseCase) UserList(ctx context.Context, current, size int32, keyWord string) ([]User, error) {
	user, err := u.repo.FindAll(current, size, keyWord, ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUseCase) FindUserByUsername(ctx context.Context, username string) (*[]User, error) {
	user, err := u.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUseCase) FindUserById(ctx context.Context, id int32) (*User, error) {
	user, err := u.repo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUseCase) FindUserByOpenId(ctx context.Context, openid string) (*User, error) {
	user, err := u.repo.FindByOpenId(ctx, openid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUseCase) FindUserByPhone(ctx context.Context, phone string) (*User, error) {
	user, err := u.repo.FindByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUseCase) UpdateUser(ctx context.Context, user *User) (*User, error) {
	res, err := u.repo.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *UserUseCase) RemoveUser(ctx context.Context, id int32) error {
	err := u.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
