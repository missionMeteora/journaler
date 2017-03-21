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
	msgFmt   = "%s %v\n"
	debugFmt = "%s (%s:%d) %v\n"
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

// Warn is for warning messages
func (j *Journal) Warn(val interface{}) {
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

// Journaler is for logging with a prefix
type Journaler struct {
	j      *Journal
	prefix string
}

// Success is for success messages
func (j *Journaler) Success(val interface{}) {
	j.j.Success(getPrefixedValue(j.prefix, val))
}

// Notification is for notification messages
func (j *Journaler) Notification(val interface{}) {
	j.j.Notification(getPrefixedValue(j.prefix, val))
}

// Warn is for warning messages
func (j *Journaler) Warn(val interface{}) {
	j.j.Warn(getPrefixedValue(j.prefix, val))
}

// Error is for error messages
func (j *Journaler) Error(err error) {
	j.j.Error(j.prefix + err.Error())
}

// Output is for custom messages
func (j *Journaler) Output(label, color string, val interface{}) {
	j.j.Output(label, color, getPrefixedValue(j.prefix, val))
}

// Debug is for debug messages
func (j *Journaler) Debug(val interface{}) {
	j.j.Debug(getPrefixedValue(j.prefix, val))
}

func getPrefixedValue(prefix string, val interface{}) string {
	return fmt.Sprintf("%s%v", prefix, val)
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
