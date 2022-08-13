package logg

import (
	"testing"
)

type loggerMock struct {
	Logged *bool
}

func (l *loggerMock) LogStatus(data *LogData) error {
	*l.Logged = true
	return nil
}

func TestNbpLogger_LogStatus_DataIsNil_ReturnsError(t *testing.T) {
	t.Run("TestNbpLogger_LogStatus_DataIsNil_ReturnsError", func(t *testing.T) {
		n := NewNbpLogger([]Logger{})

		if err := n.LogStatus(nil); err == nil {
			t.Error("LogStatus() error. Should return error if data is NIL")
		}
	})
}

func TestNbpLogger_LogStatus_RegisteredLoggersShouldLog(t *testing.T) {
	t.Run("TestNbpLogger_LogStatus_RegisteredLoggersShouldLog", func(t *testing.T) {
		logged := false
		lm := &loggerMock{Logged: &logged}
		n := NewNbpLogger([]Logger{lm})

		if err := n.LogStatus(&LogData{}); err != nil {
			t.Errorf("LogStatus() error. %s", err)
		}

		if !logged {
			t.Error("LogStatus() error. NbpLogger should log to registered logger")
		}
	})
}
