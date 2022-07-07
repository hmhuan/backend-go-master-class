package util

import "github.com/go-playground/validator/v10"

var ValidCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check currency supported
		return isSupportedCurrency(currency)
	}
	return false
}
