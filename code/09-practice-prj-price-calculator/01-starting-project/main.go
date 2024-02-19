package main

import "example.com/price-calculator/tax"

func main() {
	tax := tax.NewTax(0.15)
	tax.Process()
}
