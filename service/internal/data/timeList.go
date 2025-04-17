package data

import (
	"context"
	"service/internal/biz"
)

type TimeListRepo struct {
	data *Data
}

func (t *TimeListRepo) List(userId int32, ctx context.Context) (*biz.TimeList, error) {
	return nil, nil
}
