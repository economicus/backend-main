package request

import "main/internal/core/model"

type CreateQuantReq struct {
	model.QuantOption
}

type EditQuantReq struct {
	Active      bool   `json:"active,omitempty" example:"true"`
	Name        string `json:"name,omitempty" example:"New model name"`
	Description string `json:"description,omitempty" example:"New model description"`
}
