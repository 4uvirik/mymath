package main

import "fmt"

type CurrencyConverter struct {
	Rate float64
}

func (c *CurrencyConverter) ConvertToUSD(raschet float64) float64 {
	return raschet / c.Rate
}

func (c *CurrencyConverter) ConvertToRUB(raschet float64) float64 {
	return c.Rate * raschet
}
func main() {
	converter := CurrencyConverter{Rate: 75.5}
	dol := converter.ConvertToUSD(755)
	fmt.Println(dol)

	rub := converter.ConvertToRUB(2)
	fmt.Println(rub)
}
