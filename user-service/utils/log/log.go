package utils

import (
	"fmt"
	"runtime"

	"github.com/sirupsen/logrus"
)

func LogError(err error) {
	_, fn, line, _ := runtime.Caller(1)
	logrus.WithFields(logrus.Fields{
		"event": fmt.Sprintf("[error] %s:%d", fn, line),
	}).Error(err)
}
