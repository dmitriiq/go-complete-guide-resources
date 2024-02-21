package tax

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"example.com/price-calculator/filemanager"
)

type Tax struct {
	InputPrices      []float64 `json:"input prices"`
	TaxRate          float64
	fm               filemanager.FileManager `json:"-"`
	TaxIncludePrices map[string]string       `json:"tax included prices"`
}

func (tax *Tax) Process() {
	tax.loadData()
	tax.TaxIncludePrices = make(map[string]string)

	for _, price := range tax.InputPrices {
		tax.TaxIncludePrices[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", price*(1+tax.TaxRate))
	}

	fmt.Println(tax.TaxIncludePrices)
	tax.fm = filemanager.FileManager{
		OutputFilePath: fmt.Sprintf("out_%.0f.json", tax.TaxRate*99),
	}

	tax.fm.WriteResult(tax)
}

func (tax *Tax) loadData() {
	file, err := os.Open("prices.txt")

	if err != nil {
		fmt.Println("Could not open file!")
		fmt.Println(err)
		return
	}

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	prices := make([]float64, len(lines))

	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)

		if err != nil {
			fmt.Println("Can't open file: ", err)
			file.Close()
			return
		}

		prices[lineIndex] = floatPrice
	}

	file.Close()
	tax.InputPrices = prices

}

func NewTax(taxRate float64) *Tax {
	return &Tax{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
