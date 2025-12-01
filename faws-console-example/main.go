package main

import (
	"crypto/rand"
	"math/big"
	"time"

	"github.com/faws-vcs/console"
)

func random_int(n int) (i int) {
	n_big, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		panic(err)
	}
	i = int(n_big.Int64())
	return
}

func main() {
	console.Open()

	bars_progress := make([]float64, 8)

	console.Println("Hello world")

	console.RenderFunc(func(hud *console.Hud) {
		var spinner console.Spinner
		spinner.Stylesheet.Sequence[0] = console.Cell{'⡿', console.BrightBlue, 0}
		spinner.Stylesheet.Sequence[1] = console.Cell{'⣟', console.BrightBlue, 0}
		spinner.Stylesheet.Sequence[2] = console.Cell{'⣯', console.BrightBlue, 0}
		spinner.Stylesheet.Sequence[3] = console.Cell{'⣷', console.BrightBlue, 0}
		spinner.Stylesheet.Sequence[4] = console.Cell{'⣾', console.BrightBlue, 0}
		spinner.Stylesheet.Sequence[5] = console.Cell{'⣽', console.BrightBlue, 0}
		spinner.Stylesheet.Sequence[6] = console.Cell{'⣻', console.BrightBlue, 0}
		spinner.Stylesheet.Sequence[7] = console.Cell{'⢿', console.BrightBlue, 0}
		spinner.Frequency = 200 * time.Millisecond

		var message console.Text
		message.Stylesheet.Alignment = console.Right
		message.Stylesheet.Width = 21
		message.Add("Stuff is happening...", console.Black, console.BrightBlue)

		hud.Line(&spinner, &message)

		var progress_bar console.ProgressBar
		progress_bar.Stylesheet.Alignment = console.Left
		progress_bar.Stylesheet.Width = console.Width()

		progress_bar.Stylesheet.Sequence[console.PbCaseLeft] = console.Cell{'[', 0, 0}
		progress_bar.Stylesheet.Sequence[console.PbCaseRight] = console.Cell{']', 0, 0}
		progress_bar.Stylesheet.Sequence[console.PbFluid] = console.Cell{'#', 0, 0}
		progress_bar.Stylesheet.Sequence[console.PbVoid] = console.Cell{'.', 0, 0}
		progress_bar.Stylesheet.Sequence[console.PbTail] = console.Cell{'#', 0, 0}
		progress_bar.Stylesheet.Sequence[console.PbHead] = console.Cell{'#', 0, 0}

		for i := range bars_progress {
			progress_bar.Progress = bars_progress[i]

			hud.Line(&progress_bar)
		}
	})

	for {
		for i := range bars_progress {
			bars_progress[i] += float64(random_int(200+(random_int(50)*10))) / 10000
		}

		all_progress_bars_complete := true
		for _, bar_progress := range bars_progress {
			if bar_progress < 1.0 {
				all_progress_bars_complete = false
				break
			}
		}
		if all_progress_bars_complete {
			break
		}
		console.SwapHud()
		time.Sleep(50 * time.Millisecond)
	}
	console.SwapHud()

	console.Close()
}
