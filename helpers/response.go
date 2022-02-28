package helpers

import (
	"final-project/structs"

	"github.com/beego/beego/v2/core/validation"
)

// Default Response API
func DefResponse(data interface{}, error interface{}, code int) (response structs.Response) {
	response.Data = data
	response.Error = error
	response.Code = code
	return response
}

// Default Error Response API
func ErrResponse(data []string) (response structs.ErrResponse) {
	for i := 0; i < len(data); i++ {
		response.Data = append(response.Data, data[i])
	}
	return response
}

func ErrValidation(data []*validation.Error) (response []string) {
	for i := 0; i < len(data); i++ {
		response = append(response, data[i].String())
	}
	return response
}
