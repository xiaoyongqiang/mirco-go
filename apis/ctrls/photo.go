package ctrls

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"net/http"

	"swingbaby-go/apis/logics"
	"swingbaby-go/apis/params"
	"swingbaby-go/config"
	"swingbaby-go/models"
	"swingbaby-go/utils"

	"github.com/julienschmidt/httprouter"
	"github.com/mholt/binding"
)

//NewAlbum 创建相册
func NewAlbum(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	photoReq := new(params.AlbumReq)
	photoRep := new(params.PhotoRep)

	uID := config.Uid
	if err := binding.Bind(r, photoReq); err != nil {
		log.Printf("%v\n", err)
		b, _ := photoRep.Result(config.FAIL, config.MISSPARAMATER)
		fmt.Fprintf(rw, "%s", b)
		return
	}

	arr := strings.Split(photoReq.Ci_photos, ",")
	for _, v := range arr {
		if v == "" {
			b, _ := photoRep.Result(config.FAIL, "图片参数不能为空")
			fmt.Fprintf(rw, "%s", b)
			return
		}
	}
	if len(arr) > 9 {
		b, _ := photoRep.Result(config.FAIL, "图片上传不能超过9张")
		fmt.Fprintf(rw, "%s", b)
		return
	}

	cID, _ := strconv.ParseInt(photoReq.Ci_id, 10, 64)
	bID, _ := strconv.ParseInt(photoReq.Bb_id, 10, 64)
	if cID <= 0 {
		b, _ := photoRep.Result(config.FAIL, config.MISSPARAMATER)
		fmt.Fprintf(rw, "%s", b)
		return
	}

	if photoReq.Bb_id != "" {
		if bID <= 0 {
			b, _ := photoRep.Result(config.FAIL, config.MISSPARAMATER)
			fmt.Fprintf(rw, "%s", b)
			return
		}

		status := logics.ParentAndBabyAndClass(uID, cID, bID)
		if status <= 0 {
			b, _ := photoRep.Result(config.FAIL, config.Errflag[status])
			fmt.Fprintf(rw, "%s", b)
			return
		}
	} else {
		status := logics.TeacherOnClass(uID, cID)
		if status <= 0 {
			b, _ := photoRep.Result(config.FAIL, config.Errflag[status])
			fmt.Fprintf(rw, "%s", b)
			return
		}
	}

	photoStruct := &models.PhotoStruct{
		Ca_ci_id:            cID,
		Ca_bu_id:            uID,
		Ca_name:             photoReq.Ci_title,
		Ca_last_upload_time: time.Now().Unix(),
		Ca_created_time:     time.Now().Unix(),
		Ca_photo_num:        int64(len(arr)),
	}

	err := photoStruct.CreateAblum(arr)
	if err != nil {
		b, _ := photoRep.Result(config.FAIL, "添加相册失败")
		fmt.Fprintf(rw, "%s", b)
		log.Printf("photoStruct.CreateAblum() fail. Error:%s\n", err.Error())
		return
	}

	b, _ := photoRep.Result(config.SUCCESS, "操作成功")
	fmt.Fprintf(rw, "%s", b)
}

//NewPhoto 创建照片
func NewPhoto(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	photoReq := new(params.PhotoReq)
	photoRep := new(params.PhotoRep)

	uID := config.Uid
	if err := binding.Bind(r, photoReq); err != nil {
		log.Printf("%v\n", err)
		b, _ := photoRep.Result(config.FAIL, config.MISSPARAMATER)
		fmt.Fprintf(rw, "%s", b)
		return
	}

	arr := strings.Split(photoReq.Ci_photos, ",")
	for _, v := range arr {
		if v == "" {
			b, _ := photoRep.Result(config.FAIL, "图片参数不能为空")
			fmt.Fprintf(rw, "%s", b)
			return
		}
	}
	if len(arr) > 9 {
		b, _ := photoRep.Result(config.FAIL, "图片上传不能超过9张")
		fmt.Fprintf(rw, "%s", b)
		return
	}

	cID, _ := strconv.ParseInt(photoReq.Ci_id, 10, 64)
	bID, _ := strconv.ParseInt(photoReq.Bb_id, 10, 64)
	caID, _ := strconv.ParseInt(photoReq.Ca_id, 10, 64)
	if cID <= 0 {
		b, _ := photoRep.Result(config.FAIL, config.MISSPARAMATER)
		fmt.Fprintf(rw, "%s", b)
		return
	}

	if photoReq.Bb_id != "" {
		if bID <= 0 {
			b, _ := photoRep.Result(config.FAIL, config.MISSPARAMATER)
			fmt.Fprintf(rw, "%s", b)
			return
		}

		status := logics.ParentAndBabyAndClass(uID, cID, bID)
		if status <= 0 {
			b, _ := photoRep.Result(config.FAIL, config.Errflag[status])
			fmt.Fprintf(rw, "%s", b)
			return
		}
	} else {
		status := logics.TeacherOnClass(uID, cID)
		if status <= 0 {
			b, _ := photoRep.Result(config.FAIL, config.Errflag[status])
			fmt.Fprintf(rw, "%s", b)
			return
		}
	}

	photoStruct := &models.PhotoStruct{
		Ca_ci_id: cID,
		Ca_bu_id: uID,
		Ca_id:    caID,
	}

	err := photoStruct.AddPhoto(arr)
	if err != nil {
		b, _ := photoRep.Result(config.FAIL, "添加照片失败")
		fmt.Fprintf(rw, "%s", b)
		log.Printf("photoStruct.AddPhoto() fail. Error:%s\n", err.Error())
		return
	}

	b, _ := photoRep.Result(config.SUCCESS, "操作成功")
	fmt.Fprintf(rw, "%s", b)
}

//PhotoAddImg 上传多张图片
func PhotoAddImg(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	photoReq := new(params.PhotoImgReq)
	photoRep := new(params.PhotoRep)

	if err := binding.Bind(r, photoReq); err != nil {
		log.Printf("%v\n", err)
		b, _ := photoRep.Result(config.FAIL, config.MISSPARAMATER)
		fmt.Fprintf(rw, "%s", b)
		return
	}

	if photoReq.Typ != "1" && photoReq.Typ != "3" {
		b, _ := photoRep.Result(config.FAIL, config.MISSPARAMATER)
		fmt.Fprintf(rw, "%s", b)
		return
	}

	var path string = config.ImgPath["photo"]
	switch photoReq.Typ {
	case "1":
		path = config.ImgPath["photo"]
	case "3":
		path = config.ImgPath["notice"]
	}

	path, err := utils.Upload(path, "photo", r)
	if err != nil {
		b, _ := photoRep.Result(config.FAIL, "上传照片失败")
		fmt.Fprintf(rw, "%s", b)
		log.Printf("utils.Upload() fail. Error:%s\n", err.Error())
		return
	}

	photoRep02 := new(params.PhotoRep02)
	pathdata := map[string]interface{}{"path": path}
	b, _ := photoRep02.SuccessResult02(config.SUCCESS, pathdata)
	fmt.Fprintf(rw, "%s", b)
}
