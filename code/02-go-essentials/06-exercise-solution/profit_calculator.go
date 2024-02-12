package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getUserInpu("Revenue: ")

	expenses = getUserInpu("Expenses: ")

	taxRate = getUserInpu("Tax Rate: ")

	ebt, profit, ratio := calc(revenue, expenses, taxRate)

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

func calc(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

func getUserInpu(text string) (value float64) {
	fmt.Print("Tax Rate: ")
	fmt.Scan(&value)
	return
}
