package opusfile

import (
	"ogg"
	"os"
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
