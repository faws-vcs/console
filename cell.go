package console

import (
	"bytes"
	"fmt"
	"io"

	"github.com/mattn/go-runewidth"
)

// A cell is a colorable UTF-8 character that takes up 1 width in the terminal (Hud).
type Cell struct {
	Rune rune
	// Foreground color
	Fg Color
	// Background color
	Bg Color
}

// some runes take up multiple cells
func rune_cell_width(r rune) int {
	var condition runewidth.Condition
	condition.EastAsianWidth = true
	condition.StrictEmojiNeutral = true
	return condition.RuneWidth(r)
}

func set_color(w io.Writer, fg, bg Color) {
	fg_sgr, ok := color_codes[color_key{false, fg}]
	if !ok {
		panic("invalid foreground color")
	}
	bg_sgr, ok := color_codes[color_key{true, bg}]
	if !ok {
		panic("invalid background color")
	}

	if fg == None && bg == None {
		w.Write(ansics_reset)
		return
	}
	if bg != None && fg == None {
		// just set background
		fmt.Fprintf(w, "\x1B[0;%dm", bg_sgr)
		return
	}
	if bg == None && fg != None {
		// just set foreground
		fmt.Fprintf(w, "\x1B[0;%dm", fg_sgr)
		return
	}
	// Set both
	fmt.Fprintf(w, "\x1B[%d;%dm", fg_sgr, bg_sgr)
}

// Convert a line into bytes with ANSI escape sequences
func render_line(buffer *bytes.Buffer, line []Cell) {
	var (
		current_fg_color Color
		current_bg_color Color
	)

	set_color(buffer, None, None)
	for _, cell := range line {
		if cell.Fg != current_fg_color || cell.Bg != current_bg_color {
			set_color(buffer, cell.Fg, cell.Bg)
		}
		current_bg_color = cell.Bg
		current_fg_color = cell.Fg
		if cell.Rune == 0 {
			buffer.WriteByte(' ')
			continue
		}
		if cell.Rune == 0x80 {
			// this is a padding cell. these are injected after a wide character is used, so we don't need to do anything here.
			continue
		}
		if cell.Rune == '\n' || cell.Rune == '\r' {
			break
		}
		buffer.WriteRune(cell.Rune)
	}
}
