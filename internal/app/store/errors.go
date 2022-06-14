package store

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrNoUser         = errors.New("no user")
	ErrBadPassword    = errors.New("bad password")
)
