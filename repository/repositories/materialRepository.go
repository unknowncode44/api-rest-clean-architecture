package repositories

import "github.com/unknowncode44/api-rest-clean-architecture/domain"

type MaterialRepository interface {
	InsertMaterialData(in *domain.InsertMaterialDto) error
}
