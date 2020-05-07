package opusfile

import (
	"encoding/binary"
)

type OpusTags struct {
	/**The array of comment string vectors.*/
	user_comments []string
	/**An array of the corresponding length of each vector, in bytes.*/
	comment_lengths []int
	/**The total number of comment streams.*/
	comments int
	/**The null-terminated vendor string.
	  This identifies the software used to encode the stream.*/
	vendor string
}

func opus_tags_parse(data []byte, tags *OpusTags) int {
	_len := len(data)
	if _len < 8 {
		return OP_ENOTFORMAT
	}
	if string(data[:8]) != "OpusTags" {
		return OP_ENOTFORMAT
	}
	if _len < 16 {
		return OP_EBADHEADER
	}

	c := 8
	count := int(binary.LittleEndian.Uint32(data[c:]))

	c += 4
	tags.vendor = string(data[c : c+count])

	c += count
	tags.comments = int(binary.LittleEndian.Uint32(data[c:]))

	c += 4
	for i := 0; i < tags.comments; i++ {
		comment_length := int(binary.LittleEndian.Uint32(data[c:]))
		tags.comment_lengths = append(tags.comment_lengths, comment_length)
		c += 4
		tags.user_comments = append(tags.user_comments, string(data[c:c+comment_length]))
	}

	return 0
}
