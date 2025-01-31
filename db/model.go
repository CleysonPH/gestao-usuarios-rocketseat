package db

import "github.com/google/uuid"

type Id uuid.UUID

type User struct {
	ID        Id
	FirstName string
	LastName  string
	Biography string
}
