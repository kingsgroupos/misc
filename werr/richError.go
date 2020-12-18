package werr

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/kingsgroupos/misc"
)

type RichInfoStyle int

const (
	RichInfoStyle_JSON RichInfoStyle = iota
	RichInfoStyle_JSONFriendly
)

var Style RichInfoStyle

type RichError struct {
	err     error
	Details []interface{}
}

func NewRichError(err error, details ...interface{}) *RichError {
	if len(details) == 0 {
		if richErr, ok := err.(*RichError); ok {
			return richErr
		}
	}
	return &RichError{
		err:     err,
		Details: details,
	}
}

func (this *RichError) Error() string {
	var buf bytes.Buffer
	this.stringImpl(&buf)
	return buf.String()
}

func (this *RichError) stringImpl(w io.Writer) {
	_, _ = fmt.Fprint(w, this.err)
	this.printDetails(w)
}

func (this *RichError) printDetails(w io.Writer) {
	n := len(this.Details)
	if n == 0 {
		return
	}

	if Style == RichInfoStyle_JSONFriendly {
		for i := 0; i < n; i += 2 {
			k, v := this.Details[i], this.Details[i+1]
			const format = ", %s: %s"
			_, _ = fmt.Fprintf(w, format, fmt.Sprint(k), misc.ToPlainJSON(v))
		}
	} else {
		for i := 0; i < n; i += 2 {
			k, v := this.Details[i], this.Details[i+1]
			format := ", %q: %s"
			if i == 0 {
				format = " {%q: %s"
			}
			_, _ = fmt.Fprintf(w, format, fmt.Sprint(k), misc.ToPlainJSON(v))
		}
		_, _ = fmt.Fprint(w, "}")
	}
}

func (this *RichError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = fmt.Fprintf(s, "%+v", this.err)
			this.printDetails(s)
			return
		}
		fallthrough
	case 's', 'q':
		this.stringImpl(s)
	}
}

func (this *RichError) Unwrap() error {
	return this.err
}

func AsRichError(err error) (*RichError, bool) {
	var x *RichError
	ok := errors.As(err, &x)
	return x, ok
}
