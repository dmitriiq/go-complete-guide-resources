package tax

import "fmt"

type Tax struct {
	InputPrices []float64
	TaxRate     float64
}

func (tax Tax) Process() {
	result := make(map[string]float64)

	for _, price := range tax.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + tax.TaxRate)
	}

	fmt.Println(result)
}

func NewTax(taxRate float64) *Tax {
	return &Tax{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
