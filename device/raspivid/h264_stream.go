package raspivid

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"log"
)

const (
	readBufferSize = 4096
	bufferSizeKB   = 256
)

// nalSeparator NAL break
var nalSeparator = []byte{0, 0, 0, 1}

func SendVideoStream(ctx context.Context, opt Options, sender chan []byte) {
	var (
		stdoutPipe io.ReadCloser
		err        error
	)
	cmd := makeCmd(ctx, opt)
	defer func() {
		if err != nil {
			log.Printf("SendVideoStream Run error: %s\n", err)
		}
		cmd.Wait()
		close(sender)
		log.Println("SendVideoStream Stop [raspivid]")
	}()

	stdoutPipe, err = cmd.StdoutPipe()
	if err != nil {
		return
	}
	defer stdoutPipe.Close()
	if err = cmd.Start(); err != nil {
		return
	}
	log.Println("SendVideoStream Started: ", cmd.Args)

	scanner := bufio.NewScanner(stdoutPipe)
	scanner.Split(scanNalSeparator)
LOOP:
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			buf := scanner.Bytes()
			if len(buf) == 0 {
				continue
			}
			data := make([]byte, 0, len(buf)+4)
			data = append(data, nalSeparator...)
			sender <- append(data, buf...)
		}
	}
}

func scanNalSeparator(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, nalSeparator); i >= 0 {
		return i + len(nalSeparator), data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}
