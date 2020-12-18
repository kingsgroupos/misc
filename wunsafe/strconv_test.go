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
