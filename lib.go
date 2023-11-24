// Package testmod do sth
package testmod

import (
	"fmt"
)

var line = 1

func Hi(name string) string {
	line++
	return fmt.Sprintf("%d: Hi, %s!", line, name)
}
