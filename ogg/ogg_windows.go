package ogg

import (
	"reflect"
	"unsafe"

	"golang.org/x/sys/windows"
)

var oggDLL *windows.DLL

var (
	oggSyncInitProc,
	oggSyncClearProc,
	oggSyncDestroyProc,
	oggSyncPageSeekProc,
	oggSyncBufferProc,
	oggSyncWroteProc,
	oggPageSerialnoProc *windows.Proc
)

func init() {
	oggDLL = windows.MustLoadDLL("../lib/ogg.dll")

	oggSyncInitProc = oggDLL.MustFindProc("ogg_sync_init")
	oggSyncClearProc = oggDLL.MustFindProc("ogg_sync_clear")
	oggSyncDestroyProc = oggDLL.MustFindProc("ogg_sync_destroy")

	oggSyncPageSeekProc = oggDLL.MustFindProc("ogg_sync_pageseek")
	oggSyncBufferProc = oggDLL.MustFindProc("ogg_sync_buffer")
	oggSyncWroteProc = oggDLL.MustFindProc("ogg_sync_wrote")
	oggPageSerialnoProc = oggDLL.MustFindProc("ogg_page_serialno")
}

func oggSyncInit(oy *SyncState) int32 {
	r1, _, _ := oggSyncInitProc.Call(uintptr(unsafe.Pointer(oy)))
	return int32(r1)
}

func oggSyncClear(oy *SyncState) {
	oggSyncClearProc.Call(uintptr(unsafe.Pointer(oy)))
}

func oggSyncDestroy(oy *SyncState) {
	oggSyncDestroyProc.Call(uintptr(unsafe.Pointer(oy)))
}

func oggSyncPageSeek(oy *SyncState, og *Page) int32 {
	r1, _, _ := oggSyncPageSeekProc.Call(uintptr(unsafe.Pointer(oy)), uintptr(unsafe.Pointer(og)))
	return int32(r1)
}

func oggSyncBuffer(oy *SyncState, size int) []byte {
	r1, _, _ := oggSyncBufferProc.Call(uintptr(unsafe.Pointer(oy)), uintptr(size))

	var bytes []byte
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	sliceHeader.Cap = size
	sliceHeader.Len = size
	sliceHeader.Data = r1

	return bytes
}

func oggSyncWrote(oy *SyncState, bytes int) int32 {
	r1, _, _ := oggSyncWroteProc.Call(uintptr(unsafe.Pointer(oy)), uintptr(bytes))
	return int32(r1)
}

func oggPageSerialno(og *Page) int32 {
	r1, _, _ := oggPageSerialnoProc.Call(uintptr(unsafe.Pointer(og)))
	return int32(r1)
}
