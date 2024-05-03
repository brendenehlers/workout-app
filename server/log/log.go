package log

import (
	"log"
)

func Debug(v ...interface{}) {
	log.Printf("[DEBUG] %v\n", v...)
}

func Info(v ...interface{}) {
	log.Printf("[INFO] %v\n", v...)
}

func Warn(v ...interface{}) {
	log.Printf("[WARN] %v\n", v...)
}

func Err(err error) {
	log.Printf("[ERROR] %v\n", err)
}

func Error(v ...interface{}) {
	log.Printf("[ERROR] %v\n", v...)
}

func Fatal(v ...interface{}) {
	log.Printf("[FATAL] %v\n", v...)
}
