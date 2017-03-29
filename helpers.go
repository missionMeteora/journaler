package journaler

import (
	"fmt"
	"io"
	"runtime"
	"strings"
)

func writeMsg(w io.Writer, msgFmt, status, msg string, args ...interface{}) {
	vs := make([]interface{}, 0, len(args)+2)
	vs = append(vs, status)
	vs = append(vs, args...)
	vs = append(vs, msg)
	fmt.Fprintf(j.w, msgFmt, vs...)
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

func getMsg(mfmt string, vals []interface{}) string {
	if len(mfmt) == 0 {
		mfmt = "%v"
	}

	switch len(vals) {
	case 0:
		return mfmt

	case 1:
		return fmt.Sprintf(mfmt, vals[0])

	default:
		return fmt.Sprintf(mfmt, vals...)
	}
}

func getPrefixed(pfmt, fmt string) string {
	if len(fmt) == 0 {
		fmt = "%v"
	}

	return pfmt + fmt
}
