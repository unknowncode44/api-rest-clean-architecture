package repositories

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/unknowncode44/api-rest-clean-architecture/database"
	"github.com/unknowncode44/api-rest-clean-architecture/domain"
)

type materialPgRepository struct {
	db database.Database
}

func NewMaterialPgRepository(db database.Database) domain.MaterialRepository {
	return &materialPgRepository{db: db}
}

func (r *materialPgRepository) InsertMaterialData(in *domain.InsertMaterialDto) error {
	data := &domain.Material{
		SystemCode:      in.SystemCode,
		Description:     in.Description,
		Category:        in.Category,
		SubCategory:     in.SubCategory,
		Supplier:        in.Supplier,
		MeasurementUnit: in.MeasurementUnit,
	}

	result := r.db.GetDb().Create(data)

	if result.Error != nil {
		log.Errorf("InsertMaterialData: %v", result.Error)
		return result.Error
	}

	log.Debugf("InsertMaterialData: %v", result.RowsAffected)
	return nil
}

func (r *materialPgRepository) FetchMaterialData(material *int) (domain.Material, error) {
	var mat domain.Material
	result := r.db.GetDb().Preload("Suppliers").Preload("Updates").First(mat, material)

	if result.Error != nil {
		log.Errorf("InsertMaterialData: %v", result.Error)
		fmt.Println(result.Error)
		return mat, result.Error
	}

	fmt.Println(mat)

	log.Debugf("InsertMaterialData: %v", result.RowsAffected)
	return mat, nil
}

func (r *materialPgRepository) UpdateMaterialData(material *domain.InsertMaterialDto) error {

	result := r.db.GetDb().Save(material)

	if result.Error != nil {
		log.Errorf("UpdateMaterialData: %v", result.Error)
		return result.Error
	}

	log.Debugf("UpdateMaterialData: %v", result.RowsAffected)
	return nil
}

func (r *materialPgRepository) FindAllMaterials() ([]domain.Material, error) {
	var materials []domain.Material
	result := r.db.GetDb().Find(&materials)

	if result.Error != nil {
		log.Errorf("RetrieveAllMaterials: %v", result.Error)
		return materials, result.Error
	}

	log.Debugf("RetrieveAllMaterials: %v", result.RowsAffected)
	return materials, nil

}
