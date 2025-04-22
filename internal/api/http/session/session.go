package session

type Session struct {
	apiSecret []byte

	cookieName     string
	cookiePath     string
	cookieHttpOnly bool

	tokenKey          string
	bearerTokenHeader string
	tokenHeader       string
}

func New(apiSecret []byte) *Session {
	return &Session{
		apiSecret: apiSecret,

		cookieName:     "auth_token",
		cookiePath:     "/",
		cookieHttpOnly: true,

		tokenKey:          "token",
		bearerTokenHeader: "Authorization",
		tokenHeader:       "x-token",
	}
}
