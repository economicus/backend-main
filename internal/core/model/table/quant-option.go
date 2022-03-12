package table

import (
	"main/internal/core/model/quant"
	"time"
)

type QuantOption struct {
	QuantID             uint               `gorm:"primaryKey;column:quant_id"`
	Name                string             `gorm:"column:name"`
	MainSectors         []quant.MainSector `gorm:"constraint:OnDelete:CASCADE;"`
	NetRevenue          quant.IntPair      `gorm:"embedded;embeddedPrefix:net_revenue_"`
	NetRevenueRate      quant.DoublePair   `gorm:"embedded;embeddedPrefix:net_revenue_rate_"`
	NetProfit           quant.IntPair      `gorm:"embedded;embeddedPrefix:net_profit_"`
	NetProfitRate       quant.DoublePair   `gorm:"embedded;embeddedPrefix:net_profit_rate_"`
	DERatio             quant.DoublePair   `gorm:"embedded;embeddedPrefix:de_ratio_"`
	Per                 quant.DoublePair   `gorm:"embedded;embeddedPrefix:per_"`
	Psr                 quant.DoublePair   `gorm:"embedded;embeddedPrefix:psr_"`
	Pbr                 quant.DoublePair   `gorm:"embedded;embeddedPrefix:pbr_"`
	Pcr                 quant.DoublePair   `gorm:"embedded;embeddedPrefix:pcr_"`
	Operating           quant.DoublePair   `gorm:"embedded;embeddedPrefix:operating_"`
	Investing           quant.DoublePair   `gorm:"embedded;embeddedPrefix:investing_"`
	Financing           quant.DoublePair   `gorm:"embedded;embeddedPrefix:financing_"`
	DividendYield       quant.DoublePair   `gorm:"embedded;embeddedPrefix:dividend_yield_"`
	DividendPayoutRatio quant.DoublePair   `gorm:"embedded;embeddedPrefix:dividend_payout_ratio_"`
	Roa                 quant.DoublePair   `gorm:"embedded;embeddedPrefix:roa_"`
	Roe                 quant.DoublePair   `gorm:"embedded;embeddedPrefix:roe_"`
	MarketCap           quant.IntPair      `gorm:"embedded;embeddedPrefix:market_cap_"`
	StartDate           time.Time          `time_format:"2006-01-02T15:04:05.000Z"`
	EndDate             time.Time          `time_format:"2006-01-02T15:04:05.000Z"`
}
