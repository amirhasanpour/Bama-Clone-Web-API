package main

import (
	"github.com/amirhasanpour/car-sale-management-wep-api/src/api"
	"github.com/amirhasanpour/car-sale-management-wep-api/src/config"
	"github.com/amirhasanpour/car-sale-management-wep-api/src/data/cache"
)

func main() {
	cfg := config.GetConfig()
	cache.InitRedis(cfg)
	defer cache.CloseRedis()
	api.InitServer(cfg)
}