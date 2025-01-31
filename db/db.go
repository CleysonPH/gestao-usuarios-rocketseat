package db

import (
	"sync"

	"github.com/google/uuid"
)

type UserRepository struct {
	data map[uuid.UUID]User
	m    sync.Mutex
}

func NewUserRepository() UserRepository {
	return UserRepository{
		data: make(map[uuid.UUID]User),
	}
}

func (ur *UserRepository) FindAll() []User {
	ur.m.Lock()
	defer ur.m.Unlock()

	users := make([]User, 0, len(ur.data))
	for _, user := range ur.data {
		users = append(users, user)
	}
	return users
}

func (ur *UserRepository) Insert(u User) User {
	ur.m.Lock()
	defer ur.m.Unlock()

	u.ID = uuid.New()
	ur.data[u.ID] = u
	return u
}

func (ur *UserRepository) FindById(id string) (User, error) {
	ur.m.Lock()
	defer ur.m.Unlock()

	uuid, err := uuid.Parse(id)
	if err != nil {
		return User{}, ErrInvalidUUID
	}

	u, ok := ur.data[uuid]
	if !ok {
		return User{}, ErrUserNotFound
	}

	return u, nil
}

func (ur *UserRepository) DeleteById(id string) error {
	ur.m.Lock()
	defer ur.m.Unlock()

	uuid, err := uuid.Parse(id)
	if err != nil {
		return ErrInvalidUUID
	}

	_, ok := ur.data[uuid]
	if !ok {
		return ErrUserNotFound
	}

	delete(ur.data, uuid)
	return nil
}

func (ur *UserRepository) UpdateById(id string, u User) (User, error) {
	ur.m.Lock()
	defer ur.m.Unlock()

	uuid, err := uuid.Parse(id)
	if err != nil {
		return User{}, ErrInvalidUUID
	}

	user, ok := ur.data[uuid]
	if !ok {
		return User{}, ErrUserNotFound
	}

	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.Biography = u.Biography

	ur.data[uuid] = user

	return user, nil
}
