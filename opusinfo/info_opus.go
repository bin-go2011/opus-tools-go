package main

import (
	"encoding/binary"
	"fmt"
	"ogg"
)

func info_opus_end(stream *stream_processor) {
	fmt.Printf("Opus stream %d:\n", stream.num)
}

func info_opus_process(stream *stream_processor, page *ogg.Page) {
	var packet ogg.Packet
	var h OpusHeader

	streamState := &(stream.os)
	streamState.Pagein(page)

	inf := stream.data
	inf.last_eos = int(page.EndOfStream())

	for {
		res := streamState.Packetout(&packet)
		if res == 0 {
			break
		}

		if inf.doneheaders < 2 {
			if inf.doneheaders == 0 {
				opus_header_parse(packet.Data(), &h)
			} else if inf.doneheaders == 1 {
				data := packet.Data()
				if len(data) < 8 || string(data)[:8] != "OpusTags" {
					err := fmt.Errorf("Could not decode OpusTags header packet %d - invalid Opus stream (%d)\n",
						inf.doneheaders, stream.num)
					panic(err)
				}

				c := 8
				len := int(binary.LittleEndian.Uint32(data[c:]))

				c += 4
				fmt.Printf("Encoded with %s\n", string(data[c:c+len]))

				c += len
				nb_fields := int(binary.LittleEndian.Uint32(data[c:]))

				if nb_fields > 0 {
					c += 4
					fmt.Println("User comments section follows...")
				}

				for i := 0; i < nb_fields; i++ {
					len = int(binary.LittleEndian.Uint32(data[c:]))
					c += 4
					fmt.Printf("\t%s\n", string(data[c:c+len]))
				}
			}

			inf.doneheaders += 1
		}
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
