package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

var L *log.Logger

func InitLogger(logDir, logLevel, format string) error {
	stat, err := os.Stat(logDir)
	if err != nil && os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "log-dir %s not exists", logDir)
		return err
	}

	if !stat.IsDir() {
		fmt.Fprintf(os.Stderr, "log-dir %s is not a valid directory", logDir)
		return errors.New("log-dir is not a directory")
	}

	// log initialization
	logPath := filepath.Join(logDir, "notifier.log")
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}

	parsedLevel, err := log.ParseLevel(logLevel)
	if err != nil {
		parsedLevel = log.DebugLevel
	}

	L = &log.Logger{
		Out:       logFile,
		Formatter: new(log.JSONFormatter),
		Level:     parsedLevel,
	}

	L.SetReportCaller(true)

	return nil
}
