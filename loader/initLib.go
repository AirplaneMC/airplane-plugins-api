package loader

import (
	"github.com/AirplaneMC/airplane-plugins-api/controller/library"
	"github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

func initLibs(l *lua.LState) {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.DebugLevel
	logger := library.Logger{Send: log}

	l.PreloadModule("log", logger.InitLogLIB)
}
