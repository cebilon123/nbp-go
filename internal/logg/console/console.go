package console

import (
	"github.com/cebilon123/nbp-go/internal/logg"
	"github.com/cebilon123/nbp-go/internal/logg/log_format"
	"log"
)

// consoleLogger logs to console and implements Logger interface
type consoleLogger struct {
}

func NewConsoleLogger() logg.Logger {
	return &consoleLogger{}
}

func (c consoleLogger) LogStatus(data *logg.LogData) error {
	log.Println(log_format.GetString(data))
	return nil
}
