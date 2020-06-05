package validator

import "github.com/asaskevich/govalidator"

var MustCheck = false

func Init() {
	MustCheck = true
}

func Struct(payload interface{}) error {
	if MustCheck {
		//validate valid tag
		if _, err := govalidator.ValidateStruct(payload); err != nil {
			return err
		}
	}
	return nil
}
