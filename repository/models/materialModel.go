package models

import "github.com/unknowncode44/api-rest-clean-architecture/domain"

type AddMaterialData struct {
	SystemCode      string            `json:"system_code"`
	Description     string            `json:"description"`
	Category        string            `json:"category"`
	SubCategory     string            `json:"subcategory"`
	Supplier        []domain.Supplier `json:"supplier"`
	MeasurementUnit string            `json:"measurement_unit"`
}
