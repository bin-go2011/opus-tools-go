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
	oggPageSerialnoProc,
	oggStreamInitProc,
	oggStreamPageinProc,
	oggStreamPacketoutProc,
	oggStreamClearProc,
	oggStreamDestroyProc,
	oggStreamResetSerialnoProc,
	oggPageBosProc,
	oggPageEosProc,
	oggPageNoProc *windows.Proc
)

func init() {
	oggDLL = windows.MustLoadDLL("../lib/ogg.dll")

	oggSyncInitProc = oggDLL.MustFindProc("ogg_sync_init")
	oggSyncClearProc = oggDLL.MustFindProc("ogg_sync_clear")
	oggSyncDestroyProc = oggDLL.MustFindProc("ogg_sync_destroy")

	oggSyncPageSeekProc = oggDLL.MustFindProc("ogg_sync_pageseek")
	oggSyncBufferProc = oggDLL.MustFindProc("ogg_sync_buffer")
	oggSyncWroteProc = oggDLL.MustFindProc("ogg_sync_wrote")

	oggStreamInitProc = oggDLL.MustFindProc("ogg_stream_init")
	oggStreamPageinProc = oggDLL.MustFindProc("ogg_stream_pagein")
	oggStreamPacketoutProc = oggDLL.MustFindProc("ogg_stream_packetout")
	oggStreamClearProc = oggDLL.MustFindProc("ogg_stream_clear")
	oggStreamDestroyProc = oggDLL.MustFindProc("ogg_stream_destroy")
	oggStreamResetSerialnoProc = oggDLL.MustFindProc("ogg_stream_reset_serialno")

	oggPageSerialnoProc = oggDLL.MustFindProc("ogg_page_serialno")
	oggPageBosProc = oggDLL.MustFindProc("ogg_page_bos")
	oggPageEosProc = oggDLL.MustFindProc("ogg_page_eos")
	oggPageNoProc = oggDLL.MustFindProc("ogg_page_pageno")
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

func oggStreamInit(os *StreamState, serialno int) int32 {
	r1, _, _ := oggStreamInitProc.Call(uintptr(unsafe.Pointer(os)), uintptr(serialno))
	return int32(r1)
}

func oggStreamPagein(os *StreamState, og *Page) int32 {
	r1, _, _ := oggStreamPageinProc.Call(uintptr(unsafe.Pointer(os)), uintptr(unsafe.Pointer(og)))
	return int32(r1)
}

func oggStreamPacketout(os *StreamState, op *Packet) int32 {
	r1, _, _ := oggStreamPacketoutProc.Call(uintptr(unsafe.Pointer(os)), uintptr(unsafe.Pointer(op)))
	return int32(r1)
}

func oggStreamClear(os *StreamState) {
	oggStreamClearProc.Call(uintptr(unsafe.Pointer(os)))
}

func oggStreamDestroy(os *StreamState) {
	oggStreamDestroyProc.Call(uintptr(unsafe.Pointer(os)))
}

func oggStreamResetSerialno(os *StreamState, serialno int) int32 {
	r1, _, _ := oggStreamResetSerialnoProc.Call(uintptr(unsafe.Pointer(os)), uintptr(serialno))
	return int32(r1)
}

func oggPageSerialno(og *Page) int32 {
	r1, _, _ := oggPageSerialnoProc.Call(uintptr(unsafe.Pointer(og)))
	return int32(r1)
}

func oggPageBos(og *Page) int32 {
	r1, _, _ := oggPageBosProc.Call(uintptr(unsafe.Pointer(og)))
	return int32(r1)
}

func oggPageEos(og *Page) int32 {
	r1, _, _ := oggPageEosProc.Call(uintptr(unsafe.Pointer(og)))
	return int32(r1)
}

func oggPageNo(og *Page) int32 {
	r1, _, _ := oggPageNoProc.Call(uintptr(unsafe.Pointer(og)))
	return int32(r1)
}
