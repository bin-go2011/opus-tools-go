package main

import (
	"fmt"
	"ogg"
	"os"
)

const CHUNK = 4500

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	ogsync := ogg.SyncState{}
	page := ogg.Page{}

	ogsync.Init()
	defer ogsync.Clear()

	ogsync.PageSeek(&page)

	bytes := ogsync.NewBuffer(CHUNK)

	n, err := file.Read(bytes)
	if err != nil {
		panic(err)
	}
	ogsync.Wrote(n)
	fmt.Printf("%#v\n", ogsync)
}
