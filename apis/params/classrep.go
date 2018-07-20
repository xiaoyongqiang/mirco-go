package params

import (
	"encoding/json"
)

type ClassRep struct {
	RetCode string ` json:"code,omitempty"`
	RetMsg  string `json:"msg,omitempty"`
}

type ClassRep01 struct {
	RetCode string                   ` json:"code,omitempty"`
	RetData []map[string]interface{} `json:"data"`
}

type ClassRep02 struct {
	RetCode string                 ` json:"code,omitempty"`
	RetData map[string]interface{} `json:"data"`
}

func (class *ClassRep) Result(code, msg string) ([]byte, error) {
	class.RetCode = code
	class.RetMsg = msg

	return json.Marshal(class)
}

func (class *ClassRep01) SuccessResult01(code string, data []map[string]interface{}) ([]byte, error) {
	class.RetCode = code
	class.RetData = data

	return json.Marshal(class)
}

func (class *ClassRep02) SuccessResult02(code string, data map[string]interface{}) ([]byte, error) {
	class.RetCode = code
	class.RetData = data

	return json.Marshal(class)
}
