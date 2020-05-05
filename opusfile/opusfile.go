package opusfile

import (
	"fmt"
	"ogg"
	"os"
)

const (
	CHUNK = 4500
)

type OpusHead struct {
}

type OpusTags struct {
}

type OggOpusFile struct {
	oy     ogg.SyncState
	stream *os.File
}

func (of *OggOpusFile) Close() {
	of.stream.Close()
	of.oy.Destroy()
}

func Open(file string) (*OggOpusFile, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	of := &OggOpusFile{
		stream: f,
	}
	of.oy.Init()

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
