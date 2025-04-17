package service

import (
	"context"
	v2 "service/api/blackpig/v2"
	"service/internal/biz"
)

type ReserveService struct {
	v2.UnimplementedTimeTableServiceServer
	uc *biz.Reserve
}

func revertReserve(reserve *v2.TimeTable) *biz.Reserve {
	return &biz.Reserve{
		StartTime: reserve.StartTime,
		EndTime:   reserve.EndTime,
		Day:       reserve.Day,
		UserId:    reserve.UserId,
		OrderId:   reserve.OrderedId,
		CreatedAt: reserve.CreatedAt,
		UpdatedAt: reserve.UpdatedAt,
	}
}

func (s *ReserveService) AddTimeTables(ctx context.Context, request *v2.Ordered) []*v2.TimeTable {
	reserves := make([]*biz.Reserve, 0)
	for _, reserve := range request.TimeTables {
		reserves = append(reserves, revertReserve(reserve))
	}
	err, out := s.uc.AddReserve(reserves, ctx)
	if err != nil {
		return nil
	}
	return out
}

func (s *ReserveService) UpdateTimeTables(ctx context.Context, request *v2.Ordered) *v2.Ordered {
	reserves := make([]*biz.Reserve, 0)
	for _, reserve := range request.TimeTables {
		reserves = append(reserves, revertReserve(reserve))
	}
	err, out := s.uc.UpdateReserve(reserves, ctx)
	if err != nil {
		return nil
	}
	return &v2.Ordered{TimeTables: out}
}
func (s *ReserveService) DeleteTimeTables(ctx context.Context, request *v2.DeleteTimeTableRequest) *v2.DeleteTimeTableResponse {
	id := request.Id
	if id == 0 {
		return &v2.DeleteTimeTableResponse{Success: false}
	}
	err := s.uc.DeleteReserve(id, ctx)
	if err != nil {
		return &v2.DeleteTimeTableResponse{Success: false}
	}
	return &v2.DeleteTimeTableResponse{Success: true}
}

// GetOrderedByOrderedId 导员获取所有的被预约的时间
func (s *ReserveService) GetOrderedByOrderedId(ctx context.Context, request *v2.GerOrderedRequest) *v2.Ordered {
	in := &biz.Reserve{
		OrderId: request.OrderedId,
	}
	reserves, err := s.uc.GetReserve(in, ctx)
	if err != nil {
		return &v2.Ordered{}
	}
	return &v2.Ordered{TimeTables: reserves}
}

// GetTimeTablesOrdered 导员获取某一天所有的被预约的时间
func (s *ReserveService) GetTimeTablesOrdered(ctx context.Context, request *v2.GetTimeTablesRequest) *v2.Ordered {
	in := &biz.Reserve{
		UserId: request.UserId,
		Day:    request.Day,
	}
	reserves, err := s.uc.GetReserve(in, ctx)
	if err != nil {
		return &v2.Ordered{}
	}
	return &v2.Ordered{TimeTables: reserves}
}
