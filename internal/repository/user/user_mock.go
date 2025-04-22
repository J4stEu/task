package user

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/J4stEu/task/internal/model"
	repoMessage "github.com/J4stEu/task/internal/repository/message"
)

// User repository implementation with mock data
type UserMock struct{}

// Get User repository implementation with mock data
func NewMock() User {
	return &UserMock{}
}

// Mock data
func getMockedUsers() []model.User {
	users := make([]model.User, 0, 10)

	randomDays := 24 * (1 + rand.Intn(30-1))
	for i := range 10 {
		createdAt := time.Now().Add(-time.Hour * time.Duration(randomDays))

		users = append(users, model.User{
			ID: uint32(i) + 1,

			Login:    fmt.Sprintf("User %v", uint32(i)+1),
			Password: "",

			CreatedAt: createdAt,
			UpdatedAt: createdAt,
		})
	}

	adminCreatedAt := time.Now().Add(-time.Hour * time.Duration(randomDays))
	users = append(users, model.User{
		ID:        11,
		Login:     "admin",
		Password:  "$2a$10$OLQy6vlCtfupVFJgdSJ80OMZc/77e0AiWCIfqBbqbgYDxmXJrhNn6", // test
		CreatedAt: adminCreatedAt,
		UpdatedAt: adminCreatedAt,
	})

	return users
}

func (*UserMock) New(user model.User) (model.User, error) {
	users := getMockedUsers()
	user.ID = users[len(users)-1].ID + 1
	user.Login = fmt.Sprintf("Task %v", user.ID)

	createdAt := time.Now()
	user.CreatedAt = createdAt
	user.UpdatedAt = createdAt

	return user, nil
}

func (*UserMock) GetByID(id uint32) (model.User, error) {
	users := getMockedUsers()
	for i := range users {
		user := users[i]
		if user.ID == id {
			return user, nil
		}
	}

	return model.User{}, repoMessage.ErrRecordNotFound
}

func (*UserMock) GetByLogin(login string) (model.User, error) {
	users := getMockedUsers()
	for i := range users {
		user := users[i]
		if user.Login == login {
			return user, nil
		}
	}

	return model.User{}, repoMessage.ErrRecordNotFound
}
