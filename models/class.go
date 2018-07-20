package models

import (
	"database/sql"
	"errors"
	"fmt"
	"swingbaby-go/config"
	"time"
)

type ClassStruct struct {
	Ci_name          string
	Ci_telephone     string
	Ci_id            int64
	Ci_bu_id         int64
	Ci_is_open       int64
	Ci_status        int64
	Ci_created_time  int64
	Cbr_status       int64
	Cbr_id           int64
	Cbr_bb_id        int64
	Cbr_ci_id        int64
	Cbr_created_time int64
	Cbr_is_officer   int64
	Cbr_memo         string
}

func GetInfoByCId(classId int64) (*ClassStruct, error) {
	temp := new(ClassStruct)
	err := config.DBHandle.QueryRow(" SELECT ci_telephone, ci_is_open, ci_name, ci_id, ci_bu_id FROM swing_class_infos WHERE ci_id = ? AND ci_status = 1 ", classId).
		Scan(&temp.Ci_telephone, &temp.Ci_is_open, &temp.Ci_name, &temp.Ci_id, &temp.Ci_bu_id)

	if err != nil {
		return nil, err
	}

	return temp, nil
}

func GetRoleByCid(uId int64, classId int64) (int64, error) {
	var role int64
	err := config.DBHandle.QueryRow(" SELECT cur_is_master FROM swing_class_teacher_rels WHERE cur_ci_id = ? AND cur_bu_id = ? LIMIT 1 ", classId, uId).
		Scan(&role)

	if err != nil {
		return 0, err
	}

	return role, nil
}

func (t *ClassStruct) CreateClass() error {

	tx, err := config.DBHandle.Begin()
	defer func() {
		tx.Rollback()
	}()

	if err != nil {
		return err
	}

	check := func(typ int, rs sql.Result, err error, msg string) error {
		if err != nil {
			return fmt.Errorf("%s. Error:%s", msg, err.Error())
		}

		var rai int64
		if typ == 1 {
			rai, err = rs.RowsAffected()
		} else {
			rai, err = rs.LastInsertId()
		}

		if err != nil {
			return fmt.Errorf("%s. Error:%s", msg, err.Error())
		}

		if rai <= 0 {
			return errors.New(msg)
		}

		return nil
	}

	rs, err := tx.Exec("INSERT INTO swing_class_infos SET ci_name = ?, ci_telephone = ?, ci_bu_id = ?, ci_created_time = ?",
		t.Ci_name, t.Ci_telephone, t.Ci_bu_id, t.Ci_created_time)
	if err := check(2, rs, err, "INSERT INTO swing_class_infos fail"); err != nil {
		return err
	}

	t.Ci_id, err = rs.LastInsertId()
	if err != nil {
		return err
	}

	rs, err = tx.Exec("INSERT INTO swing_class_teacher_rels SET cur_ci_id = ?, cur_bu_id = ?, cur_is_master = ?, cur_created_time = ?",
		t.Ci_id, t.Ci_bu_id, 1, t.Ci_created_time)
	if err := check(2, rs, err, "INSERT INTO swing_class_teacher_rels fail"); err != nil {
		return err
	}

	rs, err = tx.Exec("INSERT INTO swing_class_finance SET cf_ci_id = ?", t.Ci_id)
	if err := check(2, rs, err, "INSERT INTO swing_class_finance fail"); err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}

		return err
	}

	return nil
}

func (t *ClassStruct) EditMobil() error {
	rs, err := config.DBHandle.Exec("UPDATE swing_class_infos SET ci_telephone = ? WHERE ci_id = ?", t.Ci_telephone, t.Ci_id)

	if err != nil {
		return err
	}

	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (t *ClassStruct) EditName() error {
	rs, err := config.DBHandle.Exec("UPDATE swing_class_infos SET ci_name = ? WHERE ci_id = ?", t.Ci_name, t.Ci_id)

	if err != nil {
		return err
	}

	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func BabyOnClass(clsId int64, bId int64) (*ClassStruct, error) {
	temp := new(ClassStruct)
	err := config.DBHandle.QueryRow("SELECT cbr_id, cbr_status, cbr_is_officer, cbr_memo FROM swing_class_baby_rels WHERE cbr_ci_id = ? AND cbr_bb_id = ?", clsId, bId).
		Scan(&temp.Cbr_id, &temp.Cbr_status, &temp.Cbr_is_officer, &temp.Cbr_memo)

	if err != nil {
		return nil, err
	}

	return temp, nil
}

func (t *ClassStruct) RemoveClass() error {
	rs, err := config.DBHandle.Exec("UPDATE swing_class_baby_rels SET cbr_status = 3 WHERE cbr_bb_id = ? AND cbr_ci_id = ? AND cbr_status = 1 ", t.Cbr_bb_id, t.Ci_id)

	if err != nil {
		return err
	}

	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (t *ClassStruct) EditBabyStatus() error {
	rs, err := config.DBHandle.Exec("UPDATE swing_class_baby_rels SET cbr_status = 1 WHERE cbr_bb_id = ? AND cbr_ci_id = ? AND cbr_status = 3 ", t.Cbr_bb_id, t.Ci_id)

	if err != nil {
		return err
	}

	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (t *ClassStruct) JoinClass() error {
	rs, err := config.DBHandle.Exec(" INSERT INTO swing_class_baby_rels SET cbr_bb_id = ?, cbr_ci_id = ?, cbr_status = ?, cbr_created_time = ? ", t.Cbr_bb_id, t.Ci_id, 1, time.Now().Unix())

	if err != nil {
		return err
	}

	t.Cbr_id, err = rs.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func GetListByManager(bId int64) ([]map[string]interface{}, error) {
	rows, err := config.DBHandle.Query(" SELECT cbr_ci_id FROM swing_class_baby_rels WHERE cbr_bb_id = ? AND cbr_status= 1 AND cbr_is_officer = 3 ORDER BY cbr_id DESC", bId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var (
		id  int64
		ret []map[string]interface{}
	)

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		ret = append(ret, map[string]interface{}{
			"id": id,
		})
	}

	return ret, nil
}

func GetListByTeacher(uId int64) ([]map[string]interface{}, error) {
	rows, err := config.DBHandle.Query(" SELECT cur_ci_id FROM swing_class_teacher_rels WHERE cur_bu_id = ? ORDER BY cur_id DESC ", uId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var (
		id  int64
		ret []map[string]interface{}
	)

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		ret = append(ret, map[string]interface{}{
			"id": id,
		})
	}

	return ret, nil
}

func GetListByParent(bId int64) ([]map[string]interface{}, error) {
	rows, err := config.DBHandle.Query(" SELECT cbr_ci_id, cbr_is_officer FROM swing_class_baby_rels WHERE cbr_bb_id = ? AND cbr_status= 1 ORDER BY cbr_id DESC ", bId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var (
		id      int64
		officer int64
		ret     []map[string]interface{}
	)

	for rows.Next() {
		err = rows.Scan(&id, &officer)
		if err != nil {
			return nil, err
		}

		ret = append(ret, map[string]interface{}{
			"id":      id,
			"officer": officer,
		})
	}

	return ret, nil
}

func (t *ClassStruct) ApplyClass() error {

	var old int64
	if t.Ci_is_open == 1 {
		old = 3
	} else {
		old = 1
	}

	rs, err := config.DBHandle.Exec("UPDATE swing_class_infos SET ci_is_open = ? WHERE ci_id = ? AND ci_is_open = ? AND ci_status = 1 ", t.Ci_is_open, t.Ci_id, old)

	if err != nil {
		return err
	}

	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
