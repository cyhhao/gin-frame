package controllers

import (
	"github.com/gin-gonic/gin"
	"gin-frame/models"
	"database/sql"
	. "gin-frame/utils"
)

func UserSetup(v1 *gin.RouterGroup) {
	r := v1.Group("/user")

	r.GET("login", login)
	r.POST("logout", logout)

}

func login(c *gin.Context) {
	user := models.Users{Name: "cyh", Age: sql.NullInt64{Int64: 18, Valid: true}, Email: "vvv@qqq.com"}
	err := Orm.Create(&user).Error
	if err != nil {
		Log.Errorln(err)
	}

	var searchUser models.Users
	err = Orm.Where("name = ?", "cyh").First(&searchUser).Error
	if err != nil {
		Log.Errorln(err)
	} else {
		Log.WithField("user", searchUser).Info("test")
	}

	searchUser.Age.Int64 = 24
	err = Orm.Save(&searchUser).Error
	if err != nil {
		Log.Errorln(err)
	}

}

func logout(c *gin.Context) {

}
