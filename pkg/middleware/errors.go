package middleware

import (
	"net/http"
	"todolist/pkg/errors"

	"github.com/gin-gonic/gin"
)

// ErrorHandler is a middleware that handles errors raised by the handlers.
// If there are errors in the context, it marshals them into JSON and returns
// them with a status code.
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			// Create a map to store the errors. The keys are names of the fields
			// that contain errors and the values are error messages.
			errorsMap := make(map[string]string)

			// Determine the status code to return based on the errors.
			// If there are multiple errors, the status code of the first error
			// is used.
			var statusCode int

			// Iterate over the errors and add them to the map.
			for _, err := range c.Errors {
				switch err.Err {
				case errors.ErrInvalidRequest:
					statusCode = http.StatusBadRequest
					errorsMap["request"] = "Invalid request format"
				case errors.ErrTaskNotFound:
					statusCode = http.StatusNotFound
					errorsMap["task"] = "Task not found"
				case errors.ErrInvalidTaskID:
					statusCode = http.StatusBadRequest
					errorsMap["id"] = "Invalid task ID"
				default:
					statusCode = http.StatusInternalServerError
					errorsMap["internal"] = "Internal server error"
				}
			}

			// Marshal the errors map to JSON and write it to the response.
			c.JSON(statusCode, gin.H{"errors": errorsMap})
		}
	}
}
