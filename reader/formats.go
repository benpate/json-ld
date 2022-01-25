package reader

import "time"

func AsTime(value string) time.Time {
	result, _ := time.Parse(time.RFC3339, value)
	return result
}
