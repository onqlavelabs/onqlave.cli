package utils

import (
	"github.com/go-playground/validator"
)

type (
	ServiceValidator struct {
		validator *validator.Validate
	}
)

func (cv *ServiceValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return err
		//return echo.NewHTTPError(http.StatusBadRequest, common.NewBaseErrorResponse(int32(http.StatusBadRequest), err.Error()))
	}
	return nil
}

func NewServiceValidator() *ServiceValidator {
	return &ServiceValidator{validator: validator.New()}
}
