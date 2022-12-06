package check

import (
	"fmt"
	"time"
)

// TimeToSec checks if time is correct and converts it to seconds
func TimeToSec(timeout string) (string, error) {
	t, err := time.ParseDuration(timeout)
	seconds := fmt.Sprintf("%g", t.Seconds())

	// fmt.Println("TIMEOUT:", seconds, "ERR:", err)

	return seconds, err
}
