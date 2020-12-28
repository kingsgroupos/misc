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

package wtime

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/kingsgroupos/misc/wunsafe"
)

const Day = 24 * time.Hour

func ParseDuration(s string) (time.Duration, error) {
	idx := strings.IndexRune(s, 'd')
	if idx < 0 {
		return time.ParseDuration(s)
	}

	days, err := strconv.Atoi(s[:idx])
	if err != nil {
		return 0, errors.New("time: invalid duration " + s)
	}
	d1 := Day * time.Duration(days)
	if idx+1 >= len(s) {
		return d1, nil
	}
	d2, err := time.ParseDuration(s[idx+1:])
	if err != nil {
		return 0, err
	}

	return d1 + d2, nil
}

type Duration time.Duration

func (d Duration) D() time.Duration {
	return time.Duration(d)
}

func (d Duration) MarshalJSON() ([]byte, error) {
	s := d.String()
	n := len(s) + 2
	b := make([]byte, n)
	copy(b[1:], wunsafe.StringToBytes(s))
	b[0], b[n-1] = '"', '"'
	return b, nil
}

func (this *Duration) UnmarshalJSON(data []byte) error {
	n := len(data)
	if n < 2 {
		return nil
	}
	d, err := ParseDuration(string(data[1 : n-1]))
	if err != nil {
		return err
	}
	*this = Duration(d)
	return nil
}

func (d Duration) Hours() float64 {
	return time.Duration(d).Hours()
}

func (d Duration) Microseconds() int64 {
	return time.Duration(d).Microseconds()
}

func (d Duration) Milliseconds() int64 {
	return time.Duration(d).Milliseconds()
}

func (d Duration) Minutes() float64 {
	return time.Duration(d).Minutes()
}

func (d Duration) Nanoseconds() int64 {
	return time.Duration(d).Nanoseconds()
}

func (d Duration) Round(m time.Duration) time.Duration {
	return time.Duration(d).Round(m)
}

func (d Duration) Seconds() float64 {
	return time.Duration(d).Seconds()
}

func (d Duration) String() string {
	return time.Duration(d).String()
}

func (d Duration) Truncate(m time.Duration) time.Duration {
	return time.Duration(d).Truncate(m)
}
