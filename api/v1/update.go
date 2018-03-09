package apiv1

import (
	"github.com/gin-gonic/gin"
	"timekeeper/lib/cache"
	"timekeeper/tk"
)

func updateAccount(c *gin.Context) {
	username := c.Param("username")

	acc, err := tk.UpdateAccount(username)
	if err != nil {
		code, json := handleAccountError(err)
		c.JSON(code, json)
	} else {
		cache.InvalidateAccount(acc)
		c.JSON(200, gin.H{"status": "success"})
	}
}

func handleAccountError(err error) (int, gin.H) {
	var code int
	var ecode int
	var message string

	if e, ok := err.(*tk.UpdateError); ok {
		ecode = e.Type
		switch ecode {
		case tk.UEInvalidUsername, tk.UERecentUpdate:
			code = 400
		default:
			code = 500
		}
		message = e.Message
	} else {
		code = 500
		ecode = -1
		message = "Internal error"
	}

	return code, gin.H{
		"status":       "error",
		"errorCode":    ecode,
		"errorMessage": message,
	}
}
