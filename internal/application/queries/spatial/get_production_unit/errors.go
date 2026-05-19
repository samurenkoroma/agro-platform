package getproductionunit

import "errors"

var (
	ErrProductionUnitNotFound = errors.New(
		"production unit not found",
	)
)
