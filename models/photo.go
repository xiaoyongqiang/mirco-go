package models

import (
	"database/sql"
	"errors"
	"fmt"
	"swingbaby-go/config"
	"time"
)

type PhotoStruct struct {
	Ca_id               int64
	Ca_ci_id            int64
	Ca_bu_id            int64
	Ca_name             string
	Ca_status           int64
	Ca_last_upload_time int64
	Ca_photo_num        int64
	Ca_created_time     int64
	Ap_id               int64
	Ap_ca_id            int64
	Ap_path             string
	Ap_bu_id            int64
	Ap_status           int64
	Ap_created_time     int64
}

func (t *PhotoStruct) CreateAblum(arr []string) error {

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

	rs, err := tx.Exec("INSERT INTO swing_class_albums SET ca_ci_id = ?, ca_bu_id = ?, ca_name = ?, ca_last_upload_time = ?,ca_created_time = ?,ca_photo_num = ? ",
		t.Ca_ci_id, t.Ca_bu_id, t.Ca_name, t.Ca_last_upload_time, t.Ca_created_time, t.Ca_photo_num)
	if err := check(2, rs, err, "INSERT INTO swing_class_albums fail"); err != nil {
		return err
	}

	t.Ca_id, err = rs.LastInsertId()
	if err != nil {
		return err
	}

	for _, v := range arr {
		rs, err = tx.Exec("INSERT INTO swing_album_photos SET ap_path = ?, ap_created_time = ?, ap_ca_id = ?, ap_bu_id = ?",
			v, time.Now().Unix(), t.Ca_id, t.Ca_bu_id)
		if err := check(2, rs, err, "INSERT INTO swing_album_photos fail"); err != nil {
			return err
		}
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

func (t *PhotoStruct) AddPhoto(arr []string) error {

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

	for _, v := range arr {
		rs, err := tx.Exec("INSERT INTO swing_album_photos SET ap_path = ?, ap_created_time = ?, ap_ca_id = ?, ap_bu_id = ?",
			v, time.Now().Unix(), t.Ca_id, t.Ca_bu_id)
		if err := check(2, rs, err, "INSERT INTO swing_album_photos fail"); err != nil {
			return err
		}
	}

	info, err := GetAlbumInfo(t.Ca_id)
	if err != nil {
		return err
	}

	rs, err := tx.Exec("UPDATE swing_class_albums SET ca_last_upload_time = ?, ca_photo_num = ? WHERE ca_id = ? AND ca_status = 1 ", time.Now().Unix(), info.Ca_photo_num+int64(len(arr)), t.Ca_id)
	if err := check(1, rs, err, "UPDATE swing_class_albums fail"); err != nil {
		return err
	}

	t.Ca_id, err = rs.LastInsertId()
	if err != nil {
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

func GetAlbumInfo(caId int64) (*PhotoStruct, error) {
	temp := new(PhotoStruct)
	err := config.DBHandle.QueryRow(" SELECT ca_ci_id, ca_bu_id, ca_name, ca_photo_num FROM swing_class_albums WHERE ca_id = ? AND ca_status = 1 LIMIT 1 ", caId).
		Scan(&temp.Ca_ci_id, &temp.Ca_bu_id, &temp.Ca_name, &temp.Ca_photo_num)

	if err != nil {
		return nil, err
	}

	return temp, nil
}
