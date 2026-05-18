package climatezone

import (
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Target struct {
	Temperature *vo.Range
	Humidity    *vo.Range
	CO2         *vo.Range
}
