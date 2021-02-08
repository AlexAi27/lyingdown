package assets

import (
	"encoding/json"
	"math"
)

type Spend struct {
	BeginYear   int     `json:"beginYear"`
	EndYear     int     `json:"endYear"`
	Spend       float64 `json:"spend"`
	GrouwthRate float64 `json:"grouwthRate"`
}

func (s *Spend) GetStruct() interface{} {
	return *s
}

func (s *Spend) GetDescription() string {
	bytes, _ := json.MarshalIndent(s, "", "    ")
	return string(bytes)
}

func (s *Spend) GetType() string {
	return "Spend"
}

func (s *Spend) CashAt(year int) float64 {
	return s.At(year)
}

func (s *Spend) CashRangeAt(beginYear, endYear int) []float64 {
	years := endYear - beginYear + 1
	if years <= 0 || beginYear < 0 {
		return make([]float64, 0)
	}

	return s.RangeAt(beginYear, endYear)
}

func (s *Spend) At(year int) float64 {
	years := year - s.BeginYear + 1
	if years <= 0 {
		return 0
	}
	if year >= s.EndYear {
		years = s.EndYear - s.BeginYear + 1
	}

	rate := 1 + s.GrouwthRate
	return -(s.Spend * (1 - math.Pow(rate, float64(years))) / (1 - rate))
}

func (s *Spend) RangeAt(beginYear, endYear int) []float64 {
	ret := make([]float64, 0)
	if beginYear < 0 {
		return ret
	}

	rate := 1 + s.GrouwthRate
	totalRate := 1.0
	for year := beginYear; year <= endYear; year++ {
		if year < s.BeginYear || year > s.EndYear {
			ret = append(ret, 0)
			continue
		}
		ret = append(ret, -s.Spend*totalRate)
		totalRate *= rate
	}
	for idx := range ret {
		if idx == 0 {
			continue
		}
		ret[idx] += ret[idx-1]
	}
	return ret
}
