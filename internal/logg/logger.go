package logg

import (
	"time"
)

//Logger must be implemented by all loggers
type Logger interface {
	// LogStatus logs given data as param
	LogStatus(data *LogData) error
}

// LogData represents data that should
// be passed to logger.
type LogData struct {
	Time               time.Time // Time is about how long given request has taken
	Currency           string    // Currency is a currency CODE being checked during request
	Mid                float32   // Mid is a value of currency in given request
	ResponseStatusCode int       // ResponseStatusCode is a status code of the request's response
	IsJson             bool      // IsJson is true whenever given response content-type is json
	IsJsonSyntaxValid  bool      // IsJsonSyntaxValid is true whenever there weren't any json parsing errors during request
}
