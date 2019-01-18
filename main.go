package main

import (
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

import (
	_ "gin-frame/docs"
	"net/http"
	"gin-frame/routers"
	"github.com/gin-gonic/gin"
	"gin-frame/utils"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io


// @BasePath /v1
func main() {
	// 初始化日志
	utils.SetupLogging()
	// 初始化配置
	utils.SetupConfig()
	// 初始化 ORM
	utils.SetupOrm()

	router := gin.Default()

	// 初始化 swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routers.SetupRouter(router)

	s := &http.Server{
		Addr:    ":" + utils.GetConfig("app::port").String(),
		Handler: router,
	}

	s.ListenAndServe()
}
