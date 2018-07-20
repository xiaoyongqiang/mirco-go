package params

import "encoding/json"

// TradeRep is the struct that to response to the client
type WechatRep struct {
	RetCode string      `json:"code,omitempty"`
	RetMsg  string      `json:"msg,omitempty"`
	RetData *WechatInfo `json:"data,omitempty"`
}

type WechatVerify struct {
	Openid      string `json:"openid,omitempty"`
	Session_key string `json:"session_key,omitempty"`
	Errcode     int    `json:"errcode,omitempty"`
	Errmsg      string `json:"errmsg,omitempty"`
}

type WechatUserInfo struct {
	NickName  string `json:"nickName,omitempty"`
	UnionId   string `json:"unionId,omitempty"`
	OpenId    string `json:"openId,omitempty"`
	AvatarUrl string `json:"avatarUrl,omitempty"`
}

type WechatInfo struct {
	Baby      map[string]interface{} `json:"babyInfo"`
	Class     map[string]interface{} `json:"classInfo"`
	Status    int                    `json:"type,omitempty"`
	SessionId string                 `json:"sess_id,omitempty"`
}

func (wechat *WechatRep) SuccessResult(code string, data *WechatInfo) ([]byte, error) {
	wechat.RetCode = code
	wechat.RetData = data

	return json.Marshal(wechat)
}

func (wechat *WechatRep) Result(code, msg string) ([]byte, error) {
	wechat.RetCode = code
	wechat.RetMsg = msg

	return json.Marshal(wechat)
}
