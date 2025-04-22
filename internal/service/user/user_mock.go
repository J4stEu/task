package user

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/J4stEu/task/internal/model"
	"github.com/J4stEu/task/internal/repository/user"
)

// User service implementation with mock repository
type UserMock struct {
	repository *UserMockRepo
}

// User service repository
type UserMockRepo struct {
	User user.User
}

func NewMock(userMockRepo *UserMockRepo) User {
	return &UserMock{
		repository: userMockRepo,
	}
}

func (um *UserMock) New(createData CreateUserDTO) (model.User, error) {
	hashedPassword, err := hashPassword(createData.Password)
	if err != nil {
		return model.User{}, err
	}

	return um.repository.User.New(model.User{
		Login:    createData.Login,
		Password: hashedPassword,
	})
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (um *UserMock) GetByID(id uint32) (model.User, error) {
	return um.repository.User.GetByID(id)
}
