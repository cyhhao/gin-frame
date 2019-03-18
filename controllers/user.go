package controllers

import (
	"github.com/gin-gonic/gin"
	"gin-frame/models"
	. "gin-frame/utils"
	"net/http"
			)

func UserSetup(v1 *gin.RouterGroup) {
	r := v1.Group("/user")

	r.GET("login", login)
	r.POST("logout", logout)

}

// @Tags user
// @Summary 登录
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param name query string true "姓名备注"
// @Router /user/login [get]
func login(c *gin.Context) {
	user := models.Users{Name: "cyh", Age: 18, Email: "abcdefg@qq.com"}
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

	searchUser.Age = 24
	err = Orm.Save(&searchUser).Error
	if err != nil {
		Log.Errorln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"user": searchUser,
	})
}

type Info struct {
	Name string `json:"name1"`
	Age  string
	Pig  int `json:"pig"`
}


// @Tags user
// @Summary 登出
// @Accept json
// @Produce  json
// @Param name  	formData 	string	true 	"test"
// @Param age  		formData 	int		true 	"test"
// @Param title  	formData 	string	false 	"可选"
// @Router /user/logout [post]
func logout(c *gin.Context) {


	c.JSON(http.StatusOK, gin.H{
		"user": 123,
	})
}
