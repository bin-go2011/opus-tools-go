package main

import (
	"fmt"
	"ogg"
	"os"
	"path/filepath"
)

func find_stream_processor(set *stream_set, page *ogg.Page) *stream_processor {
	var invalid int32
	var constraint int32

	serial := uint32(page.Serialno())

	for _, stream := range set.streams {
		if serial == stream.serial {
			set.in_headers = 0

			if stream.end > 0 {
				stream.isillegal = 1
				stream.constraint_violated = CONSTRAINT_PAGE_AFTER_EOS
				return stream
			}

			stream.isnew = 0
			stream.start = page.BeginningOfStream()
			stream.end = page.EndOfStream()
			stream.serial = serial

			return stream
		}
	}

	set.in_headers = 1

	stream := &stream_processor{
		isnew:               1,
		isillegal:           invalid,
		constraint_violated: constraint,
		seen_file_icons:     0,
	}
	set.streams = append(set.streams, stream)
	stream.num = len(set.streams)

	{
		var packet ogg.Packet
		streamState := &(stream.os)
		streamState.Init(int(serial))
		streamState.Pagein(page)
		res := streamState.Packetout(&packet)
		if res > 0 {
			data := packet.Data()
			if len(data) >= 19 && byteSliceToString(data)[:8] == "OpusHead" {
				info_opus_start(stream)
			}
		}

		streamState.Clear()
		streamState.Init(int(serial))
	}
	stream.start = page.BeginningOfStream()
	stream.end = page.EndOfStream()
	stream.serial = serial
	stream.shownillegal = 0
	stream.seqno = page.PageNo()

	return stream
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
	ogsync.Init()
	defer ogsync.Clear()

	page := ogg.Page{}

	processors := create_stream_set()
	written := 0

	for i := 0; i < 2; i++ {
		get_next_page(file, &ogsync, &page, &written)

		p := find_stream_processor(processors, &page)

		if p.isnew > 0 {
			fmt.Printf("New logical stream (#%d, serial: %08x): type %s\n",
				p.num, p.serial, p.typ)
		}

		// if p.isillegal == 0 {
		// 	p.process_page(p, &page)

		// 	if p.end > 0 {
		// 		p.process_end(p)
		// 		fmt.Printf("Logical stream %d ended\n", p.num)
		// 	}
		// }

	}
}

func main() {
	process_file(os.Args[1])
}
