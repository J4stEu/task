package auth

import (
	"fmt"
	"net/http"

	"github.com/J4stEu/task/internal/api/http/session"
	"github.com/J4stEu/task/internal/service/auth"
	"github.com/J4stEu/task/internal/service/user"
)

type Auth struct {
	path    string
	session *session.Session
	service *AuthService
}

type AuthService struct {
	Auth auth.Auth
	User user.User
}

func Init(mux *http.ServeMux, sess *session.Session, path string, service *AuthService) {
	taskAPI := new(path, sess, service)
	taskAPI.init(mux)
}

func new(path string, sess *session.Session, service *AuthService) *Auth {
	return &Auth{
		path:    path,
		session: sess,
		service: service,
	}
}

func (a *Auth) init(mux *http.ServeMux) {
	taskPath := a.path + "/auth"

	mux.HandleFunc(fmt.Sprintf("POST %s/register", taskPath), a.register)
	mux.HandleFunc(fmt.Sprintf("POST %s/login", taskPath), a.login)
}
