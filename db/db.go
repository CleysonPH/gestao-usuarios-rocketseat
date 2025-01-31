package db

type userRepository struct {
	data map[Id]User
}

func NewUserRepository() userRepository {
	return userRepository{
		data: make(map[Id]User),
	}
}

func (ur *userRepository) FindAll() []User {
	users := make([]User, 0, len(ur.data))
	for _, user := range ur.data {
		users = append(users, user)
	}
	return users
}
