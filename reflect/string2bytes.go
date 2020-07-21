package main

import (
	"reflect"
	"unsafe"
)

func string2bytes(s string) []byte {
	//使用unsafe包，将string指针类型转换成非类型安全指针，再通过reflect.StringHeader转换成string底层结构指针
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&bh))
}

func bytes2string(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	sh := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}

	return *(*string)(unsafe.Pointer(&sh))
}
