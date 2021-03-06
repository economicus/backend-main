package response

import "main/internal/core/model/table"

type ProfileQuantChartResponse struct {
	QuantID uint      `json:"quant_id" bson:"quant_id" example:"5"`
	Chart   []float32 `json:"chart" bson:"chart" example:"8.31201046811529,15.13554790878776,-1.336521221573761,-1.42408166715555,10.420784591586559,8.305691643668455,17.68356243256443,9.407034979656027,-4.15162926200139,5.542443496088845,6.654446258518339"`
}

type ProfileQuantResponse struct {
	ProfileQuantChartResponse
	Name        string `json:"name" example:"model name"`
	Description string `json:"description" example:"model description"`
}

type ProfileResponse struct {
	User  table.User             `json:"user"`
	Quant []ProfileQuantResponse `json:"quant"`
}
