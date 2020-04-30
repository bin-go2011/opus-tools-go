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

type Page struct {
	header     uintptr
	header_len int32
	body       uintptr
	body_len   int32
}

func (og *Page) Serialno() int32 {
	return oggPageSerialno(og)
}

func (og *Page) Header() string {
	return bytePtrToString(og.header)
}

func (og *Page) Body() string {
	return bytePtrToString(og.body)
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

type Packet struct {
	packet uintptr
	bytes  int32
	b_o_s  int32
	e_o_s  int32

	granulepos int64
	packetno   int64
}

func (op *Packet) Packet() uintptr {
	return op.packet
}

func (op *Packet) Bytes() int32 {
	return op.bytes
}

type StreamState struct {
	body_data     uintptr /* bytes from packet bodies */
	body_storage  int32   /* storage elements allocated */
	body_fill     int32   /* elements stored; fill mark */
	body_returned int32   /* elements of fill returned */

	lacing_vals  *int32 /* The values that will go to the segment table */
	granule_vals int64  /* granulepos values for headers. Not compact
	this way, but it is simple coupled to the
	lacing fifo */
	lacing_storage  int32
	lacing_fill     int32
	lacing_packet   int32
	lacing_returned int32

	header      [282]uint8 /* working space for header encode */
	header_fill int32

	e_o_s int32 /* set when we have buffered the last packet in the
	   logical bitstream */
	b_o_s int32 /* set after we've written the initial page
	   of a logical bitstream */
	serialno int32
	pageno   int32
	packetno int64 /* sequence number for decode; the framing
	   knows where there's a hole in the data,
	   but we need coupling so that the codec
	   (which is in a separate abstraction
	   layer) also knows about the gap */
	granulepos int64
}

func (os *StreamState) Init(serialno int) int32 {
	return oggStreamInit(os, serialno)
}

func (os *StreamState) Pagein(page *Page) int32 {
	return oggStreamPagein(os, page)
}

func (os *StreamState) Packetout(op *Packet) int32 {
	return oggStreamPacketout(os, op)
}

func (os *StreamState) Clear() {
	oggStreamClear(os)
}
