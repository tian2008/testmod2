package testmod2

import (
	"fmt"
)

func Hi(str string, ver string) string {
	return fmt.Sprintf("Hi, %v! version is %v", str, ver)
}
