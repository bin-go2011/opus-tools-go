package main

import (
	"fmt"
	"ogg"
	"os"
	"path/filepath"
)

const CHUNK = 4500

func get_next_page(file *os.File, ogsync *ogg.SyncState, page *ogg.Page, written *int) int {
	for {
		if ret := ogsync.PageSeek(page); ret > 0 {
			break
		} else {
			if ret < 0 {
				continue
			}
		}

		bytes := ogsync.NewBuffer(CHUNK)
		if len(bytes) == 0 {
			ogsync.Wrote(0)
			return 0
		}

		n, err := file.Read(bytes)
		if err != nil {
			panic(err)
		}
		ogsync.Wrote(n)
	}

	return 1
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	path, _ := filepath.Abs(file.Name())
	fmt.Printf("Processing file \"%s\"...\n\n", path)

	ogsync := ogg.SyncState{}
	page := ogg.Page{}

	ogsync.Init()
	defer ogsync.Clear()

	written := 0
	get_next_page(file, &ogsync, &page, &written)
	fmt.Println(ogsync)
}
