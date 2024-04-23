package main

import (
	"github.com/unknowncode44/api-rest-clean-architecture/config"
	"github.com/unknowncode44/api-rest-clean-architecture/database"
	entities "github.com/unknowncode44/api-rest-clean-architecture/materials/entity"
)

func main2() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	supplierMigrate(db)

}

func supplierMigrate(db database.Database) {
	db.GetDb().Migrator().CreateTable(&entities.Supplier{})
	db.GetDb().CreateInBatches([]entities.Supplier{
		{
			Name:     "Equipel",
			SystemId: "EQP",
		},
		{
			Name:     "Relet",
			SystemId: "RLT",
		},
		{
			Name:     "Fibromarket",
			SystemId: "FIB",
		},
		{
			Name:     "DB Estructuras",
			SystemId: "DBE",
		},
		{
			Name:     "Copegan",
			SystemId: "CDT",
		},
	}, 10)

}
