package models

import (
	"fmt"
	"swingbaby-go/config"
	//"strconv"
	//"time"
)

type UserStruct struct {
	Bu_name         string
	Bu_head_img     string
	Bu_created_time int64
	Bu_id           int64
	Bu_real_name    string
	Bu_open_id      string
	Bu_union_id     string
	Bu_mobile       string
}

func GetUserByUnionId(UnionId string) *UserStruct {

	sql := "SELECT bu_id, bu_open_id, bu_union_id, bu_mobile FROM swing_basic_users WHERE bu_union_id = '%s' LIMIT 1"
	sql = fmt.Sprintf(sql, UnionId)
	stm, _ := config.DBHandle.Prepare(sql)
	defer stm.Close()
	rows, _ := stm.Query()

	defer rows.Close()
	var temp *UserStruct = nil
	if rows.Next() {
		temp = new(UserStruct)
		rows.Scan(&temp.Bu_id, &temp.Bu_open_id, &temp.Bu_union_id, &temp.Bu_mobile)
	}
	return temp
}

func (data *UserStruct) CreateUser() int64 {
	stm, _ := config.DBHandle.Prepare("INSERT INTO swing_basic_users( bu_name, bu_head_img, bu_open_id, bu_union_id, bu_created_time) values(?,?,?,?,?)")
	res, _ := stm.Exec(data.Bu_name, data.Bu_head_img, data.Bu_open_id, data.Bu_union_id, data.Bu_created_time)
	stm.Close()
	insertId, _ := res.LastInsertId()
	return insertId
}

func (data *UserStruct) GetUserByUId() error {
	err := config.DBHandle.QueryRow("SELECT bu_open_id, bu_union_id, bu_id, bu_mobile, bu_name, bu_real_name, bu_head_img FROM swing_basic_users WHERE bu_id = ? LIMIT 1", data.Bu_id).Scan(&data.Bu_open_id, &data.Bu_union_id, &data.Bu_id, &data.Bu_mobile, &data.Bu_name, &data.Bu_real_name, &data.Bu_head_img)
	if err != nil {
		return err
	}

	return nil
}
