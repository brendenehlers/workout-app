package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type Fields = map[string]interface{}

func init() {
	log.Formatter = new(logrus.JSONFormatter)
	log.Out = os.Stdout
	log.Level = logrus.ErrorLevel
}

func SetDebug(debug bool) {
	if debug {
		log.Level = logrus.DebugLevel
	} else {
		log.Level = logrus.ErrorLevel
	}
}

func Debug(msg string, fields Fields) {
	log.WithFields(fields).Debug(msg)
}

func Error(msg string, fields Fields) {
	log.WithFields(fields).Error(msg)
}

func Err(err error, fields Fields) {
	log.WithError(err).WithFields(fields).Errorf(err.Error())
}
