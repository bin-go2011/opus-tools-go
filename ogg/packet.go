package ogg

import (
	"reflect"
	"unsafe"
)

type Packet struct {
	packet uintptr
	bytes  int32
	b_o_s  int32
	e_o_s  int32

	granulepos int64
	packetno   int64
}

func (op *Packet) Data() []byte {
	var data []byte

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sliceHeader.Cap = int(op.bytes)
	sliceHeader.Len = int(op.bytes)
	sliceHeader.Data = op.packet

	return data
}
