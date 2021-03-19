package HomeController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pdco/app/Models"
)

func Index(ctx *gin.Context) {
	check := new(Models.HealthCheckModel)
	_, _ = check.CheckHealth()
	ctx.JSON(http.StatusOK, check)
}
