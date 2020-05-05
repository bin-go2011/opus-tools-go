package ogg

import (
	"reflect"
	"unsafe"
)

type Page struct {
	header     uintptr
	header_len int32
	body       uintptr
	body_len   int32
}

func (og *Page) Serialno() int32 {
	return oggPageSerialno(og)
}

func (og *Page) Header() []byte {
	var data []byte

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sliceHeader.Cap = int(og.header_len)
	sliceHeader.Len = int(og.header_len)
	sliceHeader.Data = og.header

	return data
}

func (og *Page) Body() []byte {
	var data []byte

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sliceHeader.Cap = int(og.body_len)
	sliceHeader.Len = int(og.body_len)
	sliceHeader.Data = og.body

	return data
}

func (og *Page) BeginningOfStream() int32 {
	return oggPageBos(og)
}

func (og *Page) EndOfStream() int32 {
	return oggPageEos(og)
}

func (og *Page) PageNo() int32 {
	return oggPageNo(og)
}
