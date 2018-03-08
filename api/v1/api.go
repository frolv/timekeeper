package apiv1

import "github.com/gin-gonic/gin"

func Initialize(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.GET("/update/:username", updateAccount)
	v1.GET("/stats/:username", lookupStats)
}
