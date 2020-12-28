// BSD 3-Clause License
//
// Copyright (c) 2020, Kingsgroup
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

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
