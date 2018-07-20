package ctrls

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"swingbaby-go/config"

	"github.com/chanxuehong/rand"
	"github.com/julienschmidt/httprouter"
	mpoauth2 "gopkg.in/chanxuehong/wechat.v2/mp/oauth2"
	"gopkg.in/chanxuehong/wechat.v2/oauth2"
)

func OauthBases(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	//判断是否已登录或者白名单模块
	arr := strings.Split(r.URL.Path, "/")
	if arr[1] == "favicon.ico" {
		return
	}

	session, err := config.SessionStorage.Get("isAuth")
	if err == nil {
		if session.(string) == "true" {
			config.CheckLogin = true
		}
	}

	if arr[1] != "" {
		if arr[1] != "wechat" && arr[1] != "pay" && arr[1] != "images" && !config.CheckLogin {

			state := string(rand.NewHex())
			AuthCodeURL := mpoauth2.AuthCodeURL(config.WxAppId, config.Oauth2RedirectURI, config.Oauth2Scope, state)
			log.Println("AuthCodeURL:", AuthCodeURL)

			http.Redirect(w, r, AuthCodeURL, http.StatusFound)
			return
		}
	}

	next(w, r)
}

//Callback 公众号授权后回调页面
func Callback(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println(r.RequestURI)

	var oauth2Endpoint oauth2.Endpoint = mpoauth2.NewEndpoint(config.WxAppId, config.WxAppSecret)
	queryValues, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		io.WriteString(w, err.Error())
		log.Println(err)
		return
	}

	code := queryValues.Get("code")
	if code == "" {
		log.Println("用户禁止授权")
		return
	}

	queryState := queryValues.Get("state")
	if queryState == "" {
		log.Println("state参数不能为空")
		return
	}

	oauth2Client := oauth2.Client{
		Endpoint: oauth2Endpoint,
	}
	token, err := oauth2Client.ExchangeToken(code)
	if err != nil {
		io.WriteString(w, err.Error())
		log.Println(err)
		return
	}

	userinfo, err := mpoauth2.GetUserInfo(token.AccessToken, token.OpenId, "", nil)
	if err != nil {
		io.WriteString(w, err.Error())
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(userinfo)
	log.Printf("userinfo: %+v\r\n", userinfo)

	if err = config.SessionStorage.Add("isAuth", "true"); err != nil {
		io.WriteString(w, err.Error())
		log.Println(err)
	}

	return
}
