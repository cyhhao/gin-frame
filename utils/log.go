package utils

import (
	"github.com/sirupsen/logrus"

	"runtime"
	"path"
	"strings"
	"fmt"
)

var Log *logrus.Logger

type ContextHook struct {
}

func (hook ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook ContextHook) Fire(entry *logrus.Entry) error {
	pc := make([]uintptr, 6, 6)
	cnt := runtime.Callers(6, pc)
	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i] - 1)
		name := fu.Name()
		if !strings.Contains(name, "github.com/sirupsen/logrus") {
			file, line := fu.FileLine(pc[i] - 1)
			entry.Data["file"] = path.Base(file) + ":" + fmt.Sprintf("%d", line)
			entry.Data["func"] = path.Base(name)
			break
		}
	}
	return nil
}

func SetupLogging() {
	Log = logrus.New()
	Log.AddHook(ContextHook{})
	Log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
}

