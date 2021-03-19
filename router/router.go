package router

import (
	"github.com/gin-gonic/gin"
	"pdco/app/Controllers/HomeController"
)

func Router(router *gin.Engine) {
	router.GET("/", HomeController.Index)
}
