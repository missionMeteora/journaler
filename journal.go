package journaler

import (
	"fmt"
	"io"
	"strings"
	"sync"
)

// NewJournal will return a new Journal
// Note: This is only needed if you need to output somewhere other than os.Stdout
func NewJournal(w io.Writer) *Journal {
	var j Journal
	j.w = w
	j.SetLabel("success", "Success")
	j.SetLabel("notification", "Notification")
	j.SetLabel("warning", "Warning")
	j.SetLabel("error", "Error")
	j.SetLabel("debug", "Debug")
	return &j
}

// Journal is the back-bone of the Journalers. Manages the writer and thread-safety
type Journal struct {
	mux sync.Mutex
	// Output writer
	w io.Writer

	successStr      string
	notificationStr string
	warningStr      string
	errorStr        string
	debugStr        string
}

// SetLabel will set a label
func (j *Journal) SetLabel(key, value string) (ok bool) {
	ok = true
	j.mux.Lock()

	switch key {
	case "success":
		j.successStr = successColor.Sprintf(labelFmt, value)
	case "notification":
		j.notificationStr = defaultColor.Sprintf(labelFmt, value)
	case "warning":
		j.warningStr = warningColor.Sprintf(labelFmt, value)
	case "error":
		j.errorStr = errorColor.Sprintf(labelFmt, value)
	case "debug":
		j.debugStr = defaultColor.Sprintf(labelFmt, value)

	default:
		ok = false
	}

	j.mux.Unlock()
	return
}

// Success is for success messages
func (j *Journal) Success(val interface{}) {
	j.mux.Lock()
	fmt.Fprintf(j.w, msgFmt, j.successStr, val)
	j.mux.Unlock()
}

// Notification is for notification messages
func (j *Journal) Notification(val interface{}) {
	j.mux.Lock()
	fmt.Fprintf(j.w, msgFmt, j.notificationStr, val)
	j.mux.Unlock()
}

// Warning is for warning messages
func (j *Journal) Warning(val interface{}) {
	j.mux.Lock()
	fmt.Fprintf(j.w, msgFmt, j.warningStr, val)
	j.mux.Unlock()
}

// Error is for error messages
func (j *Journal) Error(val interface{}) {
	j.mux.Lock()
	fmt.Fprintf(j.w, msgFmt, j.errorStr, val)
	j.mux.Unlock()
}

// Output is for custom messages
func (j *Journal) Output(label, color string, val interface{}) {
	j.mux.Lock()

	switch color {
	case "green":
		label = successColor.Sprintf(labelFmt, label)
	case "yellow":
		label = warningColor.Sprintf(labelFmt, label)
	case "red":
		label = errorColor.Sprintf(labelFmt, label)

	default:
		label = defaultColor.Sprintf(labelFmt, label)
	}

	fmt.Fprintf(j.w, msgFmt, label, val)
	j.mux.Unlock()
}

// Debug is for debug messages
func (j *Journal) Debug(val interface{}) {
	// We call the unexported debug func so we have the same number of frames to skip when calling runtime.Caller
	j.debug(val)
}

// debug is for debug messages
func (j *Journal) debug(val interface{}) {
	fn, ln := getDebugVals()
	j.mux.Lock()
	fmt.Fprintf(j.w, debugFmt, j.debugStr, fn, ln, val)
	j.mux.Unlock()
}

// New returns a new Journaler
func (j *Journal) New(prefixs ...string) *Journaler {
	return &Journaler{
		j:      j,
		prefix: strings.Join(prefixs, " :: ") + " :: ",
	}
}
