package entities

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

	/*---------------------------*/
	/* Actualizaciones de precio */
	/*---------------------------*/

	InsertCostUpdateDto struct {
		Id         uint32   `gorm:"primaryKey;autoIncrement" json:"id"`
		MaterialId uint32   `gorm:"foreignKey:MaterialId; references:Material.Id" json:"material_id"`
		Material   Material `json:"material"`
		OldPrice   float64  `gorm:"type:decimal(10,2);" json:"old_price"`
		NewPrice   float64  `gorm:"type:decimal(10,2);" json:"new_price"`
		SupplierId uint32   `gorm:"foreignKey:SupplierId; references:Supplier.Id" json:"supplier_id"`
		Supplier   Supplier `json:"supplier"`
	}

	CostUpdate struct {
		Id         uint32   `json:"id"`
		MaterialId uint32   `gorm:"foreignKey:MaterialId; references:Material.Id" json:"material_id"`
		Material   Material `json:"material"`
		OldPrice   float64  `json:"old_price"`
		NewPrice   float64  `json:"new_price"`
		SupplierId uint32   `gorm:"foreignKey:SupplierId; references:Supplier.Id" json:"supplier_id"`
		Supplier   Supplier `json:"supplier"`
	}

	/*---------------------------*/
	/*       Proveedores         */
	/*---------------------------*/

	InsertSupplierDto struct {
		Id        uint32     `gorm:"primaryKey;autoIncrement" json:"id"`
		Name      string     `json:"supplier_name"`
		SystemId  string     `gorm:"size:3;" json:"system_id"`
		Materials []Material `gorm:"many2many:material_suppliers" json:"materials"`
	}

	Supplier struct {
		Id        uint32     `json:"id"`
		Name      string     `json:"supplier_name"`
		SystemId  string     `json:"system_id"`
		Materials []Material `gorm:"many2many:material_suppliers" json:"materials"`
	}
)
