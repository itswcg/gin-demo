package routers

import (
	"gin-demo/pkg/setting"
	"gin-demo/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	api_v1 := r.Group("/api/v1")
	{
		api_v1.GET("/tags", v1.GetTags)
		api_v1.POST("/tags", v1.AddTag)
		api_v1.PUT("/tags/:id", v1.EditTag)
		api_v1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
