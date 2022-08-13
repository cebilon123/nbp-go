package log_format

import (
	"fmt"
	"github.com/cebilon123/nbp-go/internal/logg"
	"time"
)

// GetString returns format string for log data
func GetString(data *logg.LogData) string {
	return fmt.Sprintf("[ %s ] Duration: %v ms | Message: %s | StatusCode: %d | IsJson: %t | IsJsonSyntaxValid: %t", data.Time.Format(time.StampMilli), data.Duration.Milliseconds(), data.Message, data.ResponseStatusCode, data.IsJson, data.IsJsonSyntaxValid)
}
