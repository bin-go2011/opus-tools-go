package opusfile

import (
	"ogg"
)

type OpusHead struct {
}

type OpusTags struct {
}

type OggOpusFile struct {
	oy ogg.SyncState
}

func (of *OggOpusFile) Close() {
	of.oy.Clear()
}

func Open(file string) *OggOpusFile {
	of := &OggOpusFile{}
	of.oy.Init()

	return of
}
