package auth

import (
	"fmt"

	"github.com/J4stEu/task/internal/message"
)

var (
	ErrPasswordInvalid = fmt.Errorf("password is %w", message.ErrInvalid)
)
