package spinner

import (
	"fmt"
	"runtime"
	"time"

	"github.com/fatih/color"
)

// Specialise the type
type status int

// Status code constants
const (
	errorStatus status = iota
	successStatus
)

// Gets the default spinner frames based on the operating system
func getDefaultSpinnerFrames() []string {
	switch runtime.GOOS {
	case "windows":
		return []string{"|", "/", "-", "\\"}
	default:
		return []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
	}
}

// Gets the default status symbols based on the operating system
func getStatusSymbols() (successSymbol, errorSymbol string) {
	switch runtime.GOOS {
	case "windows":
		return " ", " "
	default:
		return "✓", "✗"
	}
}

// Spinner defines our spinner data
// spinSpeed defaults to 100ms
type Spinner struct {
	message       string        // message to display
	stopChan      chan struct{} // exit channel
	exitStatus    status        // Status of exit
	successSymbol string        // Symbol printed when Success() called
	errorSymbol   string        // Symbol printed when Error() called
	spinFrames    []string      // Spinset frames
	spinSpeed     int           // Delay between spinner updates in milliseconds
	currentLine   string
}

// NewSpinner creates a new spinner and sets up the default values
func NewSpinner(message string) *Spinner {
	successSymbol, errorSymbol := getStatusSymbols()
	return &Spinner{
		message:       message,
		stopChan:      make(chan struct{}),
		successSymbol: successSymbol,
		errorSymbol:   errorSymbol,
		spinFrames:    getDefaultSpinnerFrames(),
		spinSpeed:     100,
	}
}

// New is solely here to make code cleaner for importers
// EG: spinner.New(...)
func New(message string) *Spinner {
	return NewSpinner(message)
}

// SetSuccessSymbol sets the symbol displayed on success
func (s *Spinner) SetSuccessSymbol(symbol string) {
	s.successSymbol = symbol
}

// SetErrorSymbol sets the symbol displayed on error
func (s *Spinner) SetErrorSymbol(symbol string) {
	s.errorSymbol = symbol
}

// SetSpinFrames makes the spinner use the given characters
func (s *Spinner) SetSpinFrames(frames []string) {
	s.spinFrames = frames
}

// SetSpinSpeed sets the speed of the spinner animation
// The lower the value, the faster the spin
func (s *Spinner) SetSpinSpeed(ms int) {
	s.spinSpeed = ms
}

// Start the spinner!
func (s *Spinner) Start() {

	// make it look tidier
	hideCursor()

	// spawn off a goroutine to handle the animation
	go func() {

		// Start at the first frame
		frameNumber := 0

		// Setup frame ticker
		ticker := time.NewTicker(time.Millisecond * time.Duration(s.spinSpeed)).C

		// Let's go!
		for {
			select {
			// For each frame tick
			case <-ticker:
				// Rewind to start of line and print the current frame and message.
				// Note: We don't fully clear the line here as this causes flickering
				// under windows
				fmt.Printf("\r")
				s.currentLine = fmt.Sprintf("%s %s", s.spinFrames[frameNumber], s.message)
				fmt.Printf(s.currentLine)

				// Move to next spinner frame and if we hit the end, loop to the start.
				frameNumber++
				frameNumber = frameNumber % len(s.spinFrames)

			// If we get a stop signal
			case <-s.stopChan:
				// Quit the animation
				return
			}
		}
	}()
}

// Restart sets the message and starts the spinner again.
// Useful for reusing a single spinner
func (s *Spinner) Restart(message string) {
	s.message = message
	s.Start()
}

// stop will stop the spinner.
// The final message will either be the current message
// or the optional, given message.
// Success status will print the message in green
// Error status will print the message in red
func (s *Spinner) stop(message ...string) {

	// Issue stop signal to animation
	s.stopChan <- struct{}{}

	// If we have an optional message, save it
	if len(message) > 0 {
		s.message = message[0]
	}

	// Clear the line, because a new message may be shorter than the original
	s.clearCurrentLine()

	// Output the symbol and message depending on the status code
	if s.exitStatus == errorStatus {
		color.HiRed("\r%s %s", s.errorSymbol, s.message)
	} else {
		color.HiGreen("\r%s %s", s.successSymbol, s.message)
	}

	// Show the cursor again
	showCursor()
}

// Error stops the spinner and sets the status code to error
// Optional message to print instead of current message
func (s *Spinner) Error(message ...string) {
	s.exitStatus = errorStatus
	s.stop(message...)
}

// Errorf stops the spinner, formats and sets the status code to error
// Formats and prints the given message instead of current message
func (s *Spinner) Errorf(format string, args ...interface{}) {
	s.exitStatus = errorStatus
	message := fmt.Sprintf(format, args...)
	s.stop(message)
}

// Success stops the spinner and sets the status code to success
// Optional message to print instead of current message
func (s *Spinner) Success(message ...string) {
	s.exitStatus = successStatus
	s.stop(message...)
}

// Successf stops the spinner, formats and sets the status code to success
// Formats and prints the given message instead of current message
func (s *Spinner) Successf(format string, args ...interface{}) {
	s.exitStatus = successStatus
	message := fmt.Sprintf(format, args...)
	s.stop(message)
}
