package console

import (
	"bytes"
	"io"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

var hud Hud

type Hud struct {
	active          atomic.Bool
	in_render       atomic.Bool
	guard           sync.Mutex
	renderer_func   RendererFunc
	num_lines       int
	lines           bytes.Buffer
	newline_printed bool
	start_time      time.Time
}

type Component interface {
	Render(line []Cell) (n int, err error)
}

type RendererFunc func(h *Hud)

// Set a function to render your Hud
func RenderFunc(f RendererFunc) {
	hud.guard.Lock()
	hud.renderer_func = f
	hud.guard.Unlock()
}

// non-goroutine-safe
// re-render the hud
func (h *Hud) render() {
	h.in_render.Store(true)
	h.num_lines = 0
	h.lines.Reset()
	h.renderer_func(h)
	h.in_render.Store(false)
}

// non-goroutine-safe
func (h *Hud) erase(w io.Writer, lines int) {
	if lines == 0 {
		return
	}
	var err error
	// erase the line
	if _, err = w.Write(ansics_erase_line); err != nil {
		panic(err)
	}
	for i := 0; i < lines-1; i++ {
		// move the cursor up 1
		if _, err = w.Write(ansics_cursor_up); err != nil {
			panic(err)
		}
		// erase the line
		if _, err = w.Write(ansics_erase_line); err != nil {
			panic(err)
		}
	}

	// move cursor to 0 cell
	if _, err = w.Write(ansi_cr); err != nil {
		panic(err)
	}
}

// render a line
func (h *Hud) Line(components ...Component) {
	if !h.in_render.Load() {
		panic("not in render")
	}

	// lines need to have a newline printed before them, not after.
	// we don't want to have an ugly space below the bottom of the Hud
	if !h.newline_printed {
		h.lines.WriteByte('\n')
		h.newline_printed = true
	}

	line_width := Width()
	if line_width == 0 {
		return
	}

	line := make([]Cell, line_width)
	index := 0
	for _, component := range components {
		n, err := component.Render(line[index:])
		if err != nil {
			WriteText(line, err.Error(), 0, 0)
			break
		}
		index += n
		if index >= line_width {
			break
		}
	}

	render_line(&h.lines, line)

	h.num_lines++
}

func (h *Hud) present(w io.Writer) {
	w.Write(h.lines.Bytes())
}

// Erases the previously printed hud and renders the updated version
func SwapHud() {
	hud.guard.Lock()
	// save the old number of lines
	num_lines := hud.num_lines
	// render the new Hud
	hud.render()
	// erase and present in one buffer to avoid flickering
	var swap_buffer bytes.Buffer
	// erase the old Hud
	hud.erase(&swap_buffer, num_lines)
	// print the new Hud
	hud.present(&swap_buffer)

	os.Stdout.Write(swap_buffer.Bytes())

	hud.guard.Unlock()
}

// Automatically swap at the designated interval
func SwapInterval(t time.Duration) {
	go func() {
		for {
			if hud.active.Load() {
				SwapHud()
			} else {
				return
			}
			time.Sleep(t)
		}
	}()
}
