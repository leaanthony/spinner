// +build !windows

package spinner

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

func init() {

}

func showCursor() {
	fmt.Printf("\033[?25h")
}

func hideCursor() {
	fmt.Printf("\033[?25l")
}

func (s *Spinner) clearCurrentLine() {
	fmt.Printf("\r\033[0K")
}

type winsize struct {
	rows    uint16
	cols    uint16
	xpixels uint16
	ypixels uint16
}

func (s *Spinner) updateTermSize() error {
	out, err := os.OpenFile("/dev/tty", syscall.O_WRONLY, 0)
	if err != nil {
		return err
	}
	var sz winsize
	_, _, _ = syscall.Syscall(syscall.SYS_IOCTL,
		out.Fd(), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&sz)))
	out.Close()
	s.termWidth.SetValue(int(sz.cols))
	s.termHeight.SetValue(int(sz.rows))
	return nil
}
