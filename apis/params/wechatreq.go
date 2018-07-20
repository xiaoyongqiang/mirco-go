package params

import (
	"net/http"

	"github.com/mholt/binding"
)

type WechatReq struct {
	Code          string `json:"code"`
	RawData       string `json:"rawData"`
	Signature     string `json:"signature"`
	EncryptedData string `json:"encryptedData"`
	Iv            string `json:"iv"`
}

func (wechat *WechatReq) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&wechat.Code: binding.Field{
			Form:         "code",
			Required:     true,
			ErrorMessage: "1",
		},
		&wechat.RawData: binding.Field{
			Form:         "rawData",
			Required:     true,
			ErrorMessage: "2",
		},
		&wechat.Signature: binding.Field{
			Form:         "signature",
			Required:     true,
			ErrorMessage: "3",
		},
		&wechat.EncryptedData: binding.Field{
			Form:         "encryptedData",
			Required:     true,
			ErrorMessage: "4",
		},
		&wechat.Iv: binding.Field{
			Form:         "iv",
			Required:     true,
			ErrorMessage: "5",
		},
	}
}
