package apiv1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"timekeeper/tk"
)

func suggest(c *gin.Context) {
	query := c.Query("q")
	if len(query) == 0 {
		c.JSON(400, gin.H{
			"status": "error",
			"error": "missing required param 'q'",
		})
		return
	}

	n := c.DefaultQuery("n", "50")
	num, err := strconv.Atoi(n)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "error",
			"error": "invalid number",
		})
		return
	}

	res := tk.AutoSuggest(query, num)
	c.JSON(200, res)
}
