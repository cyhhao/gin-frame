package routers

import (
	"github.com/gin-gonic/gin"
	"gin-frame/controllers"

)


func SetupRouter(router *gin.Engine) *gin.Engine {



	v1 := router.Group("/v1")

	controllers.UserSetup(v1)



	return router
}
