package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func Encode(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

type BaseMethod struct {
	Method string `json:"method"`
}

func Decode(msg []byte) (string, []byte, error) {
	header, content, found := bytes.Cut(msg, []byte("\r\n\r\n"))
	if !found {
		return "", nil, errors.New("Decode not found")
	}
	contentLengthBytes := header[len("Content Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return "", nil, err
	}
	var baseMethod BaseMethod
	if err := json.Unmarshal(content[:contentLength], &baseMethod); err != nil {
		return "", nil, err
	}
	return baseMethod.Method, content[:contentLength], nil
}

// type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)

	func SplitFunc(data []byte, _ bool) (advance int, token []byte, err error) {
		header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
		if !found {
			return 0, nil, nil
		}
		contentLengthBytes := header[len("Content-Length: "):]
		contentLength, err := strconv.Atoi(string(contentLengthBytes))
		if err != nil {
			return 0, nil, err
		}
		if len(content) < contentLength {
			return 0, nil, nil
		}
		totalLength := len(header) + 4 + contentLength
		return totalLength, data[:totalLength], nil
	}