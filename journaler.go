package journaler

import (
	"os"

	"github.com/fatih/color"
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
func (j *Journaler) Success(fmt string, vals ...interface{}) {
	j.j.Success(j.prefix+fmt, vals...)
}

// Notification is for notification messages
func (j *Journaler) Notification(fmt string, vals ...interface{}) {
	j.j.Notification(j.prefix+fmt, vals...)
}

// Warning is for warning messages
func (j *Journaler) Warning(fmt string, vals ...interface{}) {
	j.j.Warning(j.prefix+fmt, vals...)
}

// Error is for error messages
func (j *Journaler) Error(fmt string, vals ...interface{}) {
	j.j.Error(j.prefix+fmt, vals...)
}

// Output is for custom messages
func (j *Journaler) Output(label, color string, fmt string, vals ...interface{}) {
	j.j.Output(label, color, j.prefix+fmt, vals...)
}

// Debug is for debug messages
func (j *Journaler) Debug(fmt string, vals ...interface{}) {
	j.j.debug(j.prefix+fmt, vals...)
}
