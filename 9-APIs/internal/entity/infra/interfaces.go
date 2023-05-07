package database

import "github.com/Pedro-lms/goexpert/9-APIs/internal/entity"

type UserInference interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
