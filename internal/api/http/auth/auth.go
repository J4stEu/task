package auth

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/J4stEu/task/internal/api/http/response"
	authService "github.com/J4stEu/task/internal/service/auth"
	userService "github.com/J4stEu/task/internal/service/user"
)

func (a *Auth) register(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Response(w, http.StatusBadRequest, err)
		return
	}
	registerData := userService.CreateUserDTO{}
	if err = json.Unmarshal(body, &registerData); err != nil {
		response.Response(w, http.StatusBadRequest, err)
		return
	}

	user, err := a.service.User.New(registerData)
	if err != nil {
		response.Response(w, http.StatusBadRequest, err)
		return
	}

	response.Response(w, http.StatusCreated, user)
}

func (a *Auth) login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Response(w, http.StatusBadRequest, err)
		return
	}
	loginData := authService.LoginUserDTO{}
	if err = json.Unmarshal(body, &loginData); err != nil {
		response.Response(w, http.StatusBadRequest, err)
		return
	}

	auth, err := a.service.Auth.Authenticate(loginData)
	if err != nil {
		response.Response(w, http.StatusBadRequest, err)
		return
	}

	a.session.SetCookie(w, auth.JWT, auth.Token.Expiration)
	response.Response(w, http.StatusOK, auth)
}
