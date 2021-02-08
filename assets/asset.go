package assets

type Asset interface {
	GetStruct() interface{}
	GetDescription() string
	GetType() string
	// 现金
	CashAt(int) float64
	CashRangeAt(int, int) []float64
	// 总资产（负债）
	At(int) float64
	RangeAt(int, int) []float64
}
