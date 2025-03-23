package logger

import (
	"log"
)

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Criticalf(format string, args ...interface{}) {
	log.Printf("[CRITICAL] "+format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	log.Printf("[ERROR] "+format, args...)
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	log.Printf("[WARNING] "+format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	log.Printf("[INFO] "+format, args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	log.Printf("[DEBUG] "+format, args...)
}
