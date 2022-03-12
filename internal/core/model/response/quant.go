package response

import (
	"main/internal/core/pb"
	"main/internal/pkg/objconv"
)

type ChartData struct {
	ProfitRateData []float32 `json:"profit_rate_data" bson:"chart" example:"8.31201046811529,15.13554790878776,-1.336521221573761,-1.42408166715555,10.420784591586559,8.305691643668455,17.68356243256443,9.407034979656027,-4.15162926200139,5.542443496088845,6.654446258518339"`
}

type QuantResponse struct {
	QuantID             uint      `json:"quant_id" bson:"quant_id" example:"5"`
	CumulativeReturn    float64   `json:"cumulative_return" bson:"cumulative_return" example:"15.95"`
	AnnualAverageReturn float64   `json:"annual_average_return" bson:"annual_average_return" example:"-2.21"`
	WinningPercentage   float64   `json:"winning_percentage" bson:"winning_percentage" example:"45.45"`
	MaxLossRate         float64   `json:"max_loss_rate" bson:"max_loss_rate" example:"-26.46"`
	HoldingsCount       int32     `json:"holdings_count" bson:"holdings_count" example:"7"`
	ChartData           ChartData `json:"chart_data" bson:"chart_data,inline"`
}

func NewQuantResultFromPB(pb *pb.QuantResult) *QuantResponse {
	return &QuantResponse{
		CumulativeReturn:    pb.CumulativeReturn,
		AnnualAverageReturn: pb.AnnualAverageReturn,
		WinningPercentage:   pb.WinningPercentage,
		MaxLossRate:         pb.MaxLossRate,
		HoldingsCount:       pb.HoldingsCount,
		ChartData: ChartData{
			ProfitRateData: pb.ChartData.ProfitRateData,
		},
	}
}

func (qr *QuantResponse) ToMap() map[string]interface{} {
	return objconv.ToMap(qr)
}
