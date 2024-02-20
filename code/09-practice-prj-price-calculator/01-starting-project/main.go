package main

import "example.com/price-calculator/tax"

func main() {
	t := tax.NewTax(0.12)
	t.Process()
	t = tax.NewTax(0.15)
	t.Process()
	t = tax.NewTax(0.18)
	t.Process()
}
