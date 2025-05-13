package logger

import (
	"log"
	"os"
)

var logFile *os.File

func Init(logPath string) error {
	var err error
	logFile, err = os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	log.SetOutput(logFile)

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	return nil
}

func Close() {
	if logFile != nil {
		logFile.Close()
	}
}

func LogToFileAsync(msg string) {
	go func() {
		log.Println(msg)
	}()
}
