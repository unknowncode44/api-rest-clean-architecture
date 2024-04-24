package domain

import (
	"time"
)

type (
	/*---------------------------*/
	/*       Proveedores         */
	/*---------------------------*/

	InsertSupplierDto struct {
		Id        uint32     `gorm:"primaryKey;autoIncrement" json:"id"`
		Name      string     `json:"supplier_name"`
		SystemId  string     `gorm:"size:3;" json:"system_id"`
		Materials []Material `gorm:"many2many:material_suppliers" json:"materials"`
		CreatedAt time.Time  `json:"createdAt"`
		UpdatedAt time.Time  `json:"updatedAt"`
	}

	Supplier struct {
		Id        uint32     `json:"id"`
		Name      string     `json:"supplier_name"`
		SystemId  string     `json:"system_id"`
		Materials []Material `gorm:"many2many:material_suppliers" json:"materials"`
		CreatedAt time.Time  `json:"createdAt"`
		UpdatedAt time.Time  `json:"updatedAt"`
	}
)
