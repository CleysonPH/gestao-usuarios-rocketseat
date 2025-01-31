package db

import "github.com/google/uuid"

type UserRepository struct {
	data map[uuid.UUID]User
}

func NewUserRepository() UserRepository {
	id := uuid.New()
	return UserRepository{
		data: map[uuid.UUID]User{
			id: {
				ID:        id,
				FirstName: "Cleyson",
				LastName:  "Lima",
				Biography: "Test",
			},
		},
	}
}

func (ur *UserRepository) FindAll() []User {
	users := make([]User, 0, len(ur.data))
	for _, user := range ur.data {
		users = append(users, user)
	}
	return users
}
