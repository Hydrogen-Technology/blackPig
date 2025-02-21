package data

import (
	"context"
	"service/internal/conf"
	"service/internal/ent"
	"service/internal/ent/migrate"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql" //必须导入这个包才能成功导入驱动，但是没有提示
	"github.com/google/wire"
)

// ProviderSet 这里就是实现data层操纵数据库的接口套在对应数据库
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	db *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	// 使用ent链接数据库
	client, err := ent.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	if err != nil {
		return nil, nil, err
	}
	if err := client.Schema.Create(context.Background(), migrate.WithDropIndex(true)); err != nil {
		log.Fatalf("data: 创建数据库失败: %v", err)
	}
	return &Data{
			db: client,
		}, func() {
			if err := client.Close(); err != nil {
				return
			}
		}, nil
}
