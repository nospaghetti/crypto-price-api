package apperr

import "errors"

var (
	UpstreamUnavailable = errors.New("upstream unavailable")
	InternalError       = errors.New("internal error")
)
