package utils

import (
	"sync"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate
var once sync.Once

func GetValidator() *validator.Validate {
	once.Do(func() {
		validate = validator.New()
	})
	return validate
}
