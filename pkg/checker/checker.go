package checker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cebilon123/nbp-go/internal/limiter"
	"github.com/cebilon123/nbp-go/internal/logg"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const jsonContentType = "application/json"

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

func NewChecker(logger logg.Logger, checksAmount int, secondsAmount int, host string, ctx context.Context, cancel context.CancelFunc) *checker {
	return &checker{
		logger:        logger,
		checksAmount:  checksAmount,
		secondsAmount: secondsAmount,
		host:          host,
		ctx:           ctx,
		cancel:        cancel,
	}
}

type nbpResponse struct {
	Rates []struct {
		No            string  `json:"no"`
		EffectiveDate string  `json:"effectiveDate"`
		Mid           float32 `json:"mid"`
	} `json:"rates"`
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
						Message:            "ERROR",
						ResponseStatusCode: res.StatusCode,
						IsJson:             false,
						IsJsonSyntaxValid:  false,
					}

					if err := c.logger.LogStatus(logData); err != nil {
						log.Println(err)
					}
				}

				contentType := res.Header.Get("Content-type")

				var nbpRes nbpResponse
				resBody, err := ioutil.ReadAll(res.Body)
				isJsonSyntaxValid := false
				err1 := json.Unmarshal(resBody, &nbpRes)

				if err == nil && err1 == nil {
					isJsonSyntaxValid = true
				}

				dates := getDatesWithoutPricesInRange(nbpRes)

				logData := &logg.LogData{
					Time:               start,
					Duration:           time.Since(start),
					Message:            fmt.Sprintln(dates),
					ResponseStatusCode: res.StatusCode,
					IsJson:             strings.Contains(contentType, jsonContentType),
					IsJsonSyntaxValid:  isJsonSyntaxValid,
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

// getDatesWithoutPricesInRange checks if prices are between
// 4.5 and 4.7 PLN. If the prices aren't it will
// return array of dates when those weren't in
// range.
func getDatesWithoutPricesInRange(r nbpResponse) []string {
	var ret []string

	for _, rate := range r.Rates {
		if rate.Mid < 4.5 || rate.Mid > 4.7 {
			s := rate.EffectiveDate + " "
			ret = append(ret, s)
		}
	}

	return ret
}
