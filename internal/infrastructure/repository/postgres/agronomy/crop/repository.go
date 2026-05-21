package crop

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
)

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) repository.CropRepository {
	return &Repository{db: db}
}
