package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"timekeeper/lib/tkerr"
)

func Initialize(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.GET("/update/:username", updateAccount)
	v1.GET("/stats/:username", lookupStats)
	v1.GET("/track/:username", trackAccount)
}

// Return an HTTP status code and a JSON response for the given error.
func errorToResponse(err error) (int, gin.H) {
	var code int
	var ecode int
	var message string

	if e, ok := err.(*tkerr.TKError); ok {
		ecode = e.Code
		message = e.Message

		switch ecode {
		case tkerr.InvalidUsername, tkerr.RecentUpdate, tkerr.InvalidPeriod:
			code = http.StatusBadRequest
		case tkerr.InvalidAccount, tkerr.UntrackedAccount, tkerr.NoPointsInPeriod:
			code = http.StatusNotFound
		case tkerr.OSAPIError:
			code = http.StatusBadGateway
		default:
			code = http.StatusInternalServerError
		}
	} else {
		code = http.StatusInternalServerError
		ecode = -1
		message = "Internal error"
	}

	return code, gin.H{
		"status":       "error",
		"errorCode":    ecode,
		"errorMessage": message,
	}
}
