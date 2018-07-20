package utils

import "io"

type ReqIface interface {
	String() string
	SetContentType() string
	SetReqBuf([]byte) (io.Reader, error)
}
