package opusfile

import (
	"fmt"
	"io"
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
	c      chan ogg.Page
}

func (of *OggOpusFile) Close() {
	of.oy.Clear()
	of.stream.Close()
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

func (of *OggOpusFile) Pages() chan ogg.Page {
	if of.c == nil {
		of.c = make(chan ogg.Page, 1000)
		go of.pagesToChannel()
	}

	return of.c
}

func (of *OggOpusFile) pagesToChannel() {
	defer close(of.c)
	for {
		page, err := of.NextPage()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		of.c <- *page
	}
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
