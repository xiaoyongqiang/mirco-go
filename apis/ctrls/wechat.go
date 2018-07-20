package ctrls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"swingbaby-go/apis/params"
	"swingbaby-go/config"
	"swingbaby-go/models"
	"swingbaby-go/utils"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/mholt/binding"
	"github.com/xlstudio/wxbizdatacrypt"
)

//NewWechat 小程序微信认证登录
func NewWechat(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	wechatReq := new(params.WechatReq)
	wechatRep := new(params.WechatRep)

	if err := binding.Bind(r, wechatReq); err != nil {
		log.Printf("%v\n", err)
		b, _ := wechatRep.Result(config.FAIL, config.MISSPARAMATER)
		fmt.Fprintf(rw, "%s", b)
		return
	}

	str := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", config.APPID, config.SECRET, wechatReq.Code)
	resp, err := http.Get(str)
	if err != nil {
		log.Printf("%v\n", err)
		b, _ := wechatRep.Result(config.FAIL, "请求微信session_key失败")
		fmt.Fprintf(rw, "%s", b)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%v\n", err)
		b, _ := wechatRep.Result(config.FAIL, "读取微信session_key失败")
		fmt.Fprintf(rw, "%s", b)
		return
	}

	wechatVerify := new(params.WechatVerify)
	err = json.Unmarshal([]byte(string(body)), wechatVerify)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	log.Print(string(body))
	if wechatVerify.Errmsg != "" {
		b, _ := wechatRep.Result(config.FAIL, "微信身份认证code错误")
		fmt.Fprintf(rw, "%s", b)
		return
	}

	//验证session_key合法性
	pc := wxbizdatacrypt.WxBizDataCrypt{AppID: config.APPID, SessionKey: wechatVerify.Session_key}
	res, err := pc.Decrypt(wechatReq.EncryptedData, wechatReq.Iv, true) //JSON: true, map: false
	if err != nil {
		log.Printf("%v\n", err)
		b, _ := wechatRep.Result(config.FAIL, "获取用户信息失败")
		fmt.Fprintf(rw, "%s", b)
		return
	}

	wechatUserInfo := new(params.WechatUserInfo)
	err = json.Unmarshal([]byte(res.(string)), wechatUserInfo)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}

	//判断用户是否已经存在
	uInfo := models.GetUserByUnionId(wechatUserInfo.UnionId)

	if uInfo == nil {
		uInfo = &models.UserStruct{
			Bu_name:         wechatUserInfo.NickName,
			Bu_head_img:     wechatUserInfo.AvatarUrl,
			Bu_open_id:      wechatUserInfo.OpenId,
			Bu_union_id:     wechatUserInfo.UnionId,
			Bu_created_time: time.Now().Unix(),
		}

		uInfo.Bu_id = uInfo.CreateUser()
		if uInfo.Bu_id <= 0 {
			log.Printf("uInfo.CreateUser() %v\n", err)
			b, _ := wechatRep.Result(config.FAIL, "添加用户失败")
			fmt.Fprintf(rw, "%s", b)
			return
		}
	}

	//log.Println(uInfo.Bu_id, uInfo.Bu_union_id)
	sessID := utils.GetRandomString(16)
	err = config.RedisHandle.Set(fmt.Sprintf("%s", sessID), uInfo.Bu_id, config.LOGINOUTTIME*time.Second).Err()
	if err != nil {
		log.Printf("%v\n", err)
		return
	}

	status := 1
	if len(uInfo.Bu_mobile) > 1 {
		status = 3
	}

	babyid, _ := config.RedisHandle.Get(fmt.Sprintf("defaultBaby%d", uInfo.Bu_id)).Result()
	baby := make(map[string]interface{})
	if babyid != "" {
		bInfo := &models.BabyStruct{}
		bID, _ := strconv.ParseInt(babyid, 10, 64)
		bInfo, _ = models.GetInfoByBId(bID)
		baby["baby_id"] = bInfo.Bb_id
		baby["baby_name"] = bInfo.Bb_name
	}

	classid, _ := config.RedisHandle.Get(fmt.Sprintf("defaultClass%d", uInfo.Bu_id)).Result()
	class := make(map[string]interface{})
	if classid != "" {
		cInfo := &models.ClassStruct{}
		cID, _ := strconv.ParseInt(classid, 10, 64)
		cInfo, _ = models.GetInfoByCId(cID)
		class["class_id"] = cInfo.Ci_id
		class["class_name"] = cInfo.Ci_name
	}

	wechatInfo := new(params.WechatInfo)
	wechatInfo.Status = status
	wechatInfo.SessionId = sessID
	wechatInfo.Baby = baby
	wechatInfo.Class = class
	b, _ := wechatRep.SuccessResult(config.SUCCESS, wechatInfo)
	fmt.Fprintf(rw, "%s", b)
}
