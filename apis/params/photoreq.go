package params

import (
	"net/http"

	"github.com/mholt/binding"
)

type AlbumReq struct {
	Ci_id     string `json:"cls_id"`
	Bb_id     string `json:"b_id"`
	Ci_title  string `json:"title"`
	Ci_photos string `json:"photos"`
}

type PhotoReq struct {
	Ci_id     string `json:"cls_id"`
	Bb_id     string `json:"b_id"`
	Ca_id     string `json:"ca_id"`
	Ci_photos string `json:"photos"`
}

type PhotoImgReq struct {
	Typ string `json:"type"`
}

func (photo *AlbumReq) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&photo.Ci_id: binding.Field{
			Form:         "cls_id",
			Required:     true,
			ErrorMessage: "1",
		},
		&photo.Ci_title: binding.Field{
			Form:         "title",
			Required:     true,
			ErrorMessage: "2",
		},
		&photo.Ci_photos: binding.Field{
			Form:         "photos",
			Required:     true,
			ErrorMessage: "3",
		},
		&photo.Bb_id: binding.Field{
			Form:         "b_id",
			Required:     false,
			ErrorMessage: "4",
		},
	}
}

func (photo *PhotoReq) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&photo.Ci_id: binding.Field{
			Form:         "cls_id",
			Required:     true,
			ErrorMessage: "1",
		},
		&photo.Ca_id: binding.Field{
			Form:         "ca_id",
			Required:     true,
			ErrorMessage: "2",
		},
		&photo.Ci_photos: binding.Field{
			Form:         "photos",
			Required:     true,
			ErrorMessage: "3",
		},
		&photo.Bb_id: binding.Field{
			Form:         "b_id",
			Required:     false,
			ErrorMessage: "4",
		},
	}
}

func (photo *PhotoImgReq) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&photo.Typ: binding.Field{
			Form:         "type",
			Required:     true,
			ErrorMessage: "1",
		},
	}
}
