package utils

type RespIface interface {
	String() string
	SetText(text string) error
}
