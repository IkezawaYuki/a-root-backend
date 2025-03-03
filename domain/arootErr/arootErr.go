package arootErr

import "errors"

var (
	ErrDuplicateEmail              = errors.New("email already in use")
	ErrNotFound                    = errors.New("not found")
	ErrAuthentication              = errors.New("authentication err")
	ErrAuthorization               = errors.New("authorization err")
	ErrDuplicateKey                = errors.New("duplicate key err")
	ErrEmailUsed                   = errors.New("email is already in use")
	ErrInstagramNotLinked          = errors.New("instagram not linked")
	ErrInstagramLinkedExpired      = errors.New("instagram linked expired")
	ErrInvalidInstagramTokenStatus = errors.New("invalid instagram token status")
	ErrInvalidPaymentType          = errors.New("invalid payment type")
	ErrInvalidPaymentStatus        = errors.New("invalid payment status")
)
