package console

import (
	"bytes"
	"os"
	"time"
)

// Required: start up console systems at end of process
func Open() {
	hud.renderer_func = func(h *Hud) {}
	hud.newline_printed = true
	hud.active.Store(true)
	hud.start_time = time.Now()
}

func close_internal() {
	hud.active.Store(false)
	hud.swap()
	// Fatal() was called before Close()
	if hud.final_message != nil {
		var message bytes.Buffer
		render_line(&message, hud.final_message)
		message.WriteByte('\n')
		hud.final_message = nil
		os.Stdout.Write(message.Bytes())
	}
	os.Exit(status)
}

// Required: shut down console systems at end of process
func Close() {
	hud.guard.Lock()
	close_internal()
	hud.guard.Unlock()
}
