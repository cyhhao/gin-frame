package routers

import (
	"github.com/gin-gonic/gin"
	"gin-frame/controllers"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)


func SetupRouter(router *gin.Engine) *gin.Engine {



	v1 := router.Group("/v1")

	controllers.UserSetup(v1)

	// 初始化 swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
