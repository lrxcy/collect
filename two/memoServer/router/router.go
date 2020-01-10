package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jimweng/memoServer/controllers"
)

type Router struct {
	router *gin.Engine
}

func (r *Router) Run() {
	r.router.Run()
}

type RouterImpl interface {
	Run()
}

func NewRouter() RouterImpl {
	var rr Router
	r := gin.Default()

	// add CORS allow header
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("*")

	r.Use(cors.New(corsConfig))

	r.GET("/v1", controllers.ReturnPageInfo)
	r.DELETE("/v1", controllers.DeleteSpecificValue)
	r.POST("/v1", controllers.PostData)
	r.PUT("/v1", controllers.UpdateData)

	rr.router = r
	return &rr
}
