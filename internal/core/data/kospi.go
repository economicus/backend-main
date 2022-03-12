package data

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

var KospiData = *NewKospi()
var FixedKD []float32

const layout = "2006-01-02T15:04:05.000Z"

type Kospi struct {
	Date     map[time.Time]int
	IndexVal []float32
}

func NewKospi() *Kospi {
	file, err := os.Open("internal/core/data/kospi.csv")
	if err != nil {
		log.Panicf("error while open kospi: %v", err)
	}

	r := csv.NewReader(bufio.NewReader(file))
	rows, err := r.ReadAll()
	if err != nil {
		log.Panicf("error while reading kospi: %v", err)
	}

	k := Kospi{
		Date:     make(map[time.Time]int),
		IndexVal: []float32{},
	}

	for i := range rows {
		t, err := time.Parse(layout, rows[i][0])
		if err != nil {
			log.Panicf("error while parsing time: %v", err)
		}
		f, err := strconv.ParseFloat(rows[i][1], 32)
		if err != nil {
			log.Panicf("error while parsing index value: %v", err)
		}
		k.Date[t] = i
		k.IndexVal = append(k.IndexVal, float32(f))
		FixedKD = append(FixedKD, float32(f))
	}
	return &k
}
