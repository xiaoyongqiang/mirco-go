package logics

import (
	"log"
	"swingbaby-go/models"
)

func TeacherOnClass(uId int64, clsId int64) int64 {
	_, err := models.GetInfoByCId(clsId)
	if err != nil {
		log.Printf("%v GetInfoByCId\n", err)
		return -1
	}

	role, err := models.GetRoleByCid(uId, clsId)
	if err != nil {
		log.Printf("%v GetRoleByCid\n", err)
		return -4
	}

	return role
}

func BabyOnClass(clsId int64, bId int64) int64 {
	status, err := models.BabyOnClass(clsId, bId)
	if err != nil {
		log.Printf("%v BabyOnClass\n", err)
		return -3
	}

	return status.Cbr_status
}

func ExistsClass(clsId int64) int64 {
	info, err := models.GetInfoByCId(clsId)
	if err != nil {
		log.Printf("%v GetInfoByCId\n", err)
		return -1
	}

	return info.Ci_is_open
}

func ParentOnBaby(uId int64, bId int64) int64 {
	info, err := models.GetRoleByBid(uId, bId)
	if err != nil {
		log.Printf("%v GetRoleByBid\n", err)
		return -2
	}

	return info.Fr_is_master
}

func ParentAndBabyAndClass(uId int64, clsId int64, bId int64) int64 {
	status := ExistsClass(clsId)
	if status <= 0 {
		log.Printf("errorfalg %d ExistsClass\n", status)
		return -1
	}

	status = ParentOnBaby(uId, bId)
	if status <= 0 {
		log.Printf("errorfalg %d ParentOnBaby\n", status)
		return -2
	}

	status = BabyOnClass(clsId, bId)
	if status <= 0 {
		log.Printf("errorfalg %d BabyOnClass\n", status)
		return -3
	}

	return 1
}
