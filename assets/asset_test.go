package assets

import (
	"math"
	"testing"
)

func floatEqual(a, b float64) bool {
	return math.Abs(a-b) < 0.00001
}

func TestIncome(t *testing.T) {
	income := &Income{
		Year:        22,
		Income:      30,
		GrouwthRate: 0.1,
		IsCash:      true,
	}
	if !floatEqual(income.At(0), 0) {
		t.Errorf("income.At(0) expected 0 but get %f", income.At(0))
	}
	if !floatEqual(income.At(21), 0) {
		t.Errorf("income.At(21) expected 0 but get %f", income.At(21))
	}
	if !floatEqual(income.At(22), 30) {
		t.Errorf("income.At(22) expected 30 but get %f", income.At(22))
	}
	if !floatEqual(income.At(23), 33) {
		t.Errorf("income.At(23) expected 33 but get %f", income.At(23))
	}
	if !floatEqual(income.At(24), 36.3) {
		t.Errorf("income.At(24) expected 36.3 but get %f", income.At(24))
	}
	if !floatEqual(income.At(50), 432.629808319497702113028192243) {
		t.Errorf("income.At(50) expected 432.629808319497702113028192243 but get %f", income.At(50))
	}
}

func TestSalary(t *testing.T) {
	salary := &Salary{
		BeginYear:   22,
		EndYear:     50,
		Salary:      30,
		GrouwthRate: 0.1,
	}
	if !floatEqual(salary.At(0), 0) {
		t.Errorf("salary.At(0) expected 0 but get %f", salary.At(0))
	}
	if !floatEqual(salary.At(21), 0) {
		t.Errorf("salary.At(21) expected 0 but get %f", salary.At(21))
	}
	if !floatEqual(salary.At(22), 30) {
		t.Errorf("salary.At(22) expected 30 but get %f", salary.At(22))
	}
	if !floatEqual(salary.At(23), 63) {
		t.Errorf("salary.At(23) expected 63 but get %f", salary.At(23))
	}
	if !floatEqual(salary.At(24), 99.3) {
		t.Errorf("salary.At(24) expected 99.3 but get %f", salary.At(24))
	}
	if !floatEqual(salary.At(50), 4458.927891514474723243310114673) {
		t.Errorf("salary.At(50) expected 4458.927891514474723243310114673 but get %f", salary.At(50))
	}
	if !floatEqual(salary.At(51), 4458.927891514474723243310114673) {
		t.Errorf("salary.At(51) expected 4458.927891514474723243310114673 but get %f", salary.At(51))
	}
}
