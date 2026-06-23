package utils

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

var (
		logWriter *bufio.Writer
		logMutex sync.Mutex
	)

func InitLogFile() error {

	logPath := "src/logs/errors.log"
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return  err
	}
	logWriter = bufio.NewWriter(file)
    return err
}

func LogError(err error) {
    if err == nil || logWriter == nil {
        return
    }

    logMutex.Lock()
    defer logMutex.Unlock()
    
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    fmt.Fprintf(logWriter, "[%s] %v\n", timestamp, err)
		logWriter.Flush()
}
