package wunsafe

import (
	"reflect"
	"unsafe"
)

func BytesToString(bts []byte) string {
	return *(*string)(unsafe.Pointer(&bts))
}

func StringToBytes(str string) []byte {
	strLen := len(str)
	ptr := unsafe.Pointer(&str)
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*reflect.StringHeader)(ptr).Data,
		Len:  strLen,
		Cap:  strLen,
	}))
}
