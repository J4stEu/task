package middleware

import (
	"context"
	"net/http"

	"github.com/J4stEu/task/internal/api/http/response"
	"github.com/J4stEu/task/internal/api/http/session"
	"github.com/J4stEu/task/internal/model"
	"github.com/J4stEu/task/internal/service/auth"
	"github.com/J4stEu/task/internal/service/user"
)

type AuthMiddleware struct {
	session *session.Session
	service *AuthMiddlewareService
}

func NewAuthMiddleware(
	sess *session.Session,
	service *AuthMiddlewareService,
) *AuthMiddleware {

	return &AuthMiddleware{
		session: sess,
		service: service,
	}
}

type AuthMiddlewareService struct {
	Auth auth.Auth
	User user.User
}

type AuthData struct {
	Token model.Token
	User  model.User
}

type AuthDataKey string

const (
	ReqAuthDataKey AuthDataKey = "reqAuthData"
)

func (am *AuthMiddleware) AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, user, err := am.getAuthData(w, r)
		if err != nil {
			cookie, err := am.session.GetCookie(r)
			if err == nil && cookie != nil {
				_ = am.session.DeleteCookie(w, r) //nolint:errcheck
			}

			response.Response(w, http.StatusUnauthorized, nil)
			return
		}

		aData := AuthData{
			Token: token,
			User:  user,
		}
		aDataCtx := context.WithValue(r.Context(), ReqAuthDataKey, aData)
		next(w, r.WithContext(aDataCtx))
	}
}

func (am *AuthMiddleware) getAuthData(w http.ResponseWriter, r *http.Request) (model.Token, model.User, error) {
	tokenValue, err := am.session.GetTokenValue(r)
	if err != nil {
		return model.Token{}, model.User{}, err
	}

	token, err := am.service.Auth.GetTokenByValue(tokenValue)
	if err != nil {
		return model.Token{}, model.User{}, err
	}

	oldExp := token.Expiration
	if err = am.service.Auth.ActualizeToken(token); err != nil {
		return model.Token{}, model.User{}, err
	}
	newExp := token.Expiration
	if newExp != oldExp {
		if err = am.session.UpdateCookie(w, r, newExp); err != nil {
			return model.Token{}, model.User{}, err
		}
	}

	user, err := am.service.User.GetByID(token.UserID)
	if err != nil {
		return model.Token{}, model.User{}, err
	}

	return token, user, nil
}
