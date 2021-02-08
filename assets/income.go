package assets

import (
	"encoding/json"
	"math"
)

type Income struct {
	Year        int     `json:"year"`
	Income      float64 `json:"income"`
	GrouwthRate float64 `json:"grouwthRate"`
	IsCash      bool    `json:"isCash"`
}

func (ic *Income) GetStruct() interface{} {
	return *ic
}

func (ic *Income) GetDescription() string {
	bytes, _ := json.MarshalIndent(ic, "", "    ")
	return string(bytes)
}

func (ic *Income) GetType() string {
	return "Income"
}

func (ic *Income) CashAt(year int) float64 {
	if !ic.IsCash {
		return 0
	}
	return ic.At(year)
}

func (ic *Income) CashRangeAt(beginYear, endYear int) []float64 {
	years := endYear - beginYear + 1
	if years <= 0 || beginYear < 0 {
		return make([]float64, 0)
	}

	if !ic.IsCash {
		return make([]float64, years)
	}
	return ic.RangeAt(beginYear, endYear)
}

func (ic *Income) At(year int) float64 {
	years := year - ic.Year
	if years < 0 {
		return 0
	}
	rate := 1 + ic.GrouwthRate
	return ic.Income * math.Pow(rate, float64(years))
}

func (ic *Income) RangeAt(beginYear, endYear int) []float64 {
	ret := make([]float64, 0)
	if beginYear < 0 {
		return ret
	}
	rate := 1 + ic.GrouwthRate
	totalRate := 1.0
	for year := beginYear; year <= endYear; year++ {
		if year < ic.Year {
			ret = append(ret, 0)
			continue
		}
		ret = append(ret, ic.Income*totalRate)
		totalRate *= rate
	}
	return ret
}
