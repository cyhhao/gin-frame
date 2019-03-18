package middlewares

import (
	"github.com/gin-gonic/gin"
	"gin-frame/utils/helpers"
	. "gin-frame/utils"
	"gin-frame/utils/res"
	"gin-frame/models"
	"gin-frame/middlewares/jwtauth"
)

func AuthMiddleware(force ... bool) gin.HandlerFunc {
	var isForce = false
	if len(force) > 0 && force[0] {
		isForce = true
	}
	return func(c *gin.Context) {
		headerInfo := helpers.GetHeaderInfos(c)
		Log.Debug(headerInfo)
		userTokenStr := c.GetHeader("userToken")

		et := jwtauth.EasyToken{}
		validate, userId, _ := et.ValidateToken(userTokenStr)
		if !validate {
			if isForce {
				c.JSON(res.Failure(40102))
				c.Abort()
			} else {
				c.Set("userId", "")
				c.Next()
			}
			return

		}
		var sessionToken models.SessionToken
		err := Orm.First(&sessionToken, "uuid = ?", userId).Error
		if err != nil || sessionToken.Token != userTokenStr {
			if isForce {
				c.JSON(res.Failure(40101))
				c.Abort()
			} else {
				c.Set("userId", "")
				c.Next()
			}
			return
		}
		c.Set("userId", userId)
		c.Next()

	}
}




