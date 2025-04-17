package service

import (
	"context"
	v2 "service/api/blackpig/v2"
	"service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type TimesService struct {
	v2.UnimplementedTimeServiceServer
	log *log.Helper
	uc  *biz.TimeList
}

func revertTimes(t *v2.TimeRequest) *biz.Times {
	if t == nil {
		return nil
	}
	return &biz.Times{
		UserId:    t.UserId,
		TimeSpace: t.TimeSpace,
		StartTime: t.StartTime,
	}
}

func (s *TimesService) CreateTime(ctx context.Context, request *v2.TimeRequest) *v2.Times {
	in := revertTimes(request)
	t, err := in.CreateTimes()
	if err != nil {
		return nil
	}
	return t
}

func (s *TimesService) UpdateTime(ctx context.Context, request *v2.TimeRequest) *v2.Times {
	in := revertTimes(request)
	err, t := in.UpdateTimes()
	if err != nil {
		return nil
	}
	return t
}

func (s *TimesService) DeleteTime(ctx context.Context, request *v2.TimeRequest) bool {
	in := revertTimes(request)
	err := in.DeleteTimes()
	if err != nil {
		return false
	}
	return true
}

func (s *TimesService) ListTime(ctx context.Context, request *v2.TimeListRequest) *v2.TimeList {
	in := &biz.TimeList{
		UserId: request.Id,
	}
	t, err := in.List(ctx)
	if err != nil {
		return nil
	}
	return t
}
