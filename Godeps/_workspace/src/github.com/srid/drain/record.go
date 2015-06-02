package drain

import (
	"fmt"
	"github.com/bmizerany/lpx"
)

// Record represents a log line processed by drains
type Record struct {
	Header *lpx.Header
	Data   []byte
}

// String returns the original log string back
func (r Record) String() string {
	line := fmt.Sprintf("%s %s %s %s %s %s %s\n",
		string(r.Header.PrivalVersion),
		string(r.Header.Time),
		string(r.Header.Hostname),
		string(r.Header.Name),
		string(r.Header.Procid),
		string(r.Header.Msgid),
		string(r.Data))
	return line
}
