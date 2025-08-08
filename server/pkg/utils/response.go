package utils

import (
	"net/http"

	"github.com/francis/projectx-api/internal/model"
	"github.com/gin-gonic/gin"
)

// SendSuccess sends a successful response
func SendSuccess(c *gin.Context, statusCode int, data interface{}, message string) {
    c.JSON(statusCode, model.SuccessResponse(data, message))
}

// SendError sends an error response
func SendError(c *gin.Context, statusCode int, message string) {
    c.JSON(statusCode, model.ErrorResponse(message))
}

// SendValidationError sends a validation error response
func SendValidationError(c *gin.Context, err error) {
    c.JSON(http.StatusBadRequest, model.ErrorResponse(err.Error()))
}

// SendInternalError sends an internal server error response
func SendInternalError(c *gin.Context, message string) {
    if message == "" {
        message = "Internal server error"
    }
    c.JSON(http.StatusInternalServerError, model.ErrorResponse(message))
}

// SendUnauthorized sends an unauthorized response
func SendUnauthorized(c *gin.Context, message string) {
    if message == "" {
        message = "Unauthorized"
    }
    c.JSON(http.StatusUnauthorized, model.ErrorResponse(message))
}

// SendForbidden sends a forbidden response
func SendForbidden(c *gin.Context, message string) {
    if message == "" {
        message = "Forbidden"
    }
    c.JSON(http.StatusForbidden, model.ErrorResponse(message))
}

// SendNotFound sends a not found response
func SendNotFound(c *gin.Context, message string) {
    if message == "" {
        message = "Resource not found"
    }
    c.JSON(http.StatusNotFound, model.ErrorResponse(message))
}

// SendBadRequest sends a bad request response
func SendBadRequest(c *gin.Context, message string) {
    if message == "" {
        message = "Bad request"
    }
    c.JSON(http.StatusBadRequest, model.ErrorResponse(message))
}