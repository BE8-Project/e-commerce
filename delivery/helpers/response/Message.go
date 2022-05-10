package response

import (
	"net/http"
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
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": err.Error(),
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
