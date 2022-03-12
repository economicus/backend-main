package table

import (
	"gorm.io/gorm"
	"time"
)

type Quant struct {
	ID                  uint           `gorm:"primaryKey" json:"id" example:"1"`
	CreatedAt           time.Time      `json:"-" swaggerignore:"true"`
	UpdatedAt           time.Time      `json:"-" swaggerignore:"true"`
	DeletedAt           gorm.DeletedAt `gorm:"index;" json:"-" swaggerignore:"true"`
	UserID              uint           `json:"user_id"`
	Name                string         `gorm:"foreignKey:Name;column:name;not null;unique" json:"name" example:"quant model name"`
	Description         string         `gorm:"column:description" json:"description" example:"quant model description"`
	QuantOption         QuantOption    `gorm:"foreignKey:QuantID;constraint:OnDelete:CASCADE;foreignKey:QuantID;references:ID" json:"-" swaggerignore:"true"`
	CumulativeReturn    float32        `gorm:"column:cumulative_return" json:"cumulative_return" example:"128.2"`
	AnnualAverageReturn float32        `gorm:"column:annual_average_return" json:"annual_average_return" example:"16.0"`
	WinningPercentage   float32        `gorm:"column:winning_percentage" json:"winning_percentage" example:"66.66"`
	MaxLossRate         float32        `gorm:"column:max_loss_rate" json:"max_loss_rate" example:"-29.11"`
	HoldingsCount       int32          `gorm:"column:holdings_count" json:"holdings_count" example:"22"`
	ChartID             string         `gorm:"column:chart_id" json:"-" swaggerignore:"true"`
	Active              bool           `gorm:"column:active;default:false" json:"-" swaggerignore:"true"`
	LikedUsers          []*User        `gorm:"many2many:user_favorite_quants;" json:"-" swaggerignore:"true"`
	Private             bool           `gorm:"column:private;default:false" json:"-" swaggerignore:"true"`
	Comments            []Comment      `gorm:"constraint:OnDelete:CASCADE;" json:"-" swaggerignore:"true"`
}
