package main

import (
	"github.com/unknowncode44/api-rest-clean-architecture/config"
	"github.com/unknowncode44/api-rest-clean-architecture/database"
	entities "github.com/unknowncode44/api-rest-clean-architecture/materials/entity"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	costUpdateMigrate(db)

}

func costUpdateMigrate(db database.Database) {
	db.GetDb().Migrator().DropTable(&entities.CostUpdate{})
	db.GetDb().Migrator().CreateTable(&entities.CostUpdate{})
	db.GetDb().CreateInBatches([]entities.CostUpdate{
		{
			MaterialId: 1,
			OldPrice:   10.5,
			NewPrice:   11.5,
			SupplierId: 1,
		},
	}, 10)

	var material entities.Material
	db.GetDb().Preload("Supplier").Preload("Updates").First(&material, 1)

}
