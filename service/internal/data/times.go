package data

import "service/internal/biz"

type TimesRepo struct {
	data *Data
}

func (t *TimesRepo) Create(times *biz.Times) (*biz.Times, error) {
	return nil, nil
}

func (t *TimesRepo) Update(times *biz.Times) (*biz.Times, error) {
	return nil, nil
}

func (t *TimesRepo) Delete(times *biz.Times) error {
	return nil
}
func (t *TimesRepo) Get(times *biz.Times) (*biz.Times, error) {
	return nil, nil
}
