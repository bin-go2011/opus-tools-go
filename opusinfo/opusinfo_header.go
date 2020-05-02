package main

import "ogg"

const (
	CHUNK                      = 4500
	CONSTRAINT_PAGE_AFTER_EOS  = 1
	CONSTRAINT_MUXING_VIOLATED = 2
)

type misc_opus_info struct {
	oh                   OpusHeader
	bytes                int64
	overhead_bytes       int64
	lastlastgranulepos   int64
	lastgranulepos       int64
	firstgranule         int64
	total_samples        int64
	total_packets        int64
	total_pages          int64
	last_packet_duration int32
	last_page_duration   int32
	max_page_duration    int32
	min_page_duration    int32
	max_packet_duration  int32
	min_packet_duration  int32
	max_packet_bytes     int32
	min_packet_bytes     int32
	last_eos             int

	doneheaders int
}

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
	data   *misc_opus_info
}

type stream_set struct {
	streams    []*stream_processor
	in_headers int32
}

func create_stream_set() *stream_set {
	return &stream_set{}
}
