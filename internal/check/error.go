package check

import (
	"fmt"
)

// IfError prints error, if it is not nil
func IfError(err error) bool {
	if err == nil {
		return false
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("ERROR: %s", err))
	return true
}
