package models

import "swingbaby-go/config"

type BabyStruct struct {
	Bb_name         string
	Bb_head_img     string
	Bb_id           int64
	Bb_sex          int32
	Bb_bu_id        int64
	Bb_created_time int64
	Fr_id           int64
	Fr_bu_id        int64
	Fr_bb_id        int64
	Fr_rel          int64
	Fr_is_master    int64
	Fr_created_time int64
	Fr_role_name    string
}

func GetInfoByBId(babyId int64) (*BabyStruct, error) {
	temp := new(BabyStruct)
	err := config.DBHandle.QueryRow(" SELECT bb_id, bb_name, bb_sex, bb_head_img FROM swing_basic_babies WHERE bb_id = ? LIMIT 1 ", babyId).
		Scan(&temp.Bb_id, &temp.Bb_name, &temp.Bb_sex, &temp.Bb_head_img)

	if err != nil {
		return nil, err
	}

	return temp, nil
}

func GetRoleByBid(uId int64, bId int64) (*BabyStruct, error) {
	temp := new(BabyStruct)
	err := config.DBHandle.QueryRow(" SELECT fr_is_master, fr_rel, fr_role_name FROM swing_family_rels WHERE fr_bu_id = ? AND fr_bb_id = ? LIMIT 1 ", uId, bId).
		Scan(&temp.Fr_is_master, &temp.Fr_rel, &temp.Fr_role_name)

	if err != nil {
		return nil, err
	}

	return temp, nil
}
