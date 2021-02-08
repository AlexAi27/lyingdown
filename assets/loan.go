package assets

import (
	"encoding/json"
	"math"
)

const (
	InterestTypeEqualAll       = "等额本息"
	InterestTypeEqualPrincipal = "等额本金"
)

type Loan struct {
	BeginYear    int       `json:"beginYear"`
	EndYear      int       `json:"endYear"`
	Loan         float64   `json:"loan"`
	InterestRate float64   `json:"interestRate"`
	InterestType string    `json:"interestType"`
	accCash      []float64 `json:"-"`
}

func (l *Loan) GetStruct() interface{} {
	return *l
}

func (l *Loan) GetDescription() string {
	bytes, _ := json.MarshalIndent(l, "", "    ")
	return string(bytes)
}

func (l *Loan) GetType() string {
	return "Loan"
}

func (l *Loan) CashAt(year int) float64 {
	return l.At(year)
}

func (l *Loan) CashRangeAt(beginYear, endYear int) []float64 {
	return l.RangeAt(beginYear, endYear)
}

func (l *Loan) At(year int) float64 {
	if l.accCash == nil || len(l.accCash) == 0 {
		l.calcCashCache()
	}
	if year < l.BeginYear {
		return 0
	} else if year > l.EndYear {
		return -l.accCash[len(l.accCash)-1]
	}
	return -l.accCash[year]
}

func (l *Loan) RangeAt(beginYear, endYear int) []float64 {
	ret := make([]float64, 0)
	if beginYear < 0 {
		return ret
	}
	for year := beginYear; year <= endYear; year++ {
		ret = append(ret, l.At(year))
	}
	return ret
}

func (l *Loan) equalAllCashAt(year int) float64 {
	r := l.InterestRate/12 + 1
	beginMo := 12*(year-1) + 1
	endMo := 12 * year
	ret := 0.0
	for m := beginMo; m <= endMo; m++ {
		ret += math.Pow(r, float64(m)) / (math.Pow(r, float64(m)) - 1)
	}
	ret *= l.Loan * l.InterestRate / 12
	return ret
}

func (l *Loan) equalPrincipalCashAt(year int) float64 {
	r := l.InterestRate / 12
	beginMo := 12*(year-1) + 1
	endMo := 12 * year
	totalMo := (l.EndYear - l.BeginYear + 1) * 12
	ret := l.Loan / float64(totalMo) * 12
	for m := beginMo; m < endMo; m++ {
		ret += (l.Loan - l.Loan/float64(totalMo)*(float64(m)-1)) * r
	}
	return ret
}

func (l *Loan) calcCashCache() {
	l.accCash = make([]float64, 0)
	for y := l.BeginYear; y <= l.EndYear; y++ {
		if l.InterestType == InterestTypeEqualAll {
			l.accCash = append(l.accCash, l.equalAllCashAt(y))
		} else {
			l.accCash = append(l.accCash, l.equalPrincipalCashAt(y))
		}
	}
	for idx := range l.accCash {
		if idx == 0 {
			continue
		}
		l.accCash[idx] += l.accCash[idx-1]
	}
}
