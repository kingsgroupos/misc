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

package wlang

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	rexPos    = regexp.MustCompile(`{\d+[^}]*}`)
	rexNumFmt = regexp.MustCompile(`^:[0#]*\.([0#]+)`)
)

func Sprint(lang string, a ...interface{}) (string, error) {
	l, err := language.Parse(lang)
	return message.NewPrinter(l).Sprint(a...), err
}

func Sprintln(lang string, a ...interface{}) (string, error) {
	l, err := language.Parse(lang)
	return message.NewPrinter(l).Sprintln(a...), err
}

func convertFormat(format string) string {
	if strings.Contains(format, "{{") {
		format = strings.Replace(format, "{{", "{", -1)
	}
	if strings.Contains(format, "}}") {
		format = strings.Replace(format, "}}", "}", -1)
	}
	if strings.Contains(format, "%") {
		format = strings.Replace(format, "%", "%%", -1)
	}
	return rexPos.ReplaceAllStringFunc(format, func(s string) string {
		i := 2
		for n := len(s); i < n; i++ {
			if ch := s[i]; ch < '0' || ch > '9' {
				break
			}
		}
		v, _ := strconv.Atoi(s[1:i])
		if s[i] == ':' {
			a := rexNumFmt.FindStringSubmatch(s[i:])
			if len(a) == 2 {
				return fmt.Sprintf("%%.%d[%d]f", len(a[1]), v+1)
			}
		}
		return fmt.Sprintf("%%[%d]v", v+1)
	})
}

func Sprintf(lang string, format string, a ...interface{}) (string, error) {
	l, err := language.Parse(lang)
	f := convertFormat(format)
	return message.NewPrinter(l).Sprintf(f, a...), err
}
