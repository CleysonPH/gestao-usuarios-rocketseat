package api

import "errors"

var (
	errEmptyFirstName = errors.New("please provide a first_name")
	errEmptyLastName  = errors.New("please provide a last_name")
	errEmptyBiography = errors.New("please provide a biography")
)
