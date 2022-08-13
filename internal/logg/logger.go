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
	Time               time.Time     // Time: time when given request was done
	Duration           time.Duration // Duration is about how long given request has taken
	Message            string        // Message is a currency CODE being checked during request or ERROR if request failed
	ResponseStatusCode int           // ResponseStatusCode is a status code of the request's response
	IsJson             bool          // IsJson is true whenever given response content-type is json
	IsJsonSyntaxValid  bool          // IsJsonSyntaxValid is true whenever there weren't any json parsing errors during request
}
