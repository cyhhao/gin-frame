package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

import (
	_ "gin-frame/docs"
	"net/http"
	"gin-frame/routers"
	"github.com/gin-gonic/gin"
	"gin-frame/utils"
	"gin-frame/middleware"
)

// @title API Document 2
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
	//utils.SetupOrm()

	router := gin.New()
	router.Use(middleware.Logger(utils.Log),gin.Recovery())

	routers.SetupRouter(router)

	s := &http.Server{
		Addr:    ":" + utils.GetConfig("app::port").String(),
		Handler: router,
	}

	s.ListenAndServe()
}
