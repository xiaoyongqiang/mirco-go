package logics

import (
	"swingbaby-go/models"
)

func GetListByBid(bId int64, typ int64) ([]map[string]interface{}, error) {

	var res []map[string]interface{}
	var ret []map[string]interface{}
	switch typ {
	case 1:
		res, _ = models.GetListByManager(bId)

	case 3:
		res, _ = models.GetListByParent(bId)
	}

	if res != nil {
		for _, v := range res {
			info, err := models.GetInfoByCId(v["id"].(int64))
			if err != nil {
				continue
			}

			ret = append(ret, map[string]interface{}{
				"ci_name": info.Ci_name,
				"ci_id":   info.Ci_id,
			})
		}
	}

	return ret, nil
}

func GetListByTeacher(uId int64) ([]map[string]interface{}, error) {

	var res []map[string]interface{}
	var ret []map[string]interface{}
	res, _ = models.GetListByTeacher(uId)

	if res != nil {
		for _, v := range res {
			info, err := models.GetInfoByCId(v["id"].(int64))
			if err != nil {
				continue
			}

			ret = append(ret, map[string]interface{}{
				"ci_name": info.Ci_name,
				"ci_id":   info.Ci_id,
			})
		}
	}

	return ret, nil
}
