package errors

import "errors"

var (
	ErrSiteNotAvailable = errors.New("there's no scraper implemented for this website")
)
