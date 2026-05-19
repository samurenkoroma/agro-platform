package query

import "errors"

var (
	ErrQueryHandlerNotFound = errors.New("query handler not found")

	ErrInvalidQuery = errors.New("invalid query")
)
