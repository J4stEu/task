package auth

import "github.com/J4stEu/task/internal/model"

type CreateTokenDTO struct {
	UserID   uint32 `json:"uint32"`
	Lifetime uint64 `json:"lifetime"`
}

type LoginUserDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthData struct {
	JWT   string      `json:"jwt"`
	Token model.Token `json:"token"`
}
