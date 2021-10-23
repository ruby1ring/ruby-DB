package model

import "bytes"

const Exit = ".exit"

type InputBuffer struct {
	Buffer *bytes.Buffer
}
