package error

import "errors"

var (
	ServerErr               = errors.New("server error")
	EmptyBodyRequest        = errors.New("empty body request")
	ErrorInvalidBodyRequest = errors.New("invalid body request")
	ErrorNoQueryParam       = errors.New("query param does not exists")

	ErrorEmptyQuery = errors.New("empty query")

	ErrorUnauthorizeUser = errors.New("unauthorized user")
)
