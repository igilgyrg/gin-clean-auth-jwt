package http

import (
	"context"
	"sync"
	validatorPkg "github.com/go-playground/validator/v10"
)

var once sync.Once
var validator *validatorPkg.Validate

func init(){
	once.Do(newValidator)
}

func newValidator() {
	validator = validatorPkg.New()
}

func ValidateStruct(ctx context.Context, s interface{}) error {
	return validator.StructCtx(ctx, s)
}