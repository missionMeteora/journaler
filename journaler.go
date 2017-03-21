package journaler

import (
	"io"
	"sync"

	"fmt"
	"github.com/fatih/color"
	"os"
	"runtime"
	"strings"
)

var (
	successColor = color.New(color.FgGreen, color.Bold)
	defaultColor = color.New(color.Bold)
	warningColor = color.New(color.FgYellow, color.Bold)
	errorColor   = color.New(color.FgRed, color.Bold)
)

const (
	msgFmt   = "%s %s\n"
	debugFmt = "%s (%s:%d) %s\n"
	labelFmt = "[%s]"
)

var j = NewJournal(os.Stdout)

// New returns a new Journaler
func New(prefixs ...string) *Journaler {
	return j.New(prefixs...)
}

// SetLabel will set a label for the global Journal
func SetLabel(key, value string) bool {
	return j.SetLabel(key, value)
}

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
func (j *Journal) Success(msg string) {
	j.mux.Lock()
	fmt.Fprintf(j.w, msgFmt, j.successStr, msg)
	j.mux.Unlock()
}

// Notification is for notification messages
func (j *Journal) Notification(msg string) {
	j.mux.Lock()
	fmt.Fprintf(j.w, msgFmt, j.notificationStr, msg)
	j.mux.Unlock()
}

// Warn is for warning messages
func (j *Journal) Warn(msg string) {
	j.mux.Lock()
	fmt.Fprintf(j.w, msgFmt, j.warningStr, msg)
	j.mux.Unlock()
}

// Error is for error messages
func (j *Journal) Error(msg string) {
	j.mux.Lock()
	fmt.Fprintf(j.w, msgFmt, j.errorStr, msg)
	j.mux.Unlock()
}

// Output is for custom messages
func (j *Journal) Output(label, color, msg string) {
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

	fmt.Fprintf(j.w, msgFmt, label, msg)
	j.mux.Unlock()
}

// Debug is for debug messages
func (j *Journal) Debug(msg string) {
	fn, ln := getDebugVals()
	j.mux.Lock()
	fmt.Fprintf(j.w, debugFmt, j.debugStr, fn, ln, msg)
	j.mux.Unlock()
}

// New returns a new Journaler
func (j *Journal) New(prefixs ...string) *Journaler {
	return &Journaler{
		j:      j,
		prefix: strings.Join(prefixs, " :: ") + " :: ",
	}
}

// Journaler is for logging with a prefix
type Journaler struct {
	j      *Journal
	prefix string
}

// Success is for success messages
func (j *Journaler) Success(msg string) {
	j.j.Success(j.prefix + msg)
}

// Notification is for notification messages
func (j *Journaler) Notification(msg string) {
	j.j.Notification(j.prefix + msg)
}

// Warn is for warning messages
func (j *Journaler) Warn(msg string) {
	j.j.Warn(j.prefix + msg)
}

// Error is for error messages
func (j *Journaler) Error(err error) {
	j.j.Error(j.prefix + err.Error())
}

// Output is for custom messages
func (j *Journaler) Output(label, color, msg string) {
	j.j.Output(label, color, j.prefix+msg)
}

// Debug is for debug messages
func (j *Journaler) Debug(msg string) {
	j.j.Debug(j.prefix + msg)
}

func getShort(file string) string {
	spl := strings.Split(file, "/")
	ns := make([]string, 0, len(spl))

	for i, part := range spl {
		if part == "" {
			continue
		}

		if i+1 < len(spl) {
			ns = append(ns, string(part[0]))
		} else {
			ns = append(ns, part)
		}
	}

	return "/" + strings.Join(ns, "/")
}

func getDebugVals() (filename string, lineNumber int) {
	_, filename, lineNumber, _ = runtime.Caller(3)
	filename = getShort(filename)
	return
}
