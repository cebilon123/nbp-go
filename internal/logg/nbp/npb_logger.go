package nbp

import (
	"errors"
	"github.com/cebilon123/nbp-go/internal/logg"
	"github.com/cebilon123/nbp-go/internal/logg/console"
)

// LoggerAggregator implements Logger interface
type LoggerAggregator struct {
	registeredLoggers []logg.Logger
}

// NewLoggerAggregator creates new Nbp logger that logs
// into given logging targets
func NewLoggerAggregator(registeredLoggers []logg.Logger) logg.Logger {
	return LoggerAggregator{registeredLoggers: registeredLoggers}
}

func NewDefaultLoggerAggregator() logg.Logger {
	return LoggerAggregator{registeredLoggers: []logg.Logger{console.NewConsoleLogger()}}
}

func (n LoggerAggregator) LogStatus(data *logg.LogData) error {
	if data == nil {
		return errors.New("LoggerAggregator.LogStatus: data pointer is NIL")
	}

	for _, logger := range n.registeredLoggers {
		// for now, we are going to ignore errors here
		_ = logger.LogStatus(data)
	}

	return nil
}
