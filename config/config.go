package config

import (
	"database/sql"
	"errors"

	"github.com/chanxuehong/session"
	"github.com/go-redis/redis"
)

var (
	RedisHandle    *redis.Client
	DBHandle       *sql.DB
	Config         *CmdConfig
	SessionStorage = session.New(20*60, 60*60)
	Uid            int64
	CheckLogin     = false
)

const (
	APPID         = "wx727aef7270281d7b"
	SECRET        = "63dcbdce38d37589706e5e72602a5ae1"
	FAIL          = "500"
	SUCCESS       = "200"
	UNLOGIN       = "401"
	MISSPARAMATER = "Invalid Data"
	MOBILEFAIL    = "Telephone Invalid Data"
	LOGINOUTTIME  = 3600
)

const (
	WxAppId           = "wx9fa2bca7d1153b7e"
	WxAppSecret       = "7b46671e63c28dfe94891cc261070f6b"
	Oauth2RedirectURI = "http://api.beibei1.butup.me/wechat/oauthcb"
	Oauth2Scope       = "snsapi_userinfo"
)

var Errflag = map[int64]string{
	-1: "该班级不存在",
	-2: "您不是该宝宝的家长",
	-3: "该宝宝不存在该班级",
	-4: "您不是该班级的老师",
	-5: "用户信息不存在",
}

var ImgPath = map[string]string{
	"baby":   "./images/baby/",
	"photo":  "./images/photo/",
	"notice": "./images/notice/",
	"static": "./images/static/",
}

// InitializeConn to initialize db and redis
func InitializeConn() (err error) {
	DBHandle, err = Config.DbConnection()
	if err != nil {
		return errors.New("Config.DbConnection() failed. Error info: " + err.Error())
	}

	RedisHandle = Config.RedisConnection()
	return nil
}
