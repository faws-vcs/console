package console

import (
	"bytes"
	"fmt"
	"os"
	"unicode/utf8"
)

var status int = 0

// Print generic information to os.Stdout
func Println(args ...any) (n int, err error) {
	hud.guard.Lock()
	hud.erase(hud.num_lines)
	n, err = fmt.Fprintln(os.Stdout, args...)
	hud.present()
	hud.newline_printed = true
	hud.guard.Unlock()
	return
}

func Info(args ...any) {
	Println(args...)
}

func Fatal(args ...any) {
	hud.guard.Lock()

	// exit with this code in Fatal
	status = 1

	text := fmt.Sprint(args...)
	line := make([]Cell, utf8.RuneCountInString(text)+7)
	WriteText(line, "fatal:", White, Red)
	WriteText(line[6:], " ", None, None)
	WriteText(line[7:], text, None, None)

	hud.erase(hud.num_lines)
	var message bytes.Buffer
	render_line(&message, line)
	message.WriteByte('\n')

	os.Stdout.Write(message.Bytes())

	hud.guard.Unlock()

	Close()
}
