package validator

import (
	v "github.com/go-playground/validator/v10"
	optsGenValidator "github.com/kazhuravlev/options-gen/pkg/validator"
)

var Validator = v.New()

func init() {
	optsGenValidator.Set(Validator)
}
