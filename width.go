package console

import (
	"golang.org/x/term"
)

func Width() int {
	width, _, err := term.GetSize(0)
	if err != nil {
		return 0
	}
	return width
}
