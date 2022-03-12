package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
)

var Logger = NewLogger()

var (
	output = &lumberjack.Logger{
		Filename:   "logs/logs.log",
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     30,
	}
	level      = log.DebugLevel
	timeFormat = "2006-01-02 15:04:05"
)

type format struct {
	log.TextFormatter
}

func (f *format) Format(entry *log.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("[%s] - %s - %s\n", entry.Time.Format(f.TimestampFormat), strings.ToUpper(entry.Level.String()), entry.Message)), nil
}

func NewLogger() *log.Logger {
	logger := log.Logger{}
	logger.SetOutput(os.Stdout)
	logger.SetLevel(level)
	logger.SetFormatter(&format{
		log.TextFormatter{
			FullTimestamp:          true,
			TimestampFormat:        timeFormat,
			DisableLevelTruncation: true,
		},
	})
	return &logger
}
