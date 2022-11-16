package logger

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type Logger struct {
	LogUtil *logrus.Logger
}

func GetNewLogger() *Logger {
	return &Logger{LogUtil: logrus.New()}
}

func (l *Logger) SetLogLevel(level logrus.Level) {
	l.LogUtil.SetLevel(level)
}

// AssignOutput creates (or checks if dir exists?) a new output file for the logger and assigns it to the logger.
//
// Logs will be in under /app/logs/whatevertheservice.
//
// E.g logDirectory = ./logs/dbconnections, logCategory = dbLog
//
// or logDirectory = ./logs/charactercreation, logCategory = characterCreationLog
func (l *Logger) AssignOutput(logCategory string, logDirectory string) error {
	// Check if logDirectory exists, if not create it.
	_, err := os.Stat(logDirectory)
	if os.IsNotExist(err) {
		err := os.Mkdir(logDirectory, os.ModePerm)
		if err != nil {
			return err
		}
	}
	logFileName := fmt.Sprintf("%s+%s", logCategory, time.Now().UTC().Format("2006-01-02")+".log")
	logOutput := fmt.Sprintf("%s/%s", logDirectory, logFileName)
	var logFile *os.File
	var errRet error
	if _, err := os.Stat(logOutput); errors.Is(err, os.ErrNotExist) {
		logFile, err = os.OpenFile(logOutput, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			errRet = err
		}
	} else {
		logFile, err = os.OpenFile(logOutput, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			errRet = err
		}
	}
	l.LogUtil.SetOutput(logFile)
	return errRet
}
