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

import "time"

const (
	DateTimeFormat                   = "2006-01-02 15:04:05"
	CompactDateTimeFormat            = "20060102150405"
	CompactDateTimeFormatWithoutYear = "0102150405"
	DateFormat                       = "2006-01-02"
	TimeFormat                       = "15:04:05"
)

var (
	Unix1970 = time.Unix(0, 0).UTC()
)

func UnixMilli(t time.Time) int64 {
	return t.Unix()*1e3 + int64(t.Nanosecond()/1e6)
}

func UnixMicro(t time.Time) int64 {
	return t.Unix()*1e6 + int64(t.Nanosecond()/1e3)
}

func CompareDateUTC(t1, t2 time.Time) int {
	v1 := t1.Unix() / 86400
	v2 := t2.Unix() / 86400
	switch {
	case v1 == v2:
		return 0
	case v1 < v2:
		return -1
	default:
		return 1
	}
}

func TimestampPass(timestamp, t int64) bool {
	return t <= timestamp
}
