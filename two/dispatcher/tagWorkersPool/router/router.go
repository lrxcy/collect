package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jimweng/dispatcher/tagWorkersPool/controller"
)

func ApiRouter(r *gin.Engine) {
	router := r.Group("/")
	router.GET("g1", controller.Collector)
}
