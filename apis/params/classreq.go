package params

import (
	"net/http"

	"github.com/mholt/binding"
)

type ClassReq struct {
	Ci_name      string `json:"cls_name"`
	Ci_telephone string `json:"cls_tel"`
}

type ClassEditmobilReq struct {
	Ci_id        string `json:"cls_id"`
	Ci_telephone string `json:"cls_mobil"`
}

type ClassEditnameReq struct {
	Ci_id   string `json:"cls_id"`
	Ci_name string `json:"cls_name"`
}

type ClassTremoveReq struct {
	Ci_id string `json:"cls_id"`
	Bb_id string `json:"b_id"`
}

type ClassJoinReq struct {
	Ci_id string `json:"cls_id"`
	Bb_id string `json:"b_id"`
}

type ClassPremoveReq struct {
	Ci_id string `json:"cls_id"`
	Bb_id string `json:"b_id"`
}

type ClassMlistReq struct {
	Bb_id string `json:"b_id"`
}

type ClassApplyReq struct {
	Ci_id   string `json:"cls_id"`
	Ci_join string `json:"cls_join"`
}

func (class *ClassReq) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&class.Ci_name: binding.Field{
			Form:         "cls_name",
			Required:     true,
			ErrorMessage: "1",
		},
		&class.Ci_telephone: binding.Field{
			Form:         "cls_tel",
			Required:     true,
			ErrorMessage: "2",
		},
	}
}

func (class *ClassEditmobilReq) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&class.Ci_id: binding.Field{
			Form:         "cls_id",
			Required:     true,
			ErrorMessage: "1",
		},
		&class.Ci_telephone: binding.Field{
			Form:         "cls_mobil",
			Required:     true,
			ErrorMessage: "2",
		},
	}
}

func (class *ClassEditnameReq) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&class.Ci_id: binding.Field{
			Form:         "cls_id",
			Required:     true,
			ErrorMessage: "1",
		},
		&class.Ci_name: binding.Field{
			Form:         "cls_name",
			Required:     true,
			ErrorMessage: "2",
		},
	}
}

func (class *ClassTremoveReq) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&class.Ci_id: binding.Field{
			Form:         "cls_id",
			Required:     true,
			ErrorMessage: "1",
		},
		&class.Bb_id: binding.Field{
			Form:         "b_id",
			Required:     true,
			ErrorMessage: "2",
		},
	}
}

func (class *ClassJoinReq) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&class.Ci_id: binding.Field{
			Form:         "cls_id",
			Required:     true,
			ErrorMessage: "1",
		},
		&class.Bb_id: binding.Field{
			Form:         "b_id",
			Required:     true,
			ErrorMessage: "2",
		},
	}
}

func (class *ClassPremoveReq) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&class.Ci_id: binding.Field{
			Form:         "cls_id",
			Required:     true,
			ErrorMessage: "1",
		},
		&class.Bb_id: binding.Field{
			Form:         "b_id",
			Required:     true,
			ErrorMessage: "2",
		},
	}
}

func (class *ClassMlistReq) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&class.Bb_id: binding.Field{
			Form:         "b_id",
			Required:     true,
			ErrorMessage: "1",
		},
	}
}

func (class *ClassApplyReq) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&class.Ci_id: binding.Field{
			Form:         "cls_id",
			Required:     true,
			ErrorMessage: "1",
		},
		&class.Ci_join: binding.Field{
			Form:         "cls_join",
			Required:     true,
			ErrorMessage: "2",
		},
	}
}
