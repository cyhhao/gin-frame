package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(data interface{}) (int, gin.H) {
	return http.StatusOK, gin.H{
		"code": 0,
		"data": data,
		"msg":  ErrorCode[0],
	}
}

func Failure(code int) (int, gin.H) {
	return http.StatusOK, gin.H{
		"code": code,
		"data": nil,
		"msg":  ErrorCode[code],
	}
}

func FailureWithError(code int, err error, data interface{}) (int, gin.H) {
	if err != nil {
		return http.StatusOK, gin.H{
			"code": code,
			"msg":  err.Error(),
			"data": data,
		}
	} else {
		return http.StatusOK, gin.H{
			"code": code,
			"msg":  ErrorCode[code],
			"data": data,
		}
	}
}


func ResErrorData(c *gin.Context,data interface{},err error){
	if err == nil {
		c.JSON(Success(data))
	} else {
		c.JSON(FailureWithError(4000, err, data))
	}
}