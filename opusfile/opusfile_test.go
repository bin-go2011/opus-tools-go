package opusfile

import "testing"

const SAMPLE_FILE = "../data/big_buck_bunny.opus"

func TestOpenFile(t *testing.T) {
	file, err := Open(SAMPLE_FILE)
	if err != nil {
		t.Error(err)
		return
	}
	defer file.Close()
}
