package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code,omitempty"`
	Details string `json:"details,omitempty"`
}

func RespondWithError(c *gin.Context, httpStatus int, err error, clientMessage string) {
	c.Error(err)

	errMsg := clientMessage
	if errMsg == "" {
		errMsg = "An unexpected error occurred."
	}

	c.JSON(httpStatus, ErrorResponse{
		Message: errMsg,
	})
}

const InternalServerError = "Internal Server Error"

// RespondWithInternalError is a convenience function for 500 errors.
func RespondWithInternalError(c *gin.Context, err error) {
	RespondWithError(c, http.StatusInternalServerError, err, InternalServerError)
}

// RespondWithBadRequest is a convenience function for 400 errors.
func RespondWithBadRequest(c *gin.Context, err error, clientMessage string) {
	RespondWithError(c, http.StatusBadRequest, err, clientMessage)
}
