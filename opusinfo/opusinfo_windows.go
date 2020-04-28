package main

import (
	"ogg"
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

type stream_processor struct {
	process_page        uintptr
	process_end         uintptr
	isillegal           int32
	constraint_violated int32
	shownillegal        int32
	isnew               int32
	seqno               int32
	lostseq             int32
	seen_file_icons     int32

	start int32
	end   int32

	num int32
	typ *uint8

	serial uint32 /* must be 32 bit unsigned */
	os     ogg.StreamState
	data   uintptr
}

type stream_set struct {
	streams    []stream_processor
	in_headers int32
}

func create_stream_set() *stream_set {
	return &stream_set{}
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
