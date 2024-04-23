package main

import (
	"fmt"

	"github.com/unknowncode44/api-rest-clean-architecture/config"
	"github.com/unknowncode44/api-rest-clean-architecture/database"
	entities "github.com/unknowncode44/api-rest-clean-architecture/materials/entity"
)

func main3() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	materialsMigrate(db)
}

func materialsMigrate(db database.Database) {
	db.GetDb().Migrator().CreateTable(&entities.Material{})
	db.GetDb().CreateInBatches([]entities.Material{

		{
			SystemCode:        "MAT-1234",
			Description:       "Conector caño flexible 1/2",
			LastPrice:         10.50,
			LastPriceCurrency: "USD",
			Category:          "ELE",
			SubCategory:       "Materiales",
			MeasurementUnit:   "C/U",
		},
		{
			SystemCode:        "MAT-1233",
			Description:       "Caja de Cable UTP Cat5e Multilan x 305m",
			LastPrice:         333.00,
			LastPriceCurrency: "USD",
			Category:          "NET",
			SubCategory:       "Cables",
			MeasurementUnit:   "C/U",
		},
		{
			SystemCode:        "MAT-1232",
			Description:       "Patch Cord SC/LC 1.5m",
			LastPrice:         13.50,
			LastPriceCurrency: "USD",
			Category:          "NET",
			SubCategory:       "Fibra",
			MeasurementUnit:   "C/U",
		},
		{
			SystemCode:        "MAT-1231",
			Description:       "Soporte Artesanal 1 1/2",
			LastPrice:         29.50,
			LastPriceCurrency: "USD",
			Category:          "EST",
			SubCategory:       "Estructuras Metalicas",
			MeasurementUnit:   "C/U",
		},
		{
			SystemCode:        "MAT-1230",
			Description:       "Cupla 1/2",
			LastPrice:         10.50,
			LastPriceCurrency: "USD",
			Category:          "EST",
			SubCategory:       "Estrucuturas Metalicas",
			MeasurementUnit:   "C/U",
		},
	}, 10).Preload("Suppliers").Preload("Updates")

	mat := []entities.Material{}
	db.GetDb().Debug().Find(&mat)

	fmt.Print(mat)

}
