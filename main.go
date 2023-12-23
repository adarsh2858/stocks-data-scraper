package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Stock struct {
	company, price, change string
}

func main() {
	// tickers list
	tickers := []string{
		"MSFT",
		"IBM",
		"GE",
		"UNP",
		"COST",
		"MCD",
	}

	c := colly.NewCollector()

	// visit each ticker by appending it to the base url
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})
	// handle the methods from colly collector like onError, onRequest => print the visiting url here
	// fetch the value from the corresponding html element tag and add to the struct field
	// append the stock data in the array

	for _, val := range tickers {
		url := "https://finance.yahoo.com/quote/" + val
		c.Visit(url)
	}
}
