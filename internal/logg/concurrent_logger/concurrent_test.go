package concurrent_logger

import (
	"context"
	"github.com/cebilon123/nbp-go/internal/logg"
	"testing"
)

type concurrentWriterMock struct {
	messagesCount *int
}

func (c concurrentWriterMock) WriteString(s string) (n int, err error) {
	*c.messagesCount += 1
	return 1, nil
}

func (c concurrentWriterMock) Close() error {
	return nil
}

func Test_concurrentLogger_startConcurrentLogging_LogsSuccessfully(t *testing.T) {
	t.Run("Test_concurrentLogger_startConcurrentLogging_LogsSuccessfully", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())

		messagesSentCount := 0

		mock := concurrentWriterMock{
			messagesCount: &messagesSentCount,
		}

		cl := concurrentLogger{
			loggChan: make(chan *logg.LogData),
			ctx:      ctx,
		}

		go cl.startConcurrentLogging(mock)

		_ = cl.LogStatus(&logg.LogData{
			Message: "1",
		})
		_ = cl.LogStatus(&logg.LogData{
			Message: "2",
		})

		cancel()

		if messagesSentCount != 2 {
			t.Errorf("startConcurrentLogging() error. Expected sent count: %d, got: %d.",
				2, messagesSentCount)
		}
	})
}
