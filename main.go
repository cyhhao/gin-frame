package main

import (
	"net/http"
	"gin-frame/routers"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gin-frame/utils"
	"gin-frame/middleware"
)

func main() {
	// 初始化日志
	utils.SetupLogging()
	// 初始化配置
	utils.SetupConfig()
	// 初始化 ORM
	utils.SetupOrm()

	router := gin.New()
	router.Use(middleware.Logger(utils.Log),gin.Recovery())

	routers.SetupRouter(router)

	s := &http.Server{
		Addr:    ":" + utils.GetConfig("app::port").String(),
		Handler: router,
	}

	s.ListenAndServe()
}
