package concurrency

import (
	"runtime"
	"time"
	"fmt"
	"net/http"
	"encoding/xml"
	"io/ioutil"
)

type QuoteResponse struct {
	Status string
	Name string
	LastPrice float32
	Change float32
	ChangePercent float32
	TimeStamp string
	MSDate float32
	MarketCap int
	Volume int
	ChangeYTD float32
	ChangePercentYTD float32
	High float32
	Low float32
	Open float32
}

/*
	1. Web Service Calls
	2. Add Concurrency
	3. Add Parallelism
*/
func AsynchronousWebServices() {
	// 4 Cores CPU
	runtime.GOMAXPROCS(4)

	start := time.Now()

	stockSymbols := []string {
		"googl",
		"msft",
		"aapl",
		"bbry",
		"hpq",
		"vz",
		"t",
		"tmus",
		"s",
	}

	/* ========================WITHOUT CONCURRENCY
		Alphabet Inc: 1136.59
		Microsoft Corp: 123.35
		Apple Inc: 185.72
		BlackBerry Ltd: 8.48
		HP Inc: 18.55
		: 0.00
		: 0.00
		: 0.00
		: 0.00
		Execution time: 7.070713629s
	
	for _, symbol := range stockSymbols {
		resp, _ := http.Get("http://dev.markitondemand.com/Api/v2/Quote?symbol=" + symbol)
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		quote := new (QuoteResponse)
		xml.Unmarshal(body, &quote)

		fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)
	}
	*/

	/* ========================WITH CONCURRENCY
		Verizon Communications Inc: 56.81
		T-Mobile US Inc: 73.42
		Apple Inc: 185.72
		BlackBerry Ltd: 8.48
		HP Inc: 18.56
		Microsoft Corp: 123.35
		Sprint Corp: 5.95
		Alphabet Inc: 1136.59
		AT&T Inc: 31.15
		Execution time: 3.873263889s
	*/
	numComplete := 0

	for _, symbol := range stockSymbols {
		go func(symbol string) {
			resp, _ := http.Get("http://dev.markitondemand.com/Api/v2/Quote?symbol=" + symbol)
			defer resp.Body.Close()

			body, _ := ioutil.ReadAll(resp.Body)

			quote := new (QuoteResponse)
			xml.Unmarshal(body, &quote)

			fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)
			numComplete++
		}(symbol)
	}
	
	for numComplete < len(stockSymbols) {
		time.Sleep(10 * time.Millisecond)
	}

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}