package response

import "main/internal/core/model"

type LabQuant struct {
	QuantID             uint    `json:"quant_id" example:"1"`
	Name                string  `json:"name" example:"model name"`
	CumulativeReturn    float32 `json:"cumulative_return" example:"128.2"`
	AnnualAverageReturn float32 `json:"annual_average_return" example:"16.0"`
	WinningPercentage   float32 `json:"winning_percentage" example:"66.66"`
	MaxLossRate         float32 `json:"max_loss_rate" example:"-29.11"`
	HoldingsCount       int32   `json:"holdings_count" example:"22"`
}

type LabData struct {
	Option model.QuantOption `json:"option"`
	Chart  ChartData         `json:"chart"`
}
