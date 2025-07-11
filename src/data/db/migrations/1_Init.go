package migrations

import (
	"github.com/amirhasanpour/car-sale-management-wep-api/src/config"
	"github.com/amirhasanpour/car-sale-management-wep-api/src/data/db"
	"github.com/amirhasanpour/car-sale-management-wep-api/src/data/models"
	"github.com/amirhasanpour/car-sale-management-wep-api/src/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

func Up_1() {
	database := db.GetDb()

	tables := []any{}

	country := models.Country{}
	city := models.City{}

	if !database.Migrator().HasTable(country) {
		tables = append(tables, country)
	}

	if !database.Migrator().HasTable(city) {
		tables = append(tables, city)
	}

	database.Migrator().CreateTable(tables...)
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)

}

func Down_1() {

}