package ogg

type Page struct {
	header     uintptr
	header_len int32
	body       uintptr
	body_len   int32
}

func (og *Page) Serialno() int32 {
	return oggPageSerialno(og)
}

func (og *Page) Header() string {
	return bytePtrToString(og.header)
}

func (og *Page) Body() string {
	return bytePtrToString(og.body)
}

func (og *Page) BeginningOfStream() int32 {
	return oggPageBos(og)
}

func (og *Page) EndOfStream() int32 {
	return oggPageEos(og)
}

func (og *Page) PageNo() int32 {
	return oggPageNo(og)
}
