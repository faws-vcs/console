package console

var (
	ansi_cr = []byte{'\r'}

	// ANSI control sequences

	// ANSI escape       (0x1B)
	// CSI               ([)
	// Reset             (0)
	// SGR               (m)
	ansics_reset = []byte{0x1B, '[', '0', 'm'}

	// ANSI escape       (0x1B)
	// CSI               ([)
	// Clear entire line (2K)
	ansics_erase_line = []byte{0x1B, '[', '2', 'K'}

	// ANSI escape       (0x1B)
	// CSI               ([)
	// Move cursor 1 up  (1A)
	ansics_cursor_up = []byte{0x1B, '[', '1', 'A'}
)
