package verbose

import "fmt"

const colorCyan = "\033[36m"

type Verbose struct {
	verboseLevel int
}

func NewVerbose(verboseLevel int) Verbose {
	return Verbose{verboseLevel}
}

func (v Verbose) Log1f(format string, a ...interface{}) {
	v.Logf(1, format, a...)
}

func (v Verbose) Log2f(format string, a ...interface{}) {
	v.Logf(2, format, a...)
}

func (v Verbose) Log3f(format string, a ...interface{}) {
	v.Logf(3, format, a...)
}

func (v Verbose) Logf(level int, format string, a ...interface{}) {
	if v.verboseLevel >= level {
		msg := fmt.Sprintf(format, a...)
		fmt.Printf("%sVerbose %d: %s%s\n", colorCyan, level, msg, "\033[0m")
	}
}
