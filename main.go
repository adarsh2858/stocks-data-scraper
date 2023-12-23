package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Stock struct {
	company, price, change string
}

var stocks []Stock

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
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: \n", err.Error())
	})

	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	link := e.Attr("href")

	// 	fmt.Printf("Link found: %q -> %s", e.Text, link)
	// })
	c.OnHTML("div#quote-header-info", func(e *colly.HTMLElement) {
		stock := Stock{}
		stock.company = e.ChildText("h1")
		fmt.Println("Company:", stock.company)
		stock.price = e.ChildText("fin-streamer[data-field='regularMarketPrice']")
		fmt.Println("Price:", stock.price)
		stock.change = e.ChildText("fin-streamer[data-field='regularMarketChangePercent']")
		fmt.Println("Change:", stock.change)

		stocks = append(stocks, stock)
	})
	c.Wait()

	// c.OnResponse(func(resp *colly.Response) {
	// 	file, err := os.Create("output.txt")
	// 	if err != nil {
	// 		log.Print(err.Error())
	// 	}
	// 	defer file.Close()

	// 	file.Write(resp.Body)

	// 	var stock Stock
	// 	// stock = Stock{
	// 	// 	company: company,
	// 	// 	price:   price,
	// 	// 	change:  change,
	// 	// }
	// 	stocks = append(stocks, stock)

	// })
	// handle the methods from colly collector like onError, onRequest => print the visiting url here
	// fetch the value from the corresponding html element tag and add to the struct field
	// append the stock data in the array
	// write the data to a csv file with the headers as company name, price and change

	for _, val := range tickers {
		url := "https://finance.yahoo.com/quote/" + val
		c.Visit(url)
	}

	csvFile, err := os.Create("result.csv")
	if err != nil {
		log.Fatal("Could not create the csv file \n", err.Error())
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	headers := []string{"company name", "price", "change"}
	csvWriter.Write(headers)

	for _, stock := range stocks {
		record := []string{
			stock.company,
			stock.price,
			stock.change,
		}
		csvWriter.Write(record)
	}

	defer csvWriter.Flush()
}
