package opusfile

import (
	"fmt"
	"testing"
)

const SAMPLE_FILE = "../data/big_buck_bunny.opus"

func TestAllPages(t *testing.T) {
	f, _ := Open(SAMPLE_FILE)
	defer f.Close()

	fmt.Printf("%#v\n", f.links[0].head)
}
