// Package testmod do sth
package testmod

import (
	"errors"
	"fmt"
)

var line = 0

func Hi(name, lang string) (string, error) {
	line--
	switch lang {
	case "en":
		return fmt.Sprintf("%d: Hi, %s!", line, name), nil
	case "pt":
		return fmt.Sprintf("%d: Oi, %s!", line, name), nil
	case "es":
		return fmt.Sprintf("%d: ¡Hola, %s!", line, name), nil
	case "fr":
		return fmt.Sprintf("%d: Bonjour, %s!", line, name), nil
	default:
		return "", errors.New("unknown language")
	}
}
