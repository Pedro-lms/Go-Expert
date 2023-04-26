package entity

import (
	"errors"

	"github.com/Pedro-lms/goexpert/9-APIs/pkg/entity"
)

var (
	ErrIDsRequired = errors.New("ids required")
	ErrInvalidID = errors.New("invalid id")
	ErrNameIsRequired = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice = errors.New("invalid price")
)

type Product struct 