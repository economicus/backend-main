package request

import (
	"main/internal/core/model"
	"main/internal/pkg/objconv"
)

type EditQuantOptionReq struct {
	model.QuantOption
}

func (r *EditQuantOptionReq) ToMap() map[string]interface{} {
	return objconv.ToMap(r)
}
