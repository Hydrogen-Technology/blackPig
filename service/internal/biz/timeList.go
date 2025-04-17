package biz

import (
	"context"
	v2 "service/api/blackpig/v2"
	"service/internal/data"
)

type TimeList struct {
	Id     int64
	UserId int32
	Times  []Times
}

type TimeListRepo interface {
	List(userId int32, ctx context.Context) (*TimeList, error)
}

// !! 把biz的TimeList转换成v2的TimeList，然后返回给api层
func convertTimeList(t *TimeList) *v2.TimeList {
	return &v2.TimeList{
		Id:     t.Id,
		UserId: t.UserId,
		Times:  convertTimesSlice(t.Times),
	}
}

// !! 循环把biz的Times转换成v2的Times，然后返回给api层
func convertTimesSlice(ts []Times) []*v2.Times {
	var out []*v2.Times
	for _, t := range ts {
		out = append(out, convertTimes(&t))
	}
	return out
}

func (t *TimeList) List(ctx context.Context) (*v2.TimeList, error) {
	repo := &data.TimeListRepo{}
	out, err := repo.List(t.UserId, ctx)
	if err != nil {
		return nil, err
	}
	return convertTimeList(out), nil
}
