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

import "testing"

func TestConvertFormat(t *testing.T) {
	a := []struct {
		s1, s2 string
	}{
		{s1: "", s2: ""},
		{s1: "{{", s2: "{"},
		{s1: "}}", s2: "}"},
		{s1: "%", s2: "%%"},
		{s1: "{0}", s2: "%[1]v"},
		{s1: "{1}, {0}", s2: "%[2]v, %[1]v"},
		{s1: "%{1}, {{{0}}}", s2: "%%%[2]v, {%[1]v}"},
		{s1: "{0:0#}", s2: "%[1]v"},
		{s1: "{0:0#.00##}", s2: "%.4[1]f"},
		{s1: "{0:.00}", s2: "%.2[1]f"},
	}

	for i, v := range a {
		if r := convertFormat(v.s1); r != v.s2 {
			t.Fatalf("unexpected result. result: %q, expected: %q, i: %d", r, v.s2, i)
		}
	}
}
