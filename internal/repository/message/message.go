package message

import (
	"fmt"

	"github.com/J4stEu/task/internal/message"
)

var (
	ErrRecordNotFound = fmt.Errorf("record %w", message.ErrNotFound)
)
