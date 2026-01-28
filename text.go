package console

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"unicode/utf8"
)

type Text struct {
	Stylesheet Stylesheet
	Text       []Cell
}

func (t *Text) Add(s string, fg, bg Color) (err error) {
	text_cells := make([]Cell, utf8.RuneCountInString(s))
	_, err = WriteText(text_cells, s, fg, bg)
	if err != nil {
		return
	}
	t.Text = append(t.Text, text_cells...)

	return
}

func (t *Text) Render(line []Cell) (n int, err error) {
	margin_left := t.Stylesheet.Margin[Left]
	margin_right := t.Stylesheet.Margin[Right]
	if t.Stylesheet.Width < margin_left+margin_right {
		err = fmt.Errorf("console: not enough width to render text margins")
		return
	}
	text_cells := make([]Cell, t.Stylesheet.Width)
	copy(text_cells[margin_left:len(text_cells)-margin_right], t.Text)
	// set margins
	for i := 0; i < margin_left; i++ {
		text_cells[i] = t.Stylesheet.Sequence[Left]
	}
	right_margin_index := len(text_cells) - margin_right
	for i := 0; i < t.Stylesheet.Margin[Right]; i++ {
		text_cells[right_margin_index+i] = t.Stylesheet.Sequence[Right]
	}
	// copy based on alignment
	var start_index int

	switch t.Stylesheet.Alignment {
	case Left:
		start_index = 0
	case Right:
		start_index = len(line) - t.Stylesheet.Width
	case Center:
		start_index = (len(line) / 2) - (t.Stylesheet.Width / 2)
	}

	n = copy(line[start_index:], text_cells)

	return
}

func WriteText(text []Cell, s string, fg, bg Color) (n int, err error) {
	buffer := bytes.NewBufferString(s)

	var r rune
	var padding int
	for i := 0; i < len(text); i++ {
		// we may need to insert padding cells
		// if a previous cell was wider than 1
		if padding > 0 {
			text[i].Rune = 0x80
			n = i + 1
			padding--
			continue
		}
		r, _, err = buffer.ReadRune()
		if errors.Is(err, io.EOF) {
			err = nil
			break
		}
		text[i].Rune = r
		text[i].Fg = fg
		text[i].Bg = bg
		n = i + 1

		cell_width := rune_cell_width(r)
		if cell_width > 1 {
			padding = cell_width - 1
		}
	}
	return
}
