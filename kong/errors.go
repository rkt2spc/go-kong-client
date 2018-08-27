package kong

import "errors"

// Errors
var (
	ErrMissingConfig     = errors.New("missing config")
	ErrMissingHTTPClient = errors.New("missing http-client")
)
