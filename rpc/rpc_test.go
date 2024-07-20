package rpc

import "testing"

type EncodingExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := Encode(EncodingExample{Testing: true})
	if expected != actual {
		t.Fatalf("expected %s actual %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	msg := "Content-Length: 16\r\n\r\n{\"Method\":\"yes\"}"
	method, content, err := Decode([]byte(msg))
	if err != nil {
		t.Fatal(err)
	}
	if method != "yes" {
		t.Fatalf("Expected : hi got %s", method)
	}
	contentLength := len(content)
	if contentLength != 16 {
		t.Fatalf("expected 14 actual %d", contentLength)
	}
}
