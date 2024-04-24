package domain

import (
	"time"
)

type (
	/*---------------------------*/
	/* Actualizaciones de precio */
	/*---------------------------*/

	InsertCostUpdateDto struct {
		Id         uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
		MaterialId uint32    `gorm:"foreignKey:MaterialId; references:Material.Id" json:"material_id"`
		OldPrice   float64   `gorm:"type:decimal(10,2);" json:"old_price"`
		NewPrice   float64   `gorm:"type:decimal(10,2);" json:"new_price"`
		SupplierId uint32    `gorm:"foreignKey:SupplierId; references:Supplier.Id" json:"supplier_id"`
		CreatedAt  time.Time `json:"createdAt"`
	}

	CostUpdate struct {
		Id         uint32    `json:"id"`
		MaterialId uint32    `gorm:"foreignKey:MaterialId; references:Material.Id" json:"material_id"`
		OldPrice   float64   `json:"old_price"`
		NewPrice   float64   `json:"new_price"`
		SupplierId uint32    `gorm:"foreignKey:SupplierId; references:Supplier.Id" json:"supplier_id"`
		CreatedAt  time.Time `json:"createdAt"`
	}
)
