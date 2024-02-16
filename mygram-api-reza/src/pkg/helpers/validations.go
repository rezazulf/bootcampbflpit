package helpers

import (
	"final-project-h8-mygram-Reza/src/pkg/errs"

	"github.com/asaskevich/govalidator"
)

func ValidateStruct(payload interface{}) errs.MessageErr {
	_, err := govalidator.ValidateStruct(payload)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}

	return nil
}
