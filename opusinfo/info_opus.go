package main

import "ogg"

func info_opus_end(stream *stream_processor) {

}

func info_opus_process(stream *stream_processor, page *ogg.Page) {
	var packet ogg.Packet
	var h OpusHeader

	streamState := &(stream.os)
	streamState.Pagein(page)

	for {
		res := streamState.Packetout(&packet)
		if res == 0 {
			break
		}
		opus_header_parse(packet.Data(), &h)
	}

}

func info_opus_start(stream *stream_processor) {
	stream.typ = "opus"
	stream.process_page = info_opus_process
	stream.process_end = info_opus_end

	stream.data = &misc_opus_info{
		firstgranule:        1,
		min_packet_duration: 5760,
		min_page_duration:   5760 * 255,
		min_packet_bytes:    2147483647,
	}
}
