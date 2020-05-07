package opusfile

import (
	"fmt"
	"ogg"
	"os"
)

const (
	CHUNK = 4500
)

const (
	/*Initial state.*/
	OP_NOTOPEN = 0
	/*We've found the first Opus stream in the first link.*/
	OP_PARTOPEN = 1
	OP_OPENED   = 2
	/*We've found the first Opus stream in the current link.*/
	OP_STREAMSE = 3
	/*We've initialized the decoder for the chosen Opus stream in the current
	  link.*/
	OP_INITSET = 4
)

type OggOpusLink struct {
	/*The serial number.*/
	serialno uint32
	/*The contents of the info header.*/
	head OpusHead
	tags OpusTags
}

type OggOpusFile struct {
	oy          ogg.SyncState
	os          ogg.StreamState
	stream      *os.File
	ready_state int
	links       []OggOpusLink
}

func (of *OggOpusFile) Close() {
	of.stream.Close()
	of.oy.Destroy()
	of.os.Destroy()
}

func Open(file string) (*OggOpusFile, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	of := &OggOpusFile{
		stream:      f,
		ready_state: OP_NOTOPEN,
	}
	of.oy.Init()
	of.os.Init(-1)

	of.fetch_headers()
	return of, nil
}

func (of *OggOpusFile) NextPage() (*ogg.Page, error) {
	page := &ogg.Page{}
	var written int

	for {
		if ret := of.oy.PageSeek(page); ret > 0 {
			break
		} else {
			if ret < 0 {
				fmt.Printf("Hole in data (%d bytes) found at approximate offset %d bytes. Corrupted Ogg.\n", -ret, written)
				continue
			}
		}

		bytes := of.oy.NewBuffer(CHUNK)
		if len(bytes) == 0 {
			of.oy.Wrote(0)
			err := fmt.Errorf("ogg_sync_buffer allocates bytes=%d", bytes)
			return nil, err
		}

		n, err := of.stream.Read(bytes)
		if err != nil {
			return nil, err
		}
		of.oy.Wrote(n)
		written += n

	}

	return page, nil
}
