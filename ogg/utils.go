package ogg

import "unsafe"

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
