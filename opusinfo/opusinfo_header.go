package main

import "ogg"

type stream_processor struct {
	process_page        func(*stream_processor, *ogg.Page)
	process_end         func(*stream_processor)
	isillegal           int32
	constraint_violated int32
	shownillegal        int32
	isnew               int32
	seqno               int32
	lostseq             int32
	seen_file_icons     int32

	start int32
	end   int32

	num int
	typ string

	serial uint32 /* must be 32 bit unsigned */
	os     ogg.StreamState
	data   uintptr
}

type stream_set struct {
	streams    []stream_processor
	in_headers int32
}
