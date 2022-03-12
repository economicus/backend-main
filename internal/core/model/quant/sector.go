package quant

type MainSector struct {
	QuantOptionID uint   `json:"-"`
	Name          string `json:"name"`
}

func NewMainSectors(quantID uint, ms []string) []MainSector {
	var res []MainSector

	for _, val := range ms {
		res = append(res, MainSector{quantID, val})
	}
	return res
}

func (ms *MainSector) ToString() string {
	return ms.Name
}
