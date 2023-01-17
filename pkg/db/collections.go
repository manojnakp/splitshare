package db

import "errors"

var (
	ErrNotFound    = errors.New("record not found")
	ErrInvalidType = errors.New("invalid record type")
)

type Collection interface {
	Insert(doc any) error
	Find(id string) (any, error)
	Update(id string, doc any) error
	Delete(id string) error
}
