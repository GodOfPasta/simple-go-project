package validators

import (
	"reflect"

	"time"

	"github.com/go-playground/validator/v10"
)

func DeadlineValidator(fl validator.FieldLevel) bool {
	var minDate = time.Now()
	switch v := fl.Field(); v.Kind() {
	case reflect.Struct:
		val := v.Interface().(time.Time)
		if val.After(minDate) {
			return true
		}
	default:
		return false
	}

	return false
}
