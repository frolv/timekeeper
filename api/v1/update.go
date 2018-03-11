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
		code, json := errorToResponse(err)
		c.JSON(code, json)
	} else {
		cache.InvalidateAccount(acc)
		c.JSON(200, gin.H{"status": "success"})
	}
}
