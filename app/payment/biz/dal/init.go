package dal

import (
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/payment/biz/dal/mysql"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
