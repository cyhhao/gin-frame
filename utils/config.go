package utils

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/go-ini/ini"
	"os"
	"strings"
)

var (
	EnvMode string
	Config *ini.File
)

func SetupConfig() {
	envParam := flag.String("env", "dev", "--env dev/test/prod")
	flag.Parse()

	EnvMode = *envParam

	if EnvMode == "prod" {
		gin.SetMode(gin.ReleaseMode)
		Log.SetLevel(logrus.InfoLevel)
	} else if EnvMode == "test" {
		gin.SetMode(gin.TestMode)
		Log.SetLevel(logrus.InfoLevel)
	} else {
		gin.SetMode(gin.DebugMode)
		Log.SetLevel(logrus.DebugLevel)
	}

	Log.Info("EnvMode:", EnvMode)

	var err error
	Config, err = ini.Load("conf/app.ini", "conf/"+EnvMode+".ini")
	if err != nil {
		Log.Panicln(err)
		return
	}

	// 加入环境变量
	Config.ValueMapper = os.ExpandEnv

}
func GetConfig(key string) *ini.Key {
	parts := strings.Split(key, "::")
	section := parts[0]
	keyStr := parts[1]
	return Config.Section(section).Key(keyStr)

}
