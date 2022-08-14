package concurrent_logger

import (
	"context"
	"fmt"
	"github.com/cebilon123/nbp-go/internal/logg"
	"github.com/cebilon123/nbp-go/internal/logg/log_format"
	"log"
)

type concurrentLogger struct {
	loggChan chan *logg.LogData
	ctx      context.Context
}

func New(ctx context.Context, concurrentWriter logg.ConcurrentWriter) logg.Logger {
	logChan := make(chan *logg.LogData)
	fl := concurrentLogger{ctx: ctx, loggChan: logChan}
	go fl.startConcurrentLogging(concurrentWriter)
	return fl
}

func (cl concurrentLogger) LogStatus(data *logg.LogData) error {
	cl.loggChan <- data
	return nil
}

func (cl concurrentLogger) startConcurrentLogging(concurrentWriter logg.ConcurrentWriter) {
	defer func(fileWriter logg.ConcurrentWriter) {
		err := fileWriter.Close()
		if err != nil {
			log.Println("ConcurrentLogger: ", err)
		}
	}(concurrentWriter)

	for {
		select {
		case <-cl.ctx.Done():
			return
		case d := <-cl.loggChan:
			if _, err := concurrentWriter.WriteString(fmt.Sprintln(log_format.GetString(d))); err != nil {
				log.Println("ConcurrentLogger: ", err)
			}
		}
	}
}
