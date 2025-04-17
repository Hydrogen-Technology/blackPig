package biz

import (
	v2 "service/api/blackpig/v2"
	"service/internal/data"
)

type Times struct {
	Id        int64
	UserId    int32
	StartTime string
	TimeSpace string
}

type TimesRepo interface {
	Create(times *Times) (*Times, error)
	Update(times *Times) (*Times, error)
	Delete(times *Times) error
	Get(times *Times) (*Times, error)
}

// !! 把biz的Times类型转换成v2的Times类型
func convertTimes(t *Times) *v2.Times {
	return &v2.Times{
		Id:        t.Id,
		UserId:    t.UserId,
		StartTime: t.StartTime,
		TimeSpace: t.TimeSpace,
	}
}

func (t *Times) CreateTimes() (*v2.Times, error) {
	repo := &data.TimesRepo{}
	in, err := repo.Create(t)
	if err != nil {
		return nil, err
	}
	return convertTimes(in), nil
}

func (t *Times) UpdateTimes() (error, *v2.Times) {
	repo := &data.TimesRepo{}
	in, err := repo.Update(t)
	if err != nil {
		return err, nil
	}
	return nil, convertTimes(in)
}

func (t *Times) DeleteTimes() error {
	repo := &data.TimesRepo{}
	return repo.Delete(t)
}

func (t *Times) GetTime() (error, *v2.Times) {
	repo := &data.TimesRepo{}
	in, err := repo.Get(t)
	if err != nil {
		return err, nil
	}
	return nil, convertTimes(in)
}
