package helpers

import (
	"github.com/gin-gonic/gin"
	"gin-frame/dto"
)

func GetHeaderInfos(c *gin.Context) (dto.HeaderInfo) {

	headerInfo := dto.HeaderInfo{}
	headerInfo.ClientIP = c.ClientIP()

	//if IsDevTest() {
	//	headerInfo.ClientIP, _ = Get("http://ipv4.icanhazip.com", nil)
	//	headerInfo.ClientIP = strings.Replace(headerInfo.ClientIP, " ", "", -1)
	//	headerInfo.ClientIP = strings.Replace(headerInfo.ClientIP, "\n", "", -1)
	//}

	headerInfo.AppVersion = c.GetHeader("appVersion")
	if headerInfo.AppVersion == "" {
		headerInfo.AppVersion = "1.0.0"
	}
	headerInfo.PhoneSystem = c.GetHeader("phoneSystem")
	if headerInfo.PhoneSystem == "" {
		headerInfo.PhoneSystem = "iOS"
	}
	headerInfo.DeviceID = c.GetHeader("deviceID")
	return headerInfo
}
