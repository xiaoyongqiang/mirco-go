package logics

import (
	"swingbaby-go/models"
)

func CreateAblum(data *models.PhotoStruct, arr []string) ([]map[string]interface{}, error) {

	var res []map[string]interface{}
	var ret []map[string]interface{}
	//res, _ = models.GetListByTeacher(uId)

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
