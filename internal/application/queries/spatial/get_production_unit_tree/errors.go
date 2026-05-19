package getproductionunittree

import "errors"

var (
	ErrRootNotFound = errors.New(
		"root production unit not found",
	)
)
