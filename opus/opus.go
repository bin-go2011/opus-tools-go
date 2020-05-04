package opus

func MultistreamDecoderCreate(Fs int32,
	channels int,
	streams int,
	coupled_streams int,
	// const unsigned char *mapping,
	error *int) {
	opusMultistreamDecoderCreate(Fs, channels, streams, coupled_streams, error)
}
