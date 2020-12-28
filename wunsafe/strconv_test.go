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

import (
	"bytes"
	"testing"
)

func TestBytesToString(t *testing.T) {
	if BytesToString(nil) != "" {
		t.Fatal("should return '' when the byte slice is nil")
	}

	var bts []byte
	if BytesToString(bts) != "" {
		t.Fatal("should return '' when the byte slice is empty")
	}

	bts = []byte("hello")
	if BytesToString(bts) != "hello" {
		t.Fatal("something is wrong with BytesToString")
	}
}

func TestStringToBytes(t *testing.T) {
	if !bytes.Equal(StringToBytes(""), make([]byte, 0)) {
		t.Fatal("should return an empty byte slice when the string is empty")
	}
	if !bytes.Equal(StringToBytes("hello"), []byte("hello")) {
		t.Fatal("something is wrong with StringToBytes")
	}
}

func blackHole(interface{}) {}

func BenchmarkBytesToString(b *testing.B) {
	b.Run("Unsafe", func(b *testing.B) {
		b.ReportAllocs()
		var str string
		bts := []byte("hello")
		for i := 0; i < b.N; i++ {
			str = BytesToString(bts)
		}
		blackHole(str)
	})

	b.Run("Raw", func(b *testing.B) {
		b.ReportAllocs()
		var str string
		bts := []byte("hello")
		for i := 0; i < b.N; i++ {
			str = string(bts)
		}
		blackHole(str)
	})
}

func BenchmarkStringToBytes(b *testing.B) {
	b.Run("Unsafe", func(b *testing.B) {
		b.ReportAllocs()
		var bts []byte
		str := "hello"
		for i := 0; i < b.N; i++ {
			bts = StringToBytes(str)
		}
		blackHole(bts)
	})

	b.Run("Raw", func(b *testing.B) {
		b.ReportAllocs()
		var bts []byte
		str := "hello"
		for i := 0; i < b.N; i++ {
			bts = []byte(str)
		}
		blackHole(bts)
	})
}
