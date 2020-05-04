package opus

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var opusDLL *windows.DLL

var (
	opusMultistreamDecoderCreateProc *windows.Proc
)

func init() {
	opusDLL = windows.MustLoadDLL("../lib/opus.dll")
	opusMultistreamDecoderCreateProc = opusDLL.MustFindProc("opus_multistream_decoder_create")
}

func opusMultistreamDecoderCreate(Fs int32,
	channels int,
	streams int,
	coupled_streams int,
	// const unsigned char *mapping,
	error *int) {
	opusMultistreamDecoderCreateProc.Call(uintptr(Fs), uintptr(channels), uintptr(streams), uintptr(coupled_streams), uintptr(unsafe.Pointer(error)))
}
