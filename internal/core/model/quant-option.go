package model

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"main/internal/core/model/quant"
	"main/internal/core/model/table"
	"main/internal/core/pb"
	"main/internal/pkg/objconv"
	"time"
)

type QuantOption struct {
	QuantID             uint             `json:"quant_id" example:"1"`
	Name                string           `json:"name"`
	MainSectors         []string         `json:"main_sectors"`
	NetRevenue          quant.IntPair    `json:"net_revenue"`
	NetRevenueRate      quant.DoublePair `json:"net_revenue_rate"`
	NetProfit           quant.IntPair    `json:"net_profit"`
	NetProfitRate       quant.DoublePair `json:"net_profit_rate"`
	DERatio             quant.DoublePair `json:"de_ratio"`
	Per                 quant.DoublePair `json:"per"`
	Psr                 quant.DoublePair `json:"psr"`
	Pbr                 quant.DoublePair `json:"pbr"`
	Pcr                 quant.DoublePair `json:"pcr"`
	Operating           quant.DoublePair `json:"operating"`
	Investing           quant.DoublePair `json:"investing"`
	Financing           quant.DoublePair `json:"financing"`
	DividendYield       quant.DoublePair `json:"dividend_yield"`
	DividendPayoutRatio quant.DoublePair `json:"dividend_payout_ratio"`
	Roa                 quant.DoublePair `json:"roa"`
	Roe                 quant.DoublePair `json:"roe"`
	MarketCap           quant.IntPair    `json:"market_cap"`
	StartDate           time.Time        `time_format:"2006-01-02T15:04:05.000Z" json:"start_date" example:"2016-03-31T00:00:00.000Z"`
	EndDate             time.Time        `time_format:"2006-01-02T15:04:05.000Z" json:"end_date" example:"2021-03-31T00:00:00.000Z"`
}

func NewQuantOption(tqo *table.QuantOption) *QuantOption {
	var ms []string

	for i := range tqo.MainSectors {
		ms = append(ms, tqo.MainSectors[i].ToString())
	}
	return &QuantOption{
		QuantID:             tqo.QuantID,
		Name:                tqo.Name,
		MainSectors:         ms,
		NetRevenue:          tqo.NetRevenue,
		NetRevenueRate:      tqo.NetRevenueRate,
		NetProfit:           tqo.NetProfit,
		NetProfitRate:       tqo.NetProfitRate,
		DERatio:             tqo.DERatio,
		Per:                 tqo.Per,
		Psr:                 tqo.Psr,
		Pbr:                 tqo.Pbr,
		Pcr:                 tqo.Pcr,
		Operating:           tqo.Operating,
		Investing:           tqo.Investing,
		Financing:           tqo.Financing,
		DividendYield:       tqo.DividendYield,
		DividendPayoutRatio: tqo.DividendPayoutRatio,
		Roa:                 tqo.Roa,
		Roe:                 tqo.Roe,
		MarketCap:           tqo.MarketCap,
		StartDate:           tqo.StartDate,
		EndDate:             tqo.EndDate,
	}
}

func (q *QuantOption) ToTable() *table.QuantOption {
	ms := q.MainSectors
	var tms []quant.MainSector
	for i := range ms {
		tms = append(tms, quant.MainSector{QuantOptionID: q.QuantID, Name: ms[i]})
	}
	opt := table.QuantOption{
		QuantID:             q.QuantID,
		Name:                q.Name,
		MainSectors:         tms,
		NetRevenue:          q.NetRevenue,
		NetRevenueRate:      q.NetRevenueRate,
		NetProfit:           q.NetProfit,
		NetProfitRate:       q.NetProfitRate,
		DERatio:             q.DERatio,
		Per:                 q.Per,
		Psr:                 q.Psr,
		Pbr:                 q.Pbr,
		Pcr:                 q.Pcr,
		Operating:           q.Operating,
		Investing:           q.Investing,
		Financing:           q.Financing,
		DividendYield:       q.DividendYield,
		DividendPayoutRatio: q.DividendPayoutRatio,
		Roa:                 q.Roa,
		Roe:                 q.Roe,
		MarketCap:           q.MarketCap,
		StartDate:           q.StartDate,
		EndDate:             q.EndDate,
	}
	return &opt
}

func (q *QuantOption) ToRequest() *pb.QuantRequest {
	return &pb.QuantRequest{
		MainSector:     q.MainSectors,
		NetRevenue:     q.NetRevenue.ToPB(),
		NetRevenueRate: q.NetRevenueRate.ToPB(),
		NetProfit:      q.NetProfit.ToPB(),
		NetProfitRate:  q.NetProfitRate.ToPB(),
		DeRatio:        q.DERatio.ToPB(),
		Per:            q.Per.ToPB(),
		Psr:            q.Psr.ToPB(),
		Pbr:            q.Pbr.ToPB(),
		Pcr:            q.Pcr.ToPB(),
		Activities: &pb.Activities{
			Operating: q.Operating.ToPB(),
			Investing: q.Investing.ToPB(),
			Financing: q.Financing.ToPB(),
		},
		DividendYield:       q.DividendYield.ToPB(),
		DividendPayoutRatio: q.DividendPayoutRatio.ToPB(),
		Roa:                 q.Roa.ToPB(),
		Roe:                 q.Roe.ToPB(),
		MarketCap:           q.MarketCap.ToPB(),
		StartDate:           timestamppb.New(q.StartDate),
		EndDate:             timestamppb.New(q.EndDate),
	}
}

func (q *QuantOption) ToMap() map[string]interface{} {
	return objconv.ToMap(q)
}

func (q *QuantOption) TableName() string {
	return "quant_options"
}
