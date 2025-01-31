package db

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrInvalidUUID = errors.New("invalid uuid format")
