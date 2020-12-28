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

package misc

import (
	"bytes"
	"encoding/base64"
	"strconv"
	"strings"
	"unicode"
)

func Split(str string, sep string) []string {
	var r []string
	a := strings.Split(str, sep)
	for _, s := range a {
		trimmed := strings.TrimSpace(s)
		if len(trimmed) > 0 {
			r = append(r, trimmed)
		}
	}
	return r
}

func UnescapeUnicode(str string) string {
	if !strings.Contains(str, `\u`) {
		return str
	}

	var out bytes.Buffer
	var i int
	for i <= len(str)-6 {
		if str[i] == '\\' && str[i+1] == 'u' {
			if u, err := strconv.ParseUint(str[i+2:i+6], 16, 64); err == nil {
				out.WriteRune(rune(u))
				i += 6
				continue
			}
		}
		out.WriteByte(str[i])
		i++
	}

	if i < len(str) {
		out.WriteString(str[i:])
	}
	return out.String()
}

func UCFirst(str string) string {
	var r string
	for i, v := range str {
		switch i {
		case 0:
			r = string(unicode.ToUpper(v))
		default:
			return r + str[i:]
		}
	}
	return r
}

func LCFirst(str string) string {
	var r string
	for i, v := range str {
		switch i {
		case 0:
			r = string(unicode.ToLower(v))
		default:
			return r + str[i:]
		}
	}
	return r
}

func IndexStrings(a []string, str string) int {
	for i, s := range a {
		if s == str {
			return i
		}
	}
	return -1
}

func DecodeCompileTimeString(str string) (string, error) {
	str = strings.ReplaceAll(str, "#", "=")
	bts, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(bytes.TrimRight(bts, "\r\n")), nil
}
