package main

import (
	"net/http"

	"github.com/gin-blog/pkg/setting"
	"github.com/gin-blog/routers"
)

func main() {
	// router := gin.Default()
	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": "OK",
	// 	})
	// })

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
