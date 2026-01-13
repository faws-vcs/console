package console

import (
	"os"
	"time"
)

// Required: start up console systems at end of process
func Open() {
	hud.newline_printed = true
	hud.active.Store(true)
	hud.start_time = time.Now()
}

// Required: shut down console systems at end of process
func Close() {
	hud.active.Store(false)
	hud.guard.Lock()
	hud.swap()
	hud.guard.Unlock()
	os.Exit(status)
}
