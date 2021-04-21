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

package wunsafe

import "strconv"

type String []byte

func (s String) Equals(str string) bool {
	if len(s) > 0 {
		return BytesToString(s) == str
	} else {
		return str == ""
	}
}

func (s String) Atoi64() (int64, error) {
	const fnAtoi64 = "Atoi64"

	sLen := len(s)
	if 0 < sLen && sLen < 19 {
		// Fast path for small integers that fit int type.
		s0 := s
		if s[0] == '-' || s[0] == '+' {
			s = s[1:]
			if len(s) < 1 {
				return 0, &strconv.NumError{Func: fnAtoi64, Num: string(s0), Err: strconv.ErrSyntax}
			}
		}

		n := int64(0)
		for _, ch := range []byte(s) {
			ch -= '0'
			if ch > 9 {
				return 0, &strconv.NumError{Func: fnAtoi64, Num: string(s0), Err: strconv.ErrSyntax}
			}
			n = n*10 + int64(ch)
		}
		if s0[0] == '-' {
			n = -n
		}
		return n, nil
	}

	// Slow path for invalid, big, or underscored integers.
	i64, err := strconv.ParseInt(string(s), 10, 0)
	if nerr, ok := err.(*strconv.NumError); ok {
		nerr.Func = fnAtoi64
	}
	return i64, err
}

func (s String) Atou64() (uint64, error) {
	const fnAtou64 = "Atou64"

	sLen := len(s)
	if 0 < sLen && sLen < 19 {
		// Fast path for small integers that fit int type.
		s0 := s
		if s[0] == '-' {
			return 0, &strconv.NumError{Func: fnAtou64, Num: string(s0), Err: strconv.ErrSyntax}
		}
		if s[0] == '+' {
			s = s[1:]
			if len(s) < 1 {
				return 0, &strconv.NumError{Func: fnAtou64, Num: string(s0), Err: strconv.ErrSyntax}
			}
		}

		n := uint64(0)
		for _, ch := range []byte(s) {
			ch -= '0'
			if ch > 9 {
				return 0, &strconv.NumError{Func: fnAtou64, Num: string(s0), Err: strconv.ErrSyntax}
			}
			n = n*10 + uint64(ch)
		}
		return n, nil
	}

	// Slow path for invalid, big, or underscored integers.
	i64, err := strconv.ParseUint(string(s), 10, 0)
	if nerr, ok := err.(*strconv.NumError); ok {
		nerr.Func = fnAtou64
	}
	return i64, err
}

func (s String) String() string {
	return string(s)
}

func (s String) UnsafeString() string {
	if len(s) > 0 {
		return BytesToString(s)
	} else {
		return ""
	}
}
