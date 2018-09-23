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
	s.locks[currentLineLock].Lock()
	var length = len(s.currentLine)
	s.locks[currentLineLock].Unlock()

	for i := 0; i < length; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("\r")
}
