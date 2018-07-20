package ctrls

import (
	"fmt"
	"log"
	"strconv"

	"net/http"
	"swingbaby-go/apis/logics"

	"swingbaby-go/apis/params"
	"swingbaby-go/config"
	"swingbaby-go/models"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/mholt/binding"
)

//NewClass 创建班级
func NewClass(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	classReq := new(params.ClassReq)
	classRep := new(params.ClassRep)

	uID := config.Uid
	if err := binding.Bind(r, classReq); err != nil {
		log.Printf("%v\n", err)
		b, _ := classRep.Result(config.FAIL, config.MISSPARAMATER)
		fmt.Fprintf(rw, "%s", b)
		return
	}

	userStruct := &models.UserStruct{
		Bu_id: uID,
	}

	err := userStruct.GetUserByUId()
	if err != nil {
		b, _ := classRep.Result(config.FAIL, "用户信息不存在")
		fmt.Fprintf(rw, "%s", b)
		log.Printf("userStruct.GetUserByUId() fail. Error:%s\n", err.Error())
		return
	}

	classStruct := &models.ClassStruct{
		Ci_name:         classReq.Ci_name,
		Ci_telephone:    classReq.Ci_telephone,
		Ci_bu_id:        uID,
		Ci_created_time: time.Now().Unix(),
	}

	err = classStruct.CreateClass()

	if err != nil {
		b, _ := classRep.Result(config.FAIL, "添加班级失败")
		fmt.Fprintf(rw, "%s", b)
		log.Printf("userStruct.CreateClass() fail. Error:%s\n", err.Error())
		return
	}

	config.RedisHandle.Set(fmt.Sprintf("defaultClass%d", uID), classStruct.Ci_id, 0)
	config.RedisHandle.Del(fmt.Sprintf("defaultBaby%d", uID))

	classRep02 := new(params.ClassRep02)
	b, _ := classRep02.SuccessResult02(config.SUCCESS, map[string]interface{}{
		"ci_id": classStruct.Ci_id,
	})
	fmt.Fprintf(rw, "%s", b)
}

//ClassEditmobil 修改班级电话
func ClassEditmobil(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	classReq := new(params.ClassEditmobilReq)
	classRep := new(params.ClassRep)

	uID := config.Uid
	if err := binding.Bind(r, classReq); err != nil {
		log.Printf("%v\n", err)
		b, _ := classRep.Result(config.FAIL, config.MISSPARAMATER)
		fmt.Fprintf(rw, "%s", b)
		return
	}

	clsID, _ := strconv.ParseInt(classReq.Ci_id, 10, 64)
	role := logics.TeacherOnClass(uID, clsID)

	if role <= 0 {
		b, _ := classRep.Result(config.FAIL, "不是该班级的老师")
		fmt.Fprintf(rw, "%s", b)
		return
	}

	classStruct := &models.ClassStruct{
		Ci_id:        clsID,
		Ci_telephone: classReq.Ci_telephone,
	}

	err := classStruct.EditMobil()

	if err != nil {
		b, _ := classRep.Result(config.FAIL, "修改电话失败")
		fmt.Fprintf(rw, "%s", b)
		log.Printf("classStruct.EditMobil() fail. Error:%s\n", err.Error())
		return
	}

	b, _ := classRep.Result(config.SUCCESS, "操作成功")
	fmt.Fprintf(rw, "%s", b)
}

//ClassEditname 修改班级名字
func ClassEditname(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	classReq := new(params.ClassEditnameReq)
	classRep := new(params.ClassRep)

	uID := config.Uid
	if err := binding.Bind(r, classReq); err != nil {
		log.Printf("%v\n", err)
		b, _ := classRep.Result(config.FAIL, config.MISSPARAMATER)
		fmt.Fprintf(rw, "%s", b)
		return
	}

	clsID, _ := strconv.ParseInt(classReq.Ci_id, 10, 64)
	role := logics.TeacherOnClass(uID, clsID)

	if role <= 0 {
		b, _ := classRep.Result(config.FAIL, "不是该班级的老师")
		fmt.Fprintf(rw, "%s", b)
		return
	}

	classStruct := &models.ClassStruct{
		Ci_id:   clsID,
		Ci_name: classReq.Ci_name,
	}

	err := classStruct.EditName()

	if err != nil {
		b, _ := classRep.Result(config.FAIL, "修改班级名称失败")
		fmt.Fprintf(rw, "%s", b)
		log.Printf("classStruct.EditName() fail. Error:%s\n", err.Error())
		return
	}

	b, _ := classRep.Result(config.SUCCESS, "操作成功")
	fmt.Fprintf(rw, "%s", b)
}

