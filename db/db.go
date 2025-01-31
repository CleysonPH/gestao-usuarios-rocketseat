package db

import "github.com/google/uuid"

type UserRepository struct {
	data map[uuid.UUID]User
}

func NewUserRepository() UserRepository {
	return UserRepository{
		data: make(map[uuid.UUID]User),
	}
}

func (ur *UserRepository) FindAll() []User {
	users := make([]User, 0, len(ur.data))
	for _, user := range ur.data {
		users = append(users, user)
	}
	return users
}

func (ur *UserRepository) Insert(u User) User {
	u.ID = uuid.New()
	ur.data[u.ID] = u
	return u
}

func (ur *UserRepository) FindById(id string) (User, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return User{}, err
	}

	u, ok := ur.data[uuid]
	if !ok {
		return User{}, ErrUserNotFound
	}

	return u, nil
}
