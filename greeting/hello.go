package greeting

import (
	"strings"
)

func Say(names []string) string {

	if len(names) < 1 {
		return "Make sure to input name argument"
	}

	return "Hello " + strings.Join(names, ", ") + "!"
}
