package journaler

import (
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

// Success is for success messages, accepts a body format and values
func (j *Journal) Success(fmt string, vals ...interface{}) {
	j.mux.Lock()
	writeMsg(j.w, msgFmt, j.successStr, getMsg(fmt, vals))
	j.mux.Unlock()
}

// Notification is for notification messages, accepts a body format and values
func (j *Journal) Notification(fmt string, vals ...interface{}) {
	j.mux.Lock()
	writeMsg(j.w, msgFmt, j.notificationStr, getMsg(fmt, vals))
	j.mux.Unlock()
}

// Warning is for warning messages, accepts a body format and values
func (j *Journal) Warning(fmt string, vals ...interface{}) {
	j.mux.Lock()
	writeMsg(j.w, msgFmt, j.warningStr, getMsg(fmt, vals))
	j.mux.Unlock()
}

// Error is for error messages, accepts a body format and values
func (j *Journal) Error(fmt string, vals ...interface{}) {
	j.mux.Lock()
	writeMsg(j.w, msgFmt, j.errorStr, getMsg(fmt, vals))
	j.mux.Unlock()
}

// Output is for custom messages
func (j *Journal) Output(label, color string, fmt string, vals ...interface{}) {
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

	writeMsg(j.w, msgFmt, label, getMsg(fmt, vals))
	j.mux.Unlock()
}

// Debug is for debug messages
func (j *Journal) Debug(fmt string, vals ...interface{}) {
	// We call the unexported debug func so we have the same number of frames to skip when calling runtime.Caller
	j.debug(fmt, vals)
}

// debug is for debug messages
func (j *Journal) debug(fmt string, vals ...interface{}) {
	fn, ln := getDebugVals()
	j.mux.Lock()
	writeDebug(j.w, j.debugStr, fn, ln, getMsg(fmt, vals))
	j.mux.Unlock()
}

// New returns a new Journaler
func (j *Journal) New(prefixs ...string) *Journaler {
	return &Journaler{
		j:      j,
		prefix: strings.Join(prefixs, " :: ") + " :: ",
	}
}
