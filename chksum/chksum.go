package chksum

import (
	"math"

	"github.com/pierrec/xxHash/xxHash32"
	"github.com/pierrec/xxHash/xxHash64"
)

func Int31(bts []byte) int32 {
	return int32(xxHash32.Checksum(bts, 0x06B80873) & math.MaxInt32)
}

func Int63(bts []byte) int64 {
	return int64(xxHash64.Checksum(bts, 0x06B8087306B80873) & math.MaxInt64)
}
