package logger

import (
	"fmt"
	"os"
	"time"
)

func WriteLog(message string) error {
	file, err := os.OpenFile("loggger/log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer file.Close()

	logMessage := fmt.Sprintf("%s - %s\n", time.Now().Format("2006-01-02 15:04:05"), message)

	_, err = file.WriteString(logMessage)
	if err != nil {
		return fmt.Errorf("failed to write to log file: %w", err)
	}

	// Ensure data is written before closing
	err = file.Sync()
	if err != nil {
		return fmt.Errorf("failed to flush log data: %w", err)
	}

	return nil
}
