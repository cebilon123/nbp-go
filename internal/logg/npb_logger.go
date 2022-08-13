package logg

import "errors"

// NbpLogger implements Logger interface
type NbpLogger struct {
	registeredLoggers []Logger
}

// NewNbpLogger creates new Nbp logger that logs
// into given logging targets
func NewNbpLogger(registeredLoggers []Logger) Logger {
	return NbpLogger{registeredLoggers: registeredLoggers}
}

func (n NbpLogger) LogStatus(data *LogData) error {
	if data == nil {
		return errors.New("NbpLogger.LogStatus: data pointer is NIL")
	}

	for _, logger := range n.registeredLoggers {
		// for now, we are going to ignore errors here
		_ = logger.LogStatus(data)
	}

	return nil
}
