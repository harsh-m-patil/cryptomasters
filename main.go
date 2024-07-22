package main

import (
	"fmt"
	"sync"

	"github.com/harsh-m-patil/cryptomasters/api"
)

func main() {
	currencies := []string{"BTC", "ETH", "BCH"}
	var wg sync.WaitGroup

	for _, currency := range currencies {
		wg.Add(1)
		go func() {
			getCurrencyData(currency)
			wg.Done()
		}()
	}

	wg.Wait()
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err != nil {
		fmt.Println(fmt.Errorf("error occured %v", err))
	}

	fmt.Printf("The rate for %v is %.2f $\n", rate.Currency, rate.Price)
}
