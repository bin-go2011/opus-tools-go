package opusfile

import (
	"fmt"
	"testing"
)

const SAMPLE_FILE = "../data/big_buck_bunny.opus"

func TestAllPages(t *testing.T) {
	f, _ := Open(SAMPLE_FILE)
	defer f.Close()

	for p := range f.Pages() {
		fmt.Println(p)
	}
}
