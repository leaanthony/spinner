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
	for i := 0; i < len(s.currentLine); i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("\r")
}
