package console

import (
	"fmt"
	"time"
)

type Spinner struct {
	Stylesheet Stylesheet
	// How long it takes to move to a new sequence
	Frequency time.Duration
}

func (s *Spinner) Render(line []Cell) (n int, err error) {
	now := time.Now()

	time_since_start := now.Sub(hud.start_time)
	ticks_since_start := int(time_since_start / s.Frequency)

	sequence_count := 0
	for _, cell := range s.Stylesheet.Sequence {
		if cell.Rune == 0 {
			break
		}
		sequence_count++
	}
	if sequence_count == 0 {
		err = fmt.Errorf("console: you can't have a spinner with no sequences")
		return
	}

	sequence_index := ticks_since_start % sequence_count
	line[0] = s.Stylesheet.Sequence[sequence_index]

	n = 1
	return
}
