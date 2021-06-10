package util

import (
	"reflect"
	"strings"
	"unsafe"
)

func StringIsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func StringIsNotEmpty(s string) bool {
	return len(strings.TrimSpace(s)) > 0
}

func StringLength(s string) int {
	return len(strings.TrimSpace(s))
}

func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

func StringToBytes(s string) (bs []byte) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return
}
