package errors

import "errors"

var (
	ErrInvalidProduct = errors.New("producto no encontrado")
	ErrNotStock       = errors.New("no hay stock suficiente")
)
