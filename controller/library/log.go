package library

import (
	"github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

type Logger struct {
	Send *logrus.Logger
}

func (l *Logger) info(s *lua.LState) int {
	msg := s.ToString(1)

	l.Send.Infoln(msg)

	return 0
}

func (l *Logger) warn(s *lua.LState) int {
	msg := s.ToString(1)

	l.Send.Warnln(msg)

	return 0
}

func (l *Logger) debug(s *lua.LState) int {
	msg := s.ToString(1)

	l.Send.Debugln(msg)

	return 0
}

func (log *Logger) InitLogLIB(l *lua.LState) int {
	var exports = map[string]lua.LGFunction{ // Library Functions
		"Info":  log.info,
		"Debug": log.debug,
		"Warn":  log.warn,
	}

	t := l.SetFuncs(l.NewTable(), exports) // Initialize functions

	l.Push(t) // Returning functions

	return 1
}
