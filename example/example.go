package main

import (
	"runtime"
	"time"

	"github.com/leaanthony/spinner"
)

func main() {

	// Default Success
	a := spinner.New("This is a success")
	a.Start()
	time.Sleep(time.Second * 2)
	a.Success()

	// Default Error
	a = spinner.New("This is an error")
	a.Start()
	time.Sleep(time.Second * 2)
	a.Error()

	// Custom Success
	a = spinner.New("This is a custom success message")
	a.Start()
	time.Sleep(time.Second * 2)
	a.Success("Awesome!")

	// Custom Error
	a = spinner.New("This is a custom error message")
	a.Start()
	time.Sleep(time.Second * 2)
	a.Error("Much sad")

	// Formatted Success
	a = spinner.New("This is a formatted custom success message")
	a.Start()
	time.Sleep(time.Second * 2)
	spin := "Spinner"
	awesome := "Awesome"
	a.Successf("%s is %s!", spin, awesome)

	// Formatted Error
	a = spinner.New("This is a formatted custom error message")
	a.Start()
	secs := 2
	time.Sleep(time.Second * time.Duration(secs))
	a.Errorf("I waited %d seconds to error!", secs)

	// Reuse spinner!
	a.Restart("Spinner reuse FTW!")
	time.Sleep(time.Second * 2)
	a.Success()

	// Custom Spinner chars + symbols
	switch runtime.GOOS {
	case "windows":
		a.SetSpinFrames([]string{"^", ">", "v", "<"})
		a.SetSuccessSymbol("+")
	default:
		a.SetSpinFrames([]string{"🌕", "🌖", "🌗", "🌘", "🌑", "🌒", "🌓", "🌔"})
		a.SetSuccessSymbol("👍")
	}
	a.Restart("Custom spinner + Success Symbol!")
	time.Sleep(time.Second * 2)
	a.Success()

	// Custom Spinner chars + symbols
	switch runtime.GOOS {
	case "windows":
		a.SetSpinFrames([]string{".", "o", "O", "@", "*"})
		a.SetErrorSymbol("!")
	default:
		a.SetSpinFrames([]string{"🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚", "🕛"})
		a.SetErrorSymbol("💩")
	}
	a.Restart("Custom spinner + Error Symbol!")
	time.Sleep(time.Second * 2)
	a.Error()
}
