package ctrls

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"swingbaby-go/apis/params"
	"swingbaby-go/config"
	"swingbaby-go/models"
)

func ControlBases(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	classRep := new(params.ClassRep)
	sessId := r.Header.Get("sessid")
	arr := strings.Split(r.URL.Path, "/")

	if arr[1] != "" {

		if arr[1] != "wechat" && arr[1] != "pay" && arr[1] != "images" {
			if sessId == "" {
				b, _ := classRep.Result(config.UNLOGIN, "Please Login")
				fmt.Fprintf(rw, "%s", b)
				return
			}

			userId, _ := config.RedisHandle.Get(sessId).Result()
			if userId == "" {
				b, _ := classRep.Result(config.UNLOGIN, "Please Login")
				fmt.Fprintf(rw, "%s", b)
				return
			}

			config.Uid, _ = strconv.ParseInt(userId, 10, 64)
			userStruct := &models.UserStruct{
				Bu_id: config.Uid,
			}

			err := userStruct.GetUserByUId()
			if err != nil {
				b, _ := classRep.Result(config.FAIL, config.Errflag[-5])
				fmt.Fprintf(rw, "%s", b)
				return
			}
		}
	}

	next(rw, r)
}
