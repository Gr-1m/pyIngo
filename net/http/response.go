package http

import (
	"errors"
	"net/http"
)

type Response struct {
	http.Response
}

type RespBody struct {
	BodyLen  int64
	BodyText []byte

	writeIndex int
}

func (rb *RespBody) makeCap() {
	rb.BodyText = make([]byte, rb.BodyLen, rb.BodyLen)
	rb.writeIndex = 0
}

func (rb *RespBody) MakeCap(length int64) {
	rb.BodyLen = length
	rb.makeCap()
}

func (rb *RespBody) Write(p []byte) (int, error) {
	if rb.BodyLen == 0 {
		return 0, errors.New("BodyLen is Zer0")
	}
	if cap(rb.BodyText) == 0 {
		rb.makeCap()
	}

	rb.writeIndex = copy(rb.BodyText[rb.writeIndex:], p)
	return len(p), nil
}

func (rb *RespBody) Close() error {

	if cap(rb.BodyText) == 0 {
		return nil
	}

	rb.BodyText = make([]byte, 0)
	rb.BodyLen = 0
	rb.writeIndex = 0

	return nil
}
