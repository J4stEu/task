package session

import (
	"net/http"
	"time"
)

func (a *Session) GetCookie(r *http.Request) (*http.Cookie, error) {
	return r.Cookie(a.cookieName)
}

func (a *Session) SetCookie(w http.ResponseWriter, signedToken string, expiration time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:     a.cookieName,
		Value:    signedToken,
		Path:     a.cookiePath,
		HttpOnly: a.cookieHttpOnly,
		Expires:  expiration,
	})
}

func (a *Session) UpdateCookie(w http.ResponseWriter, r *http.Request, expiration time.Time) error {
	c, err := a.GetCookie(r)
	if err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:     a.cookieName,
		Value:    c.Value,
		Path:     a.cookiePath,
		HttpOnly: a.cookieHttpOnly,
		Expires:  expiration,
	})
	return nil
}

func (a *Session) DeleteCookie(w http.ResponseWriter, r *http.Request) error {
	_, err := a.GetCookie(r)
	if err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:     a.cookieName,
		Value:    "",
		Path:     a.cookiePath,
		HttpOnly: a.cookieHttpOnly,
		Expires:  time.Unix(0, 0),
	})
	return nil
}
