package data

import (
	"context"
	"service/internal/biz"
)

type ReserveRepo struct {
	data *Data
}

func (t *ReserveRepo) Create(r *biz.Reserve, ctx context.Context) (*biz.Reserve, error) {
	return nil, nil
}

func (t *ReserveRepo) Update(r *biz.Reserve, ctx context.Context) (*biz.Reserve, error) {
	return nil, nil
}

func (t *ReserveRepo) Delete(id int64, ctx context.Context) error {
	return nil
}
func (t *ReserveRepo) Get(in *biz.Reserve, ctx context.Context) ([]*biz.Reserve, error) {
	return nil, nil
}
func (t *ReserveRepo) List(ctx context.Context, currentPage, pageSize, userId int) ([]*biz.Reserve, error) {
	return nil, nil
}
