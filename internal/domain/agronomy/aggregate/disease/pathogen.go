package disease

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Pathogen struct {
	ID        vo.ID
	Name      string
	Type      string
	Metadata  vo.Metadata
	CreatedAt time.Time
}
