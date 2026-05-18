package timeline

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Item struct {
	ID          vo.ID
	Source      Source
	ReferenceID vo.ID
	Title       string
	Description *string
	Timestamp   time.Time
	Metadata    vo.Metadata
}
