package raspivid

import (
	"bufio"
	"os"
	"testing"
)

func Test_bytesIndex(t *testing.T) {
	file, err := os.Open("out.h264")
	if err != nil {
		t.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(scanNalSeparator)

	for scanner.Scan() {
		t.Log(scanner.Bytes())
	}
}