// ClassTremove 老师将宝宝移除班级
func ClassTremove(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	classReq := new(params.ClassTremoveReq)
	classRep := new(params.ClassRep)

	uID := config.Uid
	if err := binding.Bind(r, classReq); err != nil {
		log.Printf("%v\n", err)
		b, _ := classRep.Result(config.FAIL, config.MISSPARAMATER)
		fmt.Fprintf(rw, "%s", b)
		return
	}

	clsID, _ := strconv.ParseInt(classReq.Ci_id, 10, 64)
	bID, _ := strconv.ParseInt(classReq.Bb_id, 10, 64)

	role := logics.TeacherOnClass(uID, clsID)
	if role <= 0 {
		b, _ := classRep.Result(config.FAIL, "不是该班级的老师")
		fmt.Fprintf(rw, "%s", b)
		return
	}

	exist := logics.BabyOnClass(clsID, bID)
	if exist != 1 {
		b, _ := classRep.Result(config.FAIL, "该宝宝不存在该班级")
		fmt.Fprintf(rw, "%s", b)
		return
	}

	classStruct := &models.ClassStruct{
		Ci_id:     clsID,
		Cbr_bb_id: bID,
	}

	err := classStruct.RemoveClass()

	if err != nil {
		b, _ := classRep.Result(config.FAIL, "将宝宝移除班级失败")
		fmt.Fprintf(rw, "%s", b)
		log.Printf("classStruct.RemoveClass() fail. Error:%s\n", err.Error())
		return
	}

	b, _ := classRep.Result(config.SUCCESS, "操作成功")
	fmt.Fprintf(rw, "%s", b)
}

// ClassJoin 宝宝加入班级
func ClassJoin(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	classReq := new(params.ClassJoinReq)
	classRep := new(params.ClassRep)

	uID := config.Uid
	if err := binding.Bind(r, classReq); err != nil {
		log.Printf("%v\n", err)
		b, _ := classRep.Result(config.FAIL, config.MISSPARAMATER)
		fmt.Fprintf(rw, "%s", b)
		return
	}

	clsID, _ := strconv.ParseInt(classReq.Ci_id, 10, 64)
	bID, _ := strconv.ParseInt(classReq.Bb_id, 10, 64)

	if logics.ExistsClass(clsID) != 1 {
		b, _ := classRep.Result(config.FAIL, "该班级不允许加入")
		fmt.Fprintf(rw, "%s", b)
		return
	}

	if logics.ParentOnBaby(uID, bID) <= 0 {
		b, _ := classRep.Result(config.FAIL, "不是该宝宝的家长没有权限")
		fmt.Fprintf(rw, "%s", b)
		return
	}

	classStruct := &models.ClassStruct{
		Ci_id:     clsID,
		Cbr_bb_id: bID,
	}

	exist := logics.BabyOnClass(clsID, bID)
	if exist == 1 {
		b, _ := classRep.Result(config.FAIL, "宝宝已经在该班级")
		fmt.Fprintf(rw, "%s", b)
		return
	}

	if exist == 3 {
		//将删除的宝宝重新加入班级
		err := classStruct.EditBabyStatus()
		if err != nil {
			b, _ := classRep.Result(config.FAIL, "宝宝重新加入班级失败")
			fmt.Fprintf(rw, "%s", b)
			log.Printf("classStruct.EditBabyStatus() fail. Error:%s\n", err.Error())
			return
		}

		b, _ := classRep.Result(config.SUCCESS, "操作成功")
		fmt.Fprintf(rw, "%s", b)
		return
	}

	//从没有加入班级的宝宝
	err := classStruct.JoinClass()
	if err != nil {
		b, _ := classRep.Result(config.FAIL, "宝宝加入班级失败")
		fmt.Fprintf(rw, "%s", b)
		log.Printf("classStruct.JoinClass() fail. Error:%s\n", err.Error())
		return
	}

	b, _ := classRep.Result(config.SUCCESS, "操作成功")
	fmt.Fprintf(rw, "%s", b)
}

//ClassPremove 家长将宝宝移除班级
func ClassPremove(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	classReq := new(params.ClassPremoveReq)
	classRep := new(params.ClassRep)

	uID := config.Uid
	if err := binding.Bind(r, classReq); err != nil {
		log.Printf("%v\n", err)
		b, _ := classRep.Result(config.FAIL, config.MISSPARAMATER)
		fmt.Fprintf(rw, "%s", b)
		return
	}

	clsID, _ := strconv.ParseInt(classReq.Ci_id, 10, 64)
	bID, _ := strconv.ParseInt(classReq.Bb_id, 10, 64)

	if logics.ParentOnBaby(uID, bID) <= 0 {
		b, _ := classRep.Result(config.FAIL, "不是该宝宝的家长没有权限")
		fmt.Fprintf(rw, "%s", b)
		return
	}

	exist := logics.BabyOnClass(clsID, bID)
	if exist != 1 {
		b, _ := classRep.Result(config.FAIL, "该宝宝不存在该班级")
		fmt.Fprintf(rw, "%s", b)
		return
	}

	classStruct := &models.ClassStruct{
		Ci_id:     clsID,
		Cbr_bb_id: bID,
	}

	err := classStruct.RemoveClass()

	if err != nil {
		b, _ := classRep.Result(config.FAIL, "将宝宝移除班级失败")
		fmt.Fprintf(rw, "%s", b)
		log.Printf("classStruct.RemoveClass() fail. Error:%s\n", err.Error())
		return
	}

	b, _ := classRep.Result(config.SUCCESS, "操作成功")
	fmt.Fprintf(rw, "%s", b)
}

