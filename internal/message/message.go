package message

import "errors"

// Common errors
var (
	ErrCreate = errors.New("create")
	ErrUpdate = errors.New("update")
	ErrGet    = errors.New("get")
	ErrDelete = errors.New("delete")

	ErrRequired = errors.New("required")
	ErrInvalid  = errors.New("invalid")

	ErrNotFound = errors.New("not found")
)
