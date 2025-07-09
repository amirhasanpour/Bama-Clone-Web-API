package main

import (
	"log"

	"github.com/amirhasanpour/car-sale-management-wep-api/src/api"
	"github.com/amirhasanpour/car-sale-management-wep-api/src/config"
	"github.com/amirhasanpour/car-sale-management-wep-api/src/data/cache"
	"github.com/amirhasanpour/car-sale-management-wep-api/src/data/db"
)

func main() {
	cfg := config.GetConfig()
	
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatal(err)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		log.Fatal(err)
	}

	api.InitServer(cfg)
}