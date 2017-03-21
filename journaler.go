package journaler

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"runtime"
	"strings"
)

const (
	msgFmt   = "%s %v\n"
	debugFmt = "%s (%s:%d) %v\n"
	labelFmt = "[%s]"
)

var (
	successColor = color.New(color.FgGreen, color.Bold)
	defaultColor = color.New(color.Bold)
	warningColor = color.New(color.FgYellow, color.Bold)
	errorColor   = color.New(color.FgRed, color.Bold)
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

var (
	// Success is for success messages
	Success = j.Success
	// Notification is for notifications
	Notification = j.Notification
	// Warning is for warnings
	Warning = j.Warning
	// Error is for error messages
	Error = j.Error
	// Debug is for debugging messages
	Debug = j.debug
)

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

// Warning is for warning messages
func (j *Journaler) Warning(val interface{}) {
	j.j.Warning(getPrefixedValue(j.prefix, val))
}

// Error is for error messages
func (j *Journaler) Error(val interface{}) {
	j.j.Error(getPrefixedValue(j.prefix, val))
}

// Output is for custom messages
func (j *Journaler) Output(label, color string, val interface{}) {
	j.j.Output(label, color, getPrefixedValue(j.prefix, val))
}

// Debug is for debug messages
func (j *Journaler) Debug(val interface{}) {
	j.j.debug(getPrefixedValue(j.prefix, val))
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
