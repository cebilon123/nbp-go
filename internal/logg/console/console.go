package console

import (
	"github.com/cebilon123/nbp-go/internal/logg"
	"log"
)

// consoleWriter logs to console and implements logg.ConcurrentWriter interface
type consoleWriter struct {
}

func (c consoleWriter) WriteString(s string) (n int, err error) {
	if len(s) == 0 {
		return 0, nil
	}
	log.Print(s)
	return len(s), nil
}

// Close it is not used here, but must be implemented
// because of logg.ConcurrentWriter interface
func (c consoleWriter) Close() error {
	return nil
}

func NewConsoleWriter() logg.ConcurrentWriter {
	return consoleWriter{}
}
