package tx

import "errors"

var (
	ErrTransactionRequired = errors.New("transaction required")
	ErrCommitFailed        = errors.New("commit failed")
	ErrRollbackFailed      = errors.New("rollback failed")
)
