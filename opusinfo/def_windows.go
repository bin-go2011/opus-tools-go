package main

import (
	"syscall"
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
