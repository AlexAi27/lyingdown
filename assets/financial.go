package assets

import (
	"encoding/json"
	"math"
)

type Financial struct {
	BeginYear   int        `json:"beginYear"`
	Base        float64    `json:"base"`
	GrouwthRate float64    `json:"grouwthRate"`
	AutoInvest  AutoInvest `json:"autoInvest"`
}

type AutoInvest struct {
	Enable      bool    `json:"enable"`
	EndYear     int     `json:"endYear"`
	GrouwthRate float64 `json:"grouwthRate"`
}

func (f *Financial) GetStruct() interface{} {
	return *f
}

func (f *Financial) GetDescription() string {
	bytes, _ := json.MarshalIndent(f, "", "    ")
	return string(bytes)
}

func (f *Financial) GetType() string {
	return "Financial"
}

func (f *Financial) CashAt(year int) float64 {
	return f.At(year)
}

func (f *Financial) CashRangeAt(beginYear, endYear int) []float64 {
	return f.RangeAt(beginYear, endYear)
}

func (f *Financial) At(year int) float64 {
	years := year - f.BeginYear
	if years < 0 {
		return 0
	}
	finanRate := 1 + f.GrouwthRate
	if !f.AutoInvest.Enable {
		return f.Base * math.Pow(finanRate, float64(years))
	}

	invRate := 1 + f.AutoInvest.GrouwthRate
	accRate := 0.0
	n := 0
	if year <= f.AutoInvest.EndYear {
		n = years
	} else {
		n = f.AutoInvest.EndYear - f.BeginYear
	}
	rate := math.Pow(finanRate, float64(n))
	for i := 0; i < n; i++ {
		accRate += rate
		rate = rate * invRate / finanRate
	}
	if years < n {
		accRate *= math.Pow(finanRate, float64(years-n))
	}

	return f.Base * accRate
}

func (f *Financial) RangeAt(beginYear, endYear int) []float64 {
	ret := make([]float64, 0)
	if beginYear < 0 {
		return ret
	}

	for year := beginYear; year <= endYear; year++ {
		ret = append(ret, f.At(year))
	}
	for idx := range ret {
		if idx == 0 {
			continue
		}
		ret[idx] += ret[idx-1]
	}
	return ret
}
