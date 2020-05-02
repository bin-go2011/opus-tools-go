package main

import (
	"syscall"
	"unsafe"
)

var msvcrtHandle syscall.Handle
var (
	callocPtr uintptr
)

func init() {
	msvcrtHandle, err := syscall.LoadLibrary("msvcrt.dll")
	if err != nil {
		panic("couldn't load msvcrt.dll")
	}
	callocPtr, err = syscall.GetProcAddress(msvcrtHandle, "calloc")
	if err != nil {
		panic("couldn't get calloc function")
	}
}

func byteSliceToString(bval []byte) string {
	for i := range bval {
		if bval[i] == 0 {
			return string(bval[:i])
		}
	}
	return string(bval[:])
}

func bytePtrToString(r uintptr) string {
	if r == 0 {
		return ""
	}
	bval := (*[1 << 10]byte)(unsafe.Pointer(r))
	return byteSliceToString(bval[:])
}
