package opusfile

import "testing"

const SAMPLE_FILE = "../data/big_buck_bunny.opus"

func TestOpenFile(t *testing.T) {
	file := Open(SAMPLE_FILE)
	defer file.Close()
}
