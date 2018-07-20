package params

import (
	"encoding/json"
)

type PhotoRep struct {
	RetCode string ` json:"code,omitempty"`
	RetMsg  string `json:"msg,omitempty"`
}

type PhotoRep01 struct {
	RetCode string                   ` json:"code,omitempty"`
	RetData []map[string]interface{} `json:"data"`
}

type PhotoRep02 struct {
	RetCode string                 ` json:"code,omitempty"`
	RetData map[string]interface{} `json:"data"`
}

func (photo *PhotoRep) Result(code, msg string) ([]byte, error) {
	photo.RetCode = code
	photo.RetMsg = msg

	return json.Marshal(photo)
}

func (photo *PhotoRep01) SuccessResult01(code string, data []map[string]interface{}) ([]byte, error) {
	photo.RetCode = code
	photo.RetData = data

	return json.Marshal(photo)
}

func (photo *PhotoRep02) SuccessResult02(code string, data map[string]interface{}) ([]byte, error) {
	photo.RetCode = code
	photo.RetData = data

	return json.Marshal(photo)
}
