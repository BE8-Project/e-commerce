package response

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type LoginDetail struct {
	User  Login  `json:"user"`
	Token string `json:"token"`
}

func StatusInvalidRequest() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "invalid request",
		"data":    nil,
	}
}

func StatusBadRequest(err error) map[string]interface{} {
	var field, tag string
	var message [] string
	
	for _, err := range err.(validator.ValidationErrors) {
		field = fmt.Sprint(err.StructField())
		tag = fmt.Sprint(err.Tag())

		message = append(message, "field "+field+" : "+tag)
	}
	
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": message,
		"data":    nil,
	}
}

func StatusUnauthorized(err error) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusUnauthorized,
		"message": err.Error(),
		"data":    nil,
	}
}

func StatusForbidden(message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusForbidden,
		"message": message,
		"data":    nil,
	}
}

func StatusNotFound(message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusNotFound,
		"message": message,
		"data":    nil,
	}
}

func StatusOK(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": message,
		"data":    data,
	}
}

func StatusCreated(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": message,
		"data":    data,
	}
}
