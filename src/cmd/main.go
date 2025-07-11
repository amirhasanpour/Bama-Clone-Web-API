package main

import (
	"github.com/amirhasanpour/car-sale-management-wep-api/src/api"
	"github.com/amirhasanpour/car-sale-management-wep-api/src/config"
	"github.com/amirhasanpour/car-sale-management-wep-api/src/data/cache"
	"github.com/amirhasanpour/car-sale-management-wep-api/src/data/db"
	"github.com/amirhasanpour/car-sale-management-wep-api/src/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}

	api.InitServer(cfg)
}