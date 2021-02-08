package assets

import (
	"encoding/json"
	"math"
)

type Salary struct {
	BeginYear   int     `json:"beginYear"`
	EndYear     int     `json:"endYear"`
	Salary      float64 `json:"salary"`
	GrouwthRate float64 `json:"grouwthRate"`
}

func (s *Salary) GetStruct() interface{} {
	return *s
}

func (s *Salary) GetDescription() string {
	bytes, _ := json.MarshalIndent(s, "", "    ")
	return string(bytes)
}

func (s *Salary) GetType() string {
	return "Salary"
}

func (s *Salary) CashAt(year int) float64 {
	return s.At(year)
}

func (s *Salary) CashRangeAt(beginYear, endYear int) []float64 {
	years := endYear - beginYear + 1
	if years <= 0 || beginYear < 0 {
		return make([]float64, 0)
	}

	return s.RangeAt(beginYear, endYear)
}

func (s *Salary) At(year int) float64 {
	years := year - s.BeginYear + 1
	if years <= 0 {
		return 0
	}
	if year >= s.EndYear {
		years = s.EndYear - s.BeginYear + 1
	}

	rate := 1 + s.GrouwthRate
	return s.Salary * (1 - math.Pow(rate, float64(years))) / (1 - rate)
}

func (s *Salary) RangeAt(beginYear, endYear int) []float64 {
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
		ret = append(ret, s.Salary*totalRate)
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
