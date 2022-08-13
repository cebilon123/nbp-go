package file

import (
	"context"
	"fmt"
	"github.com/cebilon123/nbp-go/internal/logg"
	"github.com/cebilon123/nbp-go/internal/logg/log_format"
	"log"
	"os"
)

type fileLogger struct {
	loggChan chan *logg.LogData
	ctx      context.Context
}

func NewFileLogger(ctx context.Context) logg.Logger {
	logChan := make(chan *logg.LogData)
	fl := fileLogger{ctx: ctx, loggChan: logChan}
	go fl.startFileLogging()
	return fl
}

func (f fileLogger) LogStatus(data *logg.LogData) error {
	f.loggChan <- data
	return nil
}

func (f fileLogger) startFileLogging() {
	// TODO edit it so the Writer is going to be passed as interface (to enable unit testing of the thing.)
	file, err := os.OpenFile("./log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println("FileLogger: ", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("FileLogger: ", err)
		}
	}(file)

	for {
		select {
		case <-f.ctx.Done():
			return
		case d := <-f.loggChan:
			if _, err = file.WriteString(fmt.Sprintln(log_format.GetString(d))); err != nil {
				log.Println("FileLogger: ", err)
			}
		}
	}
}
