package opusfile

import (
	"encoding/binary"
	"fmt"
	"ogg"
)

const (
	/**A request did not succeed.*/
	OP_FALSE = -1
	/*Currently not used externally.*/
	OP_EOF = -2
	/**There was a hole in the page sequence numbers (e.g., a page was corrupt or
	  missing).*/
	OP_HOLE = -3
	/**An underlying read, seek, or tell operation failed when it should have
	  succeeded.*/
	OP_EREAD = -128
	/**A <code>NULL</code> pointer was passed where one was unexpected, or an
	  internal memory allocation failed, or an internal library error was
	  encountered.*/
	OP_EFAULT = -129
	/**The stream used a feature that is not implemented, such as an unsupported
	  channel family.*/
	OP_EIMPL = -130
	/**One or more parameters to a function were invalid.*/
	OP_EINVAL = -131
	/**A purported Ogg Opus stream did not begin with an Ogg page, a purported
	  header packet did not start with one of the required strings, "OpusHead" or
	  "OpusTags", or a link in a chained file was encountered that did not
	  contain any logical Opus streams.*/
	OP_ENOTFORMAT = -132
	/**A required header packet was not properly formatted, contained illegal
	  values, or was missing altogether.*/
	OP_EBADHEADER = -133
	/**The ID header contained an unrecognized version number.*/
	OP_EVERSION = -134
	/*Currently not used at all.*/
	OP_ENOTAUDIO = -135
	/**An audio packet failed to decode properly.
	  This is usually caused by a multistream Ogg packet where the durations of
	   the individual Opus packets contained in it are not all the same.*/
	OP_EBADPACKET = -136
	/**We failed to find data we had seen before, or the bitstream structure was
	  sufficiently malformed that seeking to the target destination was
	  impossible.*/
	OP_EBADLINK = -137
	/**An operation that requires seeking was requested on an unseekable stream.*/
	OP_ENOSEEK = -138
	/**The first or last granule position of a link failed basic validity checks.*/
	OP_EBADTIMESTAMP = -139
)

type OpusHead struct {
	version           int
	channels          int /* Number of channels: 1..255 */
	preskip           int
	input_sample_rate uint32
	gain              int /* in dB S7.8 should be zero whenever possible */
}

func (of *OggOpusFile) fetch_headers() {
	page, err := of.NextPage()
	if err != nil {
		panic(err)
	}

	op := ogg.Packet{}
	of.links = append(of.links, OggOpusLink{})

	streamState := of.os
	streamState.ResetSerialno(int(page.Serialno()))
	streamState.Pagein(page)
	streamState.Packetout(&op)

	res := opus_head_parse(op.Data(), &(of.links[0].head))
	if res < 0 {
		err := fmt.Errorf("failed to fetch header, error code=%d", res)
		panic(err)
	}

	page, err = of.NextPage()
	if err != nil {
		panic(err)
	}
	streamState.Pagein(page)
	streamState.Packetout(&op)
	opus_tags_parse(op.Data(), &(of.links[0].tags))
}

func opus_head_parse(data []byte, head *OpusHead) int {
	_len := len(data)
	if _len < 8 {
		return OP_ENOTFORMAT
	}

	if string(data[:8]) != "OpusHead" {
		return OP_ENOTFORMAT
	}

	if _len < 9 {
		return OP_EBADHEADER
	}

	head.version = int(data[8])
	if head.version > 15 {
		return OP_EVERSION
	}

	if _len < 19 {
		return OP_EBADHEADER
	}

	head.channels = int(data[9])
	head.preskip = int(binary.LittleEndian.Uint16(data[10:]))
	head.input_sample_rate = binary.LittleEndian.Uint32(data[12:])
	head.gain = int(binary.LittleEndian.Uint16(data[16:]))

	return 0
}
