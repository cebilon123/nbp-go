package log_format

import (
	"fmt"
	"github.com/cebilon123/nbp-go/internal/logg"
)

// GetString returns format string for log data
func GetString(data *logg.LogData) string {
	return fmt.Sprintf("Duration: %v | Currency: %s | Mid: %f | StatusCode: %d | IsJson: %t | IsJsonSyntaxValid: %t", data.Time, data.Currency, data.Mid, data.ResponseStatusCode, data.IsJson, data.IsJsonSyntaxValid)
}
