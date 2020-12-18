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
