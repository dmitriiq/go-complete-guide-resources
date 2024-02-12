package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount int
	const expectedReturnRate = 5.5
	var years = 10

	fmt.Print("investmentAmount: ")
	fmt.Scan(&investmentAmount)

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Println(futureValue)
}
