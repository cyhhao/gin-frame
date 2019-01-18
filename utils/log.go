package utils

import "github.com/sirupsen/logrus"

var Log *logrus.Logger

func SetupLogging() {
	Log = logrus.New()
	Log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
}
