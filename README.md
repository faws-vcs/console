# faws-vcs/console

[![Go Reference](https://pkg.go.dev/badge/github.com/faws-vcs/faws.svg)](https://pkg.go.dev/github.com/faws-vcs/console)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

`console` provides an immediate-mode rendering framework for miniature terminal-based UIs.

In this package, you are provided with a live text area (called a Hud) that you can print to using your custom `RendererFunc`.

```go
// Open the console
console.Open()

// Print a normal line to the background (always printed above the Hud)
console.Println("Hello, this is normal line")

// Set a Hud renderer function
console.RenderFunc(func(hud *console.Hud) {
    var spin console.Spinner
    // foreground color = BrightBlue, background color = None (0)
    spin.Stylesheet.Sequence[0] = console.Cell{'⡿', console.BrightBlue, 0}
    spin.Stylesheet.Sequence[1] = console.Cell{'⣟', console.BrightBlue, 0}
    spin.Stylesheet.Sequence[2] = console.Cell{'⣯', console.BrightBlue, 0}
    spin.Stylesheet.Sequence[3] = console.Cell{'⣷', console.BrightBlue, 0}
    spin.Stylesheet.Sequence[4] = console.Cell{'⣾', console.BrightBlue, 0}
    spin.Stylesheet.Sequence[5] = console.Cell{'⣽', console.BrightBlue, 0}
    spin.Stylesheet.Sequence[6] = console.Cell{'⣻', console.BrightBlue, 0}
    spin.Stylesheet.Sequence[7] = console.Cell{'⢿', console.BrightBlue, 0}
    spin.Frequency = time.Second/3
    hud.Line(&spin)

    // Start displaying hud elements
    var progress_bar console.ProgressBar

    // | [====>   ]                |
    progress_bar.Stylesheet.Alignment = console.Left
    progress_bar.Stylesheet.Width = 10 // or console.Width() to fill the entire line

    progress_bar.Stylesheet.Sequence[console.PbCaseLeft] = console.Cell{'[', 0, 0}
    progress_bar.Stylesheet.Sequence[console.PbCaseRight] = console.Cell{']', 0, 0}
    progress_bar.Stylesheet.Sequence[console.PbFluid] = console.Cell{'=', 0, 0}
    progress_bar.Stylesheet.Sequence[console.PbVoid] = console.Cell{' ', 0, 0}
    progress_bar.Stylesheet.Sequence[console.PbTail] = console.Cell{'<', 0, 0}
    progress_bar.Stylesheet.Sequence[console.PbHead] = console.Cell{'>', 0, 0}

    progress_bar.Progress = 0.5

    var progress_text console.Text
    // | [====>   ] 1/2
    progress_text.Stylesheet.Alignment = console.Right
    // necessary to put space between text elements
    progress_text.Stylesheet.Margin[console.Left] = 1
    progress_text.Add(fmt.Sprintf("%d/%d", 1, 2), console.BrightGreen, console.None) 
    // render lines to the buffer
    hud.Line(&progress_bar, &progress_text)
})

// run your application here

// ...

// Manually trigger a re-render of the Hud from your app
console.SwapHud()

// Forces SwapHud to be called every second
console.SwapInterval(1 * time.Second)

console.Close()
```