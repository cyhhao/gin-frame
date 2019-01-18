package main

import (
	"net/http"
	"gin-frame/routers"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gin-frame/utils"
)

func main() {
	utils.SetupLogging()
	utils.SetupConfig()
	utils.SetupOrm()

	router := gin.Default()
	routers.SetupRouter(router)
	s := &http.Server{
		Addr:    ":" + utils.GetConfig("app::port").String(),
		Handler: router,
	}

	s.ListenAndServe()
}
