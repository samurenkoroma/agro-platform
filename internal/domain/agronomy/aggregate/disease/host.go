package disease

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Host struct {
	ID        vo.ID
	Species   string
	Variety   *string
	Metadata  vo.Metadata
	CreatedAt time.Time
}
