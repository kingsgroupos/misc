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
	"testing"
)

func TestSplit(t *testing.T) {
	inputs := []string{
		"hello world 123",
		"   hello world 123",
		"hello world 123   ",
		"hello  world  123",
	}

	for _, str := range inputs {
		a := Split(str, " ")
		if a[0] != "hello" || a[1] != "world" || a[2] != "123" {
			t.Fatalf("something is wrong with misc.Split. input: %q, result: %s", str, ToJSON(a))
		}
	}
}

func TestUnescapeUnicode(t *testing.T) {
	str1 := `Journey to the \u300EWest\u300F`
	if UnescapeUnicode(str1) != "Journey to the 『West』" {
		t.Fatalf("failed to unquote the following string: %s", str1)
	}

	str2 := `Journey to the \u300EWest\u300F - Wu Chengen`
	if UnescapeUnicode(str2) != "Journey to the 『West』 - Wu Chengen" {
		t.Fatalf("failed to unquote the following string: %s", str2)
	}
}

func TestUCFirst(t *testing.T) {
	strs1 := []string{
		"a",
		"A",
		"abc",
		"Abc",
		"程序员",
		"a程",
		"A程",
	}
	strs2 := []string{
		"A",
		"A",
		"Abc",
		"Abc",
		"程序员",
		"A程",
		"A程",
	}

	if len(strs1) != len(strs2) {
		t.Fatalf("strs1 and strs2 are not the same length. len1: %d, len2: %d", len(strs1), len(strs2))
	}
	for i := 0; i < len(strs1); i++ {
		if UCFirst(strs1[i]) != strs2[i] {
			t.Fatalf("UCFirst does not work. i: %d, s1: %s, s2: %s", i, strs1[i], strs2[i])
		}
	}
}

func TestLCFirst(t *testing.T) {
	strs1 := []string{
		"a",
		"A",
		"abc",
		"Abc",
		"程序员",
		"a程",
		"A程",
	}
	strs2 := []string{
		"a",
		"a",
		"abc",
		"abc",
		"程序员",
		"a程",
		"a程",
	}

	if len(strs1) != len(strs2) {
		t.Fatalf("strs1 and strs2 are not the same length. len1: %d, len2: %d", len(strs1), len(strs2))
	}
	for i := 0; i < len(strs1); i++ {
		if LCFirst(strs1[i]) != strs2[i] {
			t.Fatalf("LCFirst does not work. i: %d, s1: %s, s2: %s", i, strs1[i], strs2[i])
		}
	}
}
