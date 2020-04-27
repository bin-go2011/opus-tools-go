package ogg

type SyncState struct {
	data     uintptr
	storage  int32
	fill     int32
	returned int32

	unsynced    int32
	headerbytes int32
	bodybytes   int32
}

type Page struct {
	header     uintptr
	header_len int32
	body       uintptr
	body_len   int32
}

func (oy *SyncState) Init() int32 {
	return oggSyncInit(oy)
}

func (oy *SyncState) Clear() {
	oggSyncClear(oy)
}

func (oy *SyncState) Destroy() {
	oggSyncDestroy(oy)
}

func (oy *SyncState) PageSeek(og *Page) int32 {
	return oggSyncPageSeek(oy, og)
}

func (oy *SyncState) NewBuffer(size int) []byte {
	return oggSyncBuffer(oy, size)
}

func (oy *SyncState) Wrote(bytes int) int32 {
	return oggSyncWrote(oy, bytes)
}
