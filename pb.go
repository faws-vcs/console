package console

import "fmt"

const (
	// Stylesheet.Sequence indices for styling progress bars
	PbCaseLeft = iota
	PbCaseRight
	PbFluid
	PbVoid
	PbTail
	PbHead
)

type ProgressBar struct {
	Stylesheet Stylesheet
	// Must be in range from 0-1
	Progress float64
}

func (pb *ProgressBar) Render(line []Cell) (n int, err error) {
	actual_width := pb.Stylesheet.Width
	actual_width -= pb.Stylesheet.Margin[Left]
	actual_width -= pb.Stylesheet.Margin[Right]

	var start_index int

	switch pb.Stylesheet.Alignment {
	case Left:
		start_index = pb.Stylesheet.Margin[Left]
	case Right:
		start_index = len(line) - (pb.Stylesheet.Margin[Right] + pb.Stylesheet.Margin[Left] + actual_width)
	case Center:
		start_index = ((len(line) - start_index) / 2) - (actual_width / 2)
	}

	// calculate the number of progress cells to fill in
	progress_cells := int(float64(actual_width-2) * pb.Progress)

	if len(line) > (start_index + actual_width) {
		err = fmt.Errorf("console: ProgressBar.Render: not enough space")
		return
	}

	// [ ]
	line[start_index] = pb.Stylesheet.Sequence[PbCaseLeft]
	line[start_index+actual_width-1] = pb.Stylesheet.Sequence[PbCaseRight]

	for i := 1; i < (actual_width - 1); i++ {
		if progress_cells > i {
			line[start_index+i] = pb.Stylesheet.Sequence[PbFluid]
		} else if progress_cells == i {
			line[start_index+i] = pb.Stylesheet.Sequence[PbHead]
		} else if progress_cells < i {
			line[start_index+i] = pb.Stylesheet.Sequence[PbVoid]
		}
	}

	n = start_index + actual_width + pb.Stylesheet.Margin[Right]

	return
}
