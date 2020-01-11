package validators

import (
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"time"
)

//验证器
func bookableDate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
			return false
		}
	}
	return true
}
//注册验证器
func test()  {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}
}
