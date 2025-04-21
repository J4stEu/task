package task

import (
	"fmt"

	"github.com/J4stEu/task/internal/message"
)

var (
	ErrNothingToUpdate = fmt.Errorf("failed to %w task: nothing to update", message.ErrUpdate)
)
