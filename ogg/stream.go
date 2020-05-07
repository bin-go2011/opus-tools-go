package ogg

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

func (os *StreamState) Destroy() {
	oggStreamDestroy(os)
}

func (os *StreamState) ResetSerialno(n int) int32 {
	return oggStreamResetSerialno(os, n)
}

func (os *StreamState) Serialno() int32 {
	return os.serialno
}
