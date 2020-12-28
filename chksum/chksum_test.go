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

package chksum

import (
	"hash/crc32"
	"sync"
	"testing"

	"go.uber.org/atomic"
)

const str = "01234567890123456789012345678901234567890123456789012345678901234567890123456789"

func TestChecksum32_Threadsafe(t *testing.T) {
	buf := []byte(str)
	h := Int31(buf)

	var ok atomic.Bool
	ok.Store(true)

	const total = 1000
	var wg sync.WaitGroup
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				if Int31(buf) != h {
					ok.Store(false)
					return
				}
			}
		}()
	}

	wg.Wait()
	if !ok.Load() {
		t.Fatal("chksum.Int31 is not threadsafe")
	}
}

func TestChecksum64_Threadsafe(t *testing.T) {
	buf := []byte(str)
	h := Int63(buf)

	var ok atomic.Bool
	ok.Store(true)

	const total = 1000
	var wg sync.WaitGroup
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				if Int63(buf) != h {
					ok.Store(false)
					return
				}
			}
		}()
	}

	wg.Wait()
	if !ok.Load() {
		t.Fatal("chksum.Int63 is not threadsafe")
	}
}

func BenchmarkChecksum(b *testing.B) {
	b.Run("chksum.Int31", func(b *testing.B) {
		b.ReportAllocs()
		buf := []byte(str)
		for i := 0; i < b.N; i++ {
			Int31(buf)
		}
	})

	b.Run("chksum.Int63", func(b *testing.B) {
		b.ReportAllocs()
		buf := []byte(str)
		for i := 0; i < b.N; i++ {
			Int63(buf)
		}
	})

	b.Run("CRC32", func(b *testing.B) {
		b.ReportAllocs()
		buf := []byte(str)
		tab := crc32.MakeTable(0x06B80873)
		for i := 0; i < b.N; i++ {
			crc32.Checksum(buf, tab)
		}
	})
}
