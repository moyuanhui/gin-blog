package routers

import (
	v1 "gin-blog/apis/v1"
	"gin-blog/pkg/setting"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	appv1 := r.Group("api/v1")
	{
		appv1.GET("/tags", v1.GetTags)
		appv1.POST("/tags", v1.AddTag)
		appv1.PUT("/tags/:id", v1.EditTag)
		appv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
