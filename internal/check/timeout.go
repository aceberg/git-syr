package check

import (
	"fmt"
	"time"
)

func Timeout(timeout string) (string, error) {
	t, err := time.ParseDuration(timeout)
	seconds := fmt.Sprintf("%g", t.Seconds())

	// fmt.Println("TIMEOUT:", seconds, "ERR:", err)

	return seconds, err
}
