package checker

import (
	"context"
	"github.com/cebilon123/nbp-go/internal/limiter"
	"github.com/cebilon123/nbp-go/internal/logg"
	"log"
	"net/http"
	"time"
)

// checker is used to check
// state of NBP currencies and
// logs it to given logger targets
type checker struct {
	logger        logg.Logger
	checksAmount  int    // checksAmount - how many checks are going to be made in given time span
	secondsAmount int    // secondsAmount - time span
	host          string // host is URL address that checker will check

	ctx    context.Context
	cancel context.CancelFunc
}

func NewChecker(logger logg.Logger, checksAmount int, secondsAmount int, host string) *checker {
	ctx, cancel := context.WithCancel(context.Background())
	return &checker{
		logger:        logger,
		checksAmount:  checksAmount,
		secondsAmount: secondsAmount,
		host:          host,
		ctx:           ctx,
		cancel:        cancel,
	}
}

// Start starts checker which will
// do checks based on passed config.
func (c *checker) Start() {

	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			c.cancel()
		}
	}()

	l := limiter.NewLimiter(time.Duration(c.secondsAmount)*time.Second, c.checksAmount)

	go func(ctx context.Context) {
		for {
			l.Wait()
			select {
			case <-ctx.Done():
				return
			default:
				start := time.Now()

				res, _ := http.Get(c.host)

				if res.StatusCode >= 400 {
					logData := &logg.LogData{
						Time:               start,
						Duration:           time.Since(start),
						Currency:           "ERROR",
						Mid:                0,
						ResponseStatusCode: res.StatusCode,
						IsJson:             false,
						IsJsonSyntaxValid:  false,
					}

					if err := c.logger.LogStatus(logData); err != nil {
						log.Println(err)
					}
				}

				logData := &logg.LogData{
					Time:               start,
					Duration:           time.Since(start),
					Currency:           "EUR",
					Mid:                0,
					ResponseStatusCode: res.StatusCode,
					IsJson:             false,
					IsJsonSyntaxValid:  false,
				}
				if err := c.logger.LogStatus(logData); err != nil {
					log.Println(err)
				}
			}
		}
	}(c.ctx)

	<-c.ctx.Done()
}

// Close gracefully closes checker
func (c *checker) Close() {
	c.cancel()
}

// Wait waits for checker to be closed
func (c *checker) Wait() {
	<-c.ctx.Done()
}
