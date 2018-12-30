// +build windows

package spinner

import (
	"fmt"

	"github.com/leaanthony/wincursor"
)

func showCursor() {
	err := wincursor.Show()
	if err != nil {
		panic(err)
	}
}

func hideCursor() {
	err := wincursor.Hide()
	if err != nil {
		panic(err)
	}
}

func (s *Spinner) clearCurrentLine() {
	// *shudder*
	fmt.Printf("\r")

	// Get the current line length
	var length = len(s.getMessage()) + len(s.getCurrentSpinnerFrame()) + 1

	for i := 0; i < length; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("\r")
}
