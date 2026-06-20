package utils

import (
    "fmt"
    "os"
    "time"
)

var logFile *os.File

func InitLogFile() error {

	logPath := "src/logs/errors.log"
    var err error
    logFile, err = os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    return err
}

func LogError(err error) {
    if err == nil || logFile == nil {
        return
    }

    timestamp := time.Now().Format("2006-01-02 15:04:05")
    logEntry := fmt.Sprintf("[%s] %v\n", timestamp, err)
    
    logFile.WriteString(logEntry)
}
