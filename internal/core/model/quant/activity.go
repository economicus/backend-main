package quant

import "main/internal/core/pb"

type Activities struct {
	Operating DoublePair `gorm:"embedded;embeddedPrefix:operating_" json:"operating"`
	Investing DoublePair `gorm:"embedded;embeddedPrefix:investing_" json:"investing"`
	Financing DoublePair `gorm:"embedded;embeddedPrefix:financing_" json:"financing"`
}

func (a *Activities) ToPB() *pb.Activities {
	return &pb.Activities{
		Operating: a.Operating.ToPB(),
		Investing: a.Investing.ToPB(),
		Financing: a.Financing.ToPB(),
	}
}
