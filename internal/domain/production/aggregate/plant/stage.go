package plant

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type PlantStage struct {
	StageID vo.ID

	ChangedAt time.Time
}
