package main

import (
	"github.com/gin-gonic/gin"
	"pdco/router"
)

func main() {
	r := gin.Default()
	router.Router(r)
	_ = r.Run(":5000")
}
