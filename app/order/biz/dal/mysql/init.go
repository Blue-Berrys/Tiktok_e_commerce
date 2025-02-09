package mysql

import (
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/order/biz/dal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := "DYMall:DYMall@tcp(8.138.149.242:3306)/DYMall?charset=utf8mb4&parseTime=True&loc=Local" //先给写死
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	//add tracing
	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}

	//自动迁移，创建表
	if os.Getenv("GO_ENV") != "online" {
		_ = DB.AutoMigrate(&model.Order{},
			&model.OrderItem{})
	}
}
