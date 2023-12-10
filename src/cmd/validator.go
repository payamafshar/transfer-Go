package cmd

import (
	"ReservApp/src/pkg/helper"

	"github.com/go-playground/validator/v10"
)

var ValidCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		//checking currency is supported or not
		return helper.IsSupportedCurrency(currency)
	}
	return false
}
