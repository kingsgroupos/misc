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

package variable

import "testing"

func TestToCamel(t *testing.T) {
	cases := []struct {
		input  string
		output string
	}{
		{"", ""},
		{"test", "test"},
		{"Test", "test"},
		{"TEST", "test"},
		{"testCase", "testCase"},
		{"TestCase", "testCase"},
		{"TESTCase", "testCase"},
		{"test_case", "testCase"},
		{"test_Case", "testCase"},
		{"Test_case", "testCase"},
		{"Test_Case", "testCase"},
		{"TEST_case", "testCase"},
		{"TEST_Case", "testCase"},
		{"test1", "test1"},
		{"Test1", "test1"},
		{"TEST1", "test1"},
		{"test1x", "test1X"},
		{"test_1", "test1"},
		{"Test_1", "test1"},
		{"TEST_1", "test1"},
		{"test12", "test12"},
		{"Test12", "test12"},
		{"TEST12", "test12"},
		{"test_12", "test12"},
		{"Test_12", "test12"},
		{"TEST_12", "test12"},
		{"test123", "test_123"},
		{"Test123", "test_123"},
		{"TEST123", "test_123"},
		{"test_123", "test_123"},
		{"Test_123", "test_123"},
		{"TEST_123", "test_123"},
		{"_", ""},
		{"中文", ""},
		{"_a", "a"},
		{"a_", "a"},
		{"_a_", "a"},
		{"__a", "a"},
		{"a__", "a"},
		{"__a__", "a"},
		{"__a__b", "aB"},
		{"__a__bc__", "aBc"},
		{"__a__bC__", "aBC"},
		{"__a__Bc__", "aBc"},
		{"__a__BC__", "aBC"},
		{"1", ""},
		{"_1", ""},
		{"1a", "a"},
		{"_1a", "a"},
	}

	for i, c := range cases {
		if o := ToCamel(c.input); o != c.output {
			t.Fatalf("ToCamel does not work as expected. i: %d, input: %s, expected: %s, actual: %s",
				i, c.input, c.output, o)
		}
	}
}
