package constant

import "errors"

var (
	INSUFFICIENT_FUND        error = errors.New("insufficient fund value")
	UNAUTHORIZED_TRANSACTION error = errors.New("transaction unauthorized")
)
