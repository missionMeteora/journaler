package journaler

import (
	"io"
	"sync"

	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

var (
	successStr = color.New(color.FgGreen, color.Bold).Sprint("[Sucess]")
	warningStr = color.New(color.FgYellow, color.Bold).Sprint("[Warning]")
	errorStr   = color.New(color.FgRed, color.Bold).Sprint("[Error]")
)

const msgFmt = "%s %s\n"

var j = Journal{
	w: os.Stdout,
}

// New returns a new Journaler
func New(prefixs ...string) *Journaler {
	return j.New(prefixs...)
}

type Journal struct {
	mux sync.Mutex
	// Output writer
	w io.Writer
}

// Success is for success messages
func (j *Journal) Success(msg string) {
	j.mux.Lock()
	j.w.Write([]byte(fmt.Sprintf(msgFmt, successStr, msg)))
	j.mux.Unlock()
}

// Warn is for warning messages
func (j *Journal) Warn(msg string) {
	j.mux.Lock()
	j.w.Write([]byte(fmt.Sprintf(msgFmt, warningStr, msg)))
	j.mux.Unlock()
}

// Error is for error messages
func (j *Journal) Error(msg string) {
	j.mux.Lock()
	j.w.Write([]byte(fmt.Sprintf(msgFmt, errorStr, msg)))
	j.mux.Unlock()
}

// New returns a new Journaler
func (j *Journal) New(prefixs ...string) *Journaler {
	return &Journaler{
		j:      j,
		prefix: strings.Join(prefixs, " :: ") + " :: ",
	}
}

type Journaler struct {
	j      *Journal
	prefix string
}

// Success is for success messages
func (j *Journaler) Success(msg string) {
	j.j.Success(j.prefix + msg)
}

// Warn is for warning messages
func (j *Journaler) Warn(msg string) {
	j.j.Warn(j.prefix + msg)
}

// Error is for error messages
func (j *Journaler) Error(err error) {
	j.j.Error(j.prefix + err.Error())
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

	return strings.Join(ns, "/")
}
