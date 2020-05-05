package main

import (
	"encoding/binary"
)

type OpusHeader struct {
	version           int
	channels          int /* Number of channels: 1..255 */
	preskip           int
	input_sample_rate uint32
	gain              int /* in dB S7.8 should be zero whenever possible */

	// channel_mapping   int
	/* The rest is only used if channel_mapping != 0 */
	// nb_streams int
	// nb_coupled int
	// unsigned char stream_map[255];
	// unsigned char dmatrix[OPUS_DEMIXING_MATRIX_SIZE_MAX];
}

func opus_header_parse(header []byte, h *OpusHeader) int {
	if string(header[:8]) != "OpusHead" {
		return 0
	}
	h.version = int(header[8])
	h.channels = int(header[9])
	h.preskip = int(binary.LittleEndian.Uint16(header[10:]))
	h.input_sample_rate = binary.LittleEndian.Uint32(header[12:])
	h.gain = int(binary.LittleEndian.Uint16(header[16:]))

	return 1
}
