package model

import (
	"main/internal/core/model/table"
	"main/internal/pkg/objconv"
)

type Quants []Quant

type Quant struct {
	table.Quant
}

func NewQuant(userID uint, name string) *Quant {
	return &Quant{
		table.Quant{
			UserID: userID,
			Name:   name,
		},
	}
}

func (q *Quant) GetID() uint {
	return q.ID
}

func (q *Quant) GetOwnerID() uint {
	return q.UserID
}

func (q *Quant) ToMap() map[string]interface{} {
	return objconv.ToMap(q)
}

func (q *Quant) ToMapWithFields(fields []string) map[string]interface{} {
	return objconv.ToMapWithFields(q, fields)
}
