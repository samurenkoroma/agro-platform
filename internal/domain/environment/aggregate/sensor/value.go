package sensor

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Value struct {
	Current *float64

	Target *vo.Range

	LastTimestamp *time.Time
}
