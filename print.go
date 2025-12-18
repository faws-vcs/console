package console

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

var status int = 0

// Print generic information to os.Stdout
func Println(args ...any) (n int, err error) {
	hud.guard.Lock()
	hud.erase(os.Stdout, hud.num_lines)
	n, err = fmt.Fprintln(os.Stdout, args...)
	hud.present(os.Stdout)
	hud.newline_printed = true
	hud.guard.Unlock()
	return
}

func Fatal(args ...any) {
	hud.guard.Lock()

	// exit with this code in Close
	status = 1

	text := fmt.Sprint(args...)
	line := make([]Cell, utf8.RuneCountInString(text)+7)
	WriteText(line, "fatal:", White, Red)
	WriteText(line[6:], " ", None, None)
	WriteText(line[7:], text, None, None)

	hud.erase(os.Stdout, hud.num_lines)
	var message bytes.Buffer
	render_line(&message, line)
	message.WriteByte('\n')

	os.Stdout.Write(message.Bytes())

	hud.guard.Unlock()

	Close()
}

// str
// ---
func Header(str string) {
	Println(str)
	var sb strings.Builder
	count := utf8.RuneCountInString(str)
	for i := 0; i < count; i++ {
		sb.WriteByte('-')
	}
	Println(sb.String())
}

func Quote(args ...any) {
	var buffer strings.Builder
	fmt.Fprint(&buffer, args...)

	lines := strings.Split(buffer.String(), "\n")
	Println()
	for _, line := range lines {
		Println(" ", line)
	}
	Println()
}
