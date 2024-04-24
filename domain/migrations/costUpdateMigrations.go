package main

import (
	"github.com/unknowncode44/api-rest-clean-architecture/config"
	"github.com/unknowncode44/api-rest-clean-architecture/database"
	"github.com/unknowncode44/api-rest-clean-architecture/domain"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	costUpdateMigrate(db)

}

func costUpdateMigrate(db database.Database) {
	db.GetDb().Migrator().DropTable(&domain.CostUpdate{})
	db.GetDb().Migrator().CreateTable(&domain.CostUpdate{})
	db.GetDb().CreateInBatches([]domain.CostUpdate{
		{
			MaterialId: 1,
			OldPrice:   10.5,
			NewPrice:   11.5,
			SupplierId: 1,
		},
	}, 10)

}
