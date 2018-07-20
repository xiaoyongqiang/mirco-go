package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func Md5(params interface{}, md5key string, exclude map[string]string) (string, error) {
	s, err := JoinStr(params, exclude)
	if err != nil {
		return "", err
	}

	hasher := md5.New()
	if _, ok := exclude["___key___"]; ok == true {
		s += md5key
		hasher.Write([]byte(s))
		return hex.EncodeToString(hasher.Sum(nil)), nil
	}

	s += "&key=" + md5key
	hasher.Write([]byte(s))
	return strings.ToUpper(hex.EncodeToString(hasher.Sum(nil))), nil
}

func VerifyMd5(params interface{}, md5key string, exclude map[string]string) (bool, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return false, err
	}

	r := make(map[string]string)
	err = json.Unmarshal(bytes, &r)
	if err != nil {
		return false, err
	}

	vSign, err := Md5(params, md5key, exclude)
	if err != nil {
		return false, err
	}

	var k = "sign"
	if _, ok := exclude["___key___"]; ok == true {
		k = "signature"
	}

	if r[k] == vSign {
		return true, nil
	}

	return false, nil
}

func GetMd5(content string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(content))
	return fmt.Sprintf("%x", md5Ctx.Sum(nil))
}

func GetHash256(content string) string {
	h := sha256.New()
	h.Write([]byte(content))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GetHash512(content string) string {
	h := sha512.New()
	h.Write([]byte(content))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func JoinStr(params interface{}, exclude map[string]string) (string, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

	r := make(map[string]string)
	err = json.Unmarshal(bytes, &r)
	if err != nil {
		return "", err
	}

	keys := make([]string, 0, 0)
	for k, v := range r {
		if exclude != nil {
			if _, ok := exclude[k]; ok == true {
				continue
			}
		}

		if len(v) > 0 {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	pList := make([]string, 0, 0)
	for _, key := range keys {
		pList = append(pList, key+"="+r[key])
	}

	return strings.Join(pList, "&"), nil
}
