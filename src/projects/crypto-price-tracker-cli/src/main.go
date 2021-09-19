package main

import (
	"./client"
	"flag"
	"fmt"
	"log"
)

func main() {
	fiatCurrency := flag.String(
		"fiat", "EUR", "The Fiat currency to fetch the corresponding Crypto price for i.e. EUR,USD",
	)

	cryptoIds := flag.String(
		"crypto", "BTC", "Comma seperated cryptocurrencies for conversion i.e. BTC,ETH",
	)

	flag.Parse()
	result, err := client.FetchCryptoResponse(*fiatCurrency, *cryptoIds)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Print(result)
}
