package apiv1

import (
	"github.com/gin-gonic/gin"
	"timekeeper/tk"
)

const (
	TAInvalidPeriod = iota
)

func trackAccount(c *gin.Context) {
	username := c.Param("username")
	period, err := tk.ParsePeriod(c.DefaultQuery("period", "7d"))
	if err != nil {
		c.JSON(400, gin.H{
			"status":       "error",
			"errorCode":    TAInvalidPeriod,
			"errorMessage": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"username": username, "period": period})
}
