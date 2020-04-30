package main

import (
	"fmt"
	"ogg"
	"os"
	"path/filepath"
)

const CHUNK = 4500

func find_stream_processor(set *stream_set, page *ogg.Page) *stream_processor {
	var invalid int32
	var constraint int32

	serial := uint32(page.Serialno())

	for _, stream := range set.streams {
		if serial == stream.serial {

		}
	}

	set.in_headers = 1

	stream := stream_processor{
		isnew:               1,
		isillegal:           invalid,
		constraint_violated: constraint,
		seen_file_icons:     0,
	}
	set.streams = append(set.streams, stream)

	{
		var packet ogg.Packet
		streamState := &(stream.os)
		streamState.Init(int(serial))
		streamState.Pagein(page)
		res := streamState.Packetout(&packet)
		if res <= 0 {

		} else if packet.Bytes() >= 19 && bytePtrToString(packet.Packet())[:8] == "OpusHead" {
			info_opus_start(&stream)
		}

		res = streamState.Packetout(&packet)
		streamState.Clear()
		streamState.Init(int(serial))
	}
	stream.start = page.BeginningOfStream()
	stream.end = page.EndOfStream()
	stream.seqno = page.PageNo()

	fmt.Printf("%#v\n", stream)
	return nil
}

func get_next_page(file *os.File, ogsync *ogg.SyncState, page *ogg.Page, written *int) int {
	for {
		if ret := ogsync.PageSeek(page); ret > 0 {
			break
		} else {
			if ret < 0 {
				continue
			}
		}

		bytes := ogsync.NewBuffer(CHUNK)
		if len(bytes) == 0 {
			ogsync.Wrote(0)
			return 0
		}

		n, err := file.Read(bytes)
		if err != nil {
			panic(err)
		}
		ogsync.Wrote(n)
		*written += n
	}

	return 1
}

func process_file(name string) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	path, _ := filepath.Abs(file.Name())
	fmt.Printf("Processing file \"%s\"...\n\n", path)

	ogsync := ogg.SyncState{}
	page := ogg.Page{}

	ogsync.Init()
	defer ogsync.Clear()

	written := 0
	get_next_page(file, &ogsync, &page, &written)

	processors := create_stream_set()
	find_stream_processor(processors, &page)

}

func main() {
	process_file(os.Args[1])
}
