package quant

import "main/internal/core/pb"

type IntPair struct {
	Max int64 `json:"max" example:"-10000"`
	Min int64 `json:"min" example:"1000000"`
}

func (ip *IntPair) ToPB() *pb.IntPair {
	return &pb.IntPair{
		Max: ip.Max,
		Min: ip.Min,
	}
}

type DoublePair struct {
	Max float32 `json:"max" example:"-100.01"`
	Min float32 `json:"min" example:"100.01"`
}

func (dp *DoublePair) ToPB() *pb.DoublePair {
	return &pb.DoublePair{
		Max: dp.Max,
		Min: dp.Min,
	}
}
