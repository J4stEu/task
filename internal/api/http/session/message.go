package session

import (
	"fmt"

	"github.com/J4stEu/task/internal/message"
)

var (
	errInvalidToken = fmt.Errorf("token is %w", message.ErrInvalid)
	errExtractToken = fmt.Errorf("failed to extract token")
)
