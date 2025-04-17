package biz

import (
	"context"
	v2 "service/api/blackpig/v2"
	"service/internal/data"
)

type Reserve struct {
	Id        int32
	StartTime string
	EndTime   string
	Day       string
	UserId    int32
	OrderId   int32
	CreatedAt string
	UpdatedAt string
}

type ReserveRepo interface {
	Create(r *Reserve, ctx context.Context) (*Reserve, error)
	Get(id int32, ctx context.Context) (*Reserve, error)
	List(ctx context.Context, currentPage, pageSize, userId int) ([]*Reserve, error)
	Update(r *Reserve, ctx context.Context) error
	Delete(id int32, ctx context.Context) error
}

func ConvertReserve(r *Reserve) *v2.TimeTable {
	return &v2.TimeTable{
		Id:        r.Id,
		StartTime: r.StartTime,
		EndTime:   r.EndTime,
		Day:       r.Day,
		UserId:    r.UserId,
		OrderedId: r.OrderId,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

func (r *Reserve) AddReserve(reserve []*Reserve, ctx context.Context) (error, []*v2.TimeTable) {
	repo := &data.ReserveRepo{}
	reserves := make([]*v2.TimeTable, 0)
	for _, r := range reserve {
		temp, err := repo.Create(r, ctx)
		if err != nil {
			return err, nil
		}
		reserves = append(reserves, ConvertReserve(temp))
	}
	return nil, reserves
}

// GetReserve 根据所获取的timetable内的某几项信息，获取所有对应的reserve信息
func (r *Reserve) GetReserve(in *Reserve, ctx context.Context) ([]*v2.TimeTable, error) {
	repo := &data.ReserveRepo{}
	out, err := repo.Get(in, ctx)
	if err != nil {
		return nil, err
	}
	reserves := make([]*v2.TimeTable, 0)
	for _, r := range out {
		reserves = append(reserves, ConvertReserve(r))
	}
	return reserves, nil
}
func (r *Reserve) ListReserve(ctx context.Context, currentPage, pageSize, userId int) ([]*v2.TimeTable, error) {
	repo := &data.ReserveRepo{}
	reserves, err := repo.List(ctx, currentPage, pageSize, userId)
	if err != nil {
		return nil, err
	}
	reservesV2 := make([]*v2.TimeTable, 0)
	for _, reserve := range reserves {
		reservesV2 = append(reservesV2, ConvertReserve(reserve))
	}
	return reservesV2, nil
}

func (r *Reserve) DeleteReserve(id int64, ctx context.Context) error {
	repo := &data.ReserveRepo{}
	return repo.Delete(id, ctx)
}

func (r *Reserve) UpdateReserve(in []*Reserve, ctx context.Context) (error, []*v2.TimeTable) {
	repo := &data.ReserveRepo{}
	reserves := make([]*v2.TimeTable, 0)
	for _, r := range in {
		temp, err := repo.Update(r, ctx)
		if err != nil {
			return err, nil
		}
		reserves = append(reserves, ConvertReserve(temp))
	}
	return nil, reserves
}