//ClassMlist 家委获取拥有权 限班级列表
func ClassMlist(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	classReq := new(params.ClassMlistReq)
	classRep := new(params.ClassRep)

	uID := config.Uid
	if err := binding.Bind(r, classReq); err != nil {
		log.Printf("%v\n", err)
		b, _ := classRep.Result(config.FAIL, config.MISSPARAMATER)
		fmt.Fprintf(rw, "%s", b)
		return
	}

	bID, _ := strconv.ParseInt(classReq.Bb_id, 10, 64)
	if logics.ParentOnBaby(uID, bID) != 1 {
		b, _ := classRep.Result(config.FAIL, "您不是该宝宝的主家长")
		fmt.Fprintf(rw, "%s", b)
		return
	}
	res, err := logics.GetListByBid(bID, 1)

	if err != nil {
		b, _ := classRep.Result(config.FAIL, "获取班级信息失败")
		fmt.Fprintf(rw, "%s", b)
		log.Printf("classStruct.GetListByBid() fail. Error:%s\n", err.Error())
		return
	}

	classRep01 := new(params.ClassRep01)
	b, _ := classRep01.SuccessResult01(config.SUCCESS, res)
	fmt.Fprintf(rw, "%s", b)
}

//ClassPlist 家长获取宝宝所在班级列表
func ClassPlist(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	classReq := new(params.ClassMlistReq)
	classRep := new(params.ClassRep)

	uID := config.Uid
	if err := binding.Bind(r, classReq); err != nil {
		log.Printf("%v\n", err)
		b, _ := classRep.Result(config.FAIL, config.MISSPARAMATER)
		fmt.Fprintf(rw, "%s", b)
		return
	}

	bID, _ := strconv.ParseInt(classReq.Bb_id, 10, 64)
	if logics.ParentOnBaby(uID, bID) <= 0 {
		b, _ := classRep.Result(config.FAIL, "您不是该宝宝的家长")
		fmt.Fprintf(rw, "%s", b)
		return
	}
	res, err := logics.GetListByBid(bID, 3)

	if err != nil {
		b, _ := classRep.Result(config.FAIL, "获取班级信息失败")
		fmt.Fprintf(rw, "%s", b)
		log.Printf("classStruct.GetListByBid() fail. Error:%s\n", err.Error())
		return
	}

	classRep01 := new(params.ClassRep01)
	b, _ := classRep01.SuccessResult01(config.SUCCESS, res)
	fmt.Fprintf(rw, "%s", b)
}

//ClassTlist 教师获取所有班级列表
func ClassTlist(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uID := config.Uid
	res, _ := logics.GetListByTeacher(uID)
	classRep01 := new(params.ClassRep01)
	b, _ := classRep01.SuccessResult01(config.SUCCESS, res)
	fmt.Fprintf(rw, "%s", b)
}

// ClassApply 申请加入班级开关
func ClassApply(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	classReq := new(params.ClassApplyReq)
	classRep := new(params.ClassRep)

	uID := config.Uid
	if err := binding.Bind(r, classReq); err != nil {
		log.Printf("%v\n", err)
		b, _ := classRep.Result(config.FAIL, config.MISSPARAMATER)
		fmt.Fprintf(rw, "%s", b)
		return
	}

	if classReq.Ci_join != "1" && classReq.Ci_join != "3" {
		b, _ := classRep.Result(config.FAIL, config.MISSPARAMATER)
		fmt.Fprintf(rw, "%s", b)
		return
	}

	clsID, _ := strconv.ParseInt(classReq.Ci_id, 10, 64)
	clsIsOpen, _ := strconv.ParseInt(classReq.Ci_join, 10, 64)
	role := logics.TeacherOnClass(uID, clsID)

	if role <= 0 {
		b, _ := classRep.Result(config.FAIL, "不是该班级的老师")
		fmt.Fprintf(rw, "%s", b)
		return
	}

	classStruct := &models.ClassStruct{
		Ci_id:      clsID,
		Ci_is_open: clsIsOpen,
	}

	err := classStruct.ApplyClass()
	if err != nil {
		b, _ := classRep.Result(config.FAIL, "操作失败")
		fmt.Fprintf(rw, "%s", b)
		log.Printf("classStruct.ApplyClass() fail. Error:%s\n", err.Error())
		return
	}

	b, _ := classRep.Result(config.SUCCESS, "操作成功")
	fmt.Fprintf(rw, "%s", b)
}
