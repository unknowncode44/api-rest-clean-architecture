package domain

import (
	"time"
)

type (
	/*--------------------------*/
	/*        Materiales        */
	/*--------------------------*/

	InsertedMaterialDto struct {
		Id                    uint32       `gorm:"primaryKey;autoIncrement" json:"id"`
		SystemCode            string       `gorm:"unique; not null;" json:"system_code"`
		Description           string       `json:"description"`
		LastPrice             float64      `gorm:"type:decimal(10,2);" json:"last_price"`
		LastPriceCurrency     string       `gorm:"size:3;" json:"last_price_currency"`
		LastPriceExchange     float64      `gorm:"type:decimal(10,2);" json:"last_price_exchange"`
		LastPriceExchangeDate string       `gorm:"type:decimal(10,2);" json:"last_price_exchange_date"`
		CreatedAt             time.Time    `json:"createdAt"`
		UpdatedAt             time.Time    `json:"updatedAt"`
		Category              string       `json:"category"`
		SubCategory           string       `json:"subcategory"`
		Updates               []CostUpdate `gorm:"ForeignKey:Id" json:"updates"`
		Supplier              []Supplier   `gorm:"ForeignKey:Id" json:"supplier"`
		MeasurementUnit       string       `json:"measurement_unit"`
	}

	Material struct {
		Id                uint32       `gorm:"primaryKey" json:"id"`
		SystemCode        string       `json:"system_code"`
		Description       string       `json:"description"`
		LastPrice         float64      `json:"last_price"`
		LastPriceCurrency string       `json:"last_price_currency"`
		CreatedAt         time.Time    `json:"createdAt"`
		UpdatedAt         time.Time    `json:"updatedAt"`
		Category          string       `json:"category"`
		SubCategory       string       `json:"subcategory"`
		Updates           []CostUpdate `gorm:"ForeignKey:Id" json:"updates"`
		Supplier          []Supplier   `gorm:"ForeignKey:Id" json:"supplier"`
		MeasurementUnit   string       `json:"measurement_unit"`
	}
)
