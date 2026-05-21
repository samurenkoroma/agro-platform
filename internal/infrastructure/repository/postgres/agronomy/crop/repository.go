package crop

import (
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
)

type Repository struct {
	db uow.DB
}

func NewCropRepository(db uow.DB) repository.CropRepository {
	return &Repository{db: db}
}
