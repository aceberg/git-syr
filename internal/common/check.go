package common

import (
	"fmt"
)

// CheckIfError prints error, if it is not nil
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("ERROR: %s", err))
}
