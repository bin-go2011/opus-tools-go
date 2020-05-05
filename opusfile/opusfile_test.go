package opusfile

import (
	"fmt"
	"testing"
)

const SAMPLE_FILE = "../data/big_buck_bunny.opus"

func TestAllPages(t *testing.T) {
	f, _ := Open(SAMPLE_FILE)
	defer f.Close()

	for {
		page, err := f.NextPage()
		if err != nil {
			break
		}
		fmt.Println(page)
		fmt.Println(string(page.Header()[:4]))
	}
}
