package apiv1

import (
	"github.com/gin-gonic/gin"
	"timekeeper/models"
)

func updateAccount(c *gin.Context) {
	username := c.Query("username")

	if err := models.UpdateAccount(username); err != nil {
		code, json := handleAccountError(err)
		c.JSON(code, json)
	} else {
		c.JSON(200, gin.H{"status": "success"})
	}
}

func handleAccountError(err error) (int, gin.H) {
	var code int
	var message string

	if e, ok := err.(*models.AccountError); ok {
		switch e.Type {
		case models.AEInvalidUser:
			code = 422
		default:
			code = 500
		}
		message = e.Message
	} else {
		code = 500
		message = "Internal error"
	}

	return code, gin.H{"status": "error", "error": message}
}
