package testmod2

import (
	"fmt"
)

func Hi(str string, ver int) string {
	return fmt.Sprintf("Hi, %v! current version is %v !", str, ver)
}
