package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/J4stEu/task/internal/model"
	"github.com/J4stEu/task/internal/repository/token"
	"github.com/J4stEu/task/internal/repository/user"
)

// Auth service implementation with mock repository
type AuthMock struct {
	apiSecret  []byte
	repository *AuthMockRepo
}

// Auth service repository
type AuthMockRepo struct {
	Token token.Token
	User  user.User
}

func NewMock(apiSecret []byte, tokenMockRepo *AuthMockRepo) Auth {
	return &AuthMock{
		apiSecret:  apiSecret,
		repository: tokenMockRepo,
	}
}

func (tm *AuthMock) Authenticate(authUser LoginUserDTO) (AuthData, error) {
	user, err := tm.repository.User.GetByLogin(authUser.Login)
	if err != nil {
		return AuthData{}, err
	}
	if err = verifyPassword(user.Password, authUser.Password); err != nil {
		return AuthData{}, err
	}
	jwToken, token, err := generateToken(user.ID, tm.apiSecret)
	if err != nil {
		return AuthData{}, err
	}
	newToken, err := tm.repository.Token.New(token)
	if err != nil {
		return AuthData{}, err
	}

	return AuthData{
		JWT:   jwToken,
		Token: newToken,
	}, nil
}

func verifyPassword(hashedPassword, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return ErrPasswordInvalid
	}
	return nil
}

func generateToken(userID uint32, apiSecret []byte) (string, model.Token, error) {
	createdAt := time.Now()
	lifetime := 28 * 24 * time.Hour
	expiration := createdAt.Add(lifetime)

	generatedToken := model.Token{
		UserID:     userID,
		Value:      uuid.NewString(),
		Lifetime:   lifetime, // 28 standard days,
		CreatedAt:  createdAt,
		Expiration: expiration,
		RevokedAt:  time.Time{},
	}

	jwt, err := jwtFromToken(generatedToken, apiSecret)
	if err != nil {
		return "", model.Token{}, err
	}

	return jwt, generatedToken, nil
}

func jwtFromToken(token model.Token, apiSecret []byte) (string, error) {
	jwToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": token.Value,
	})

	signedToken, err := jwToken.SignedString(apiSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (tm *AuthMock) ActualizeToken(model.Token) error {
	return nil
}

func (tm *AuthMock) GetTokenByValue(value string) (model.Token, error) {
	return tm.repository.Token.GetByValue(value)
}

func (tm *AuthMock) Revoke(token model.Token) (model.Token, error) {
	token.RevokedAt = time.Now()
	return tm.repository.Token.Update(token)
}
