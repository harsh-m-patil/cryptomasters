package main

import (
	"flag"
	"fmt"

	"github.com/harsh-m-patil/cryptomasters/api"
)

func main() {
	from := flag.String("from", "BTC", "currency from which you want to convert")
	to := flag.String("to", "USD", "currency in which you want to convert")
	list := flag.Bool("list", false, "list all available currencies")

	flag.Parse()

	currencyMap := api.GetList()
	if *list {
		printAvailableCurrencies(currencyMap)
		return
	}
	_, fromOk := currencyMap[*from]
	_, toOk := currencyMap[*to]

	if fromOk && toOk {
		getCurrencyData(*from, *to)
	} else {
		fmt.Println("Invalid currency see list")
	}
}

func getCurrencyData(cryptoCurrency, currency string) {
	rate, err := api.GetRate(cryptoCurrency, currency)
	if err != nil {
		fmt.Println(fmt.Errorf("error(in get) occured %v", err))
	}

	fmt.Printf("The rate for %v is %.2f %s\n", rate.Currency, rate.Price, currency)
}

func printAvailableCurrencies(currMap map[string]string) {
	for key, desc := range currMap {
		fmt.Printf("%s : %s\n", key, desc)
	}
}
