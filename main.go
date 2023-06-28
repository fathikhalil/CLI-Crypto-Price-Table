package main

import (
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// Fetch Data From API
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin%2Cethereum%2Csolana%2Cripple%2Cfilecoin&vs_currencies=usd")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)

	}

	// Parse unmarshal JSON strings
	var objMap map[string]interface{}
	err = json.Unmarshal((body), &objMap)

	if err != nil {
		fmt.Println("Json Error Decode ")
	}

	// Data to Variables
	Price1 := objMap["bitcoin"].(map[string]interface{})["usd"]
	Price2 := objMap["ethereum"].(map[string]interface{})["usd"]
	Price3 := objMap["solana"].(map[string]interface{})["usd"]
	Price4 := objMap["filecoin"].(map[string]interface{})["usd"]
	Price5 := objMap["ripple"].(map[string]interface{})["usd"]

	// Timestamp
	now := time.Now().UTC()
	now1 := now.Format("15:04:05")
	date := now.Format("Jan 2, 2006")

	// Table write
	t := table.NewWriter()
	t.SetAutoIndex(true)

	t.SetTitle("Top 5 Crypto by Market Cap ")
	t.AppendHeader(table.Row{"Symbol", "Price"})
	t.AppendRow(table.Row{"BTC", Price1})
	t.AppendRow(table.Row{"ETH", Price2})
	t.AppendRows([]table.Row{{"SOL", Price3},
		{"Filecoin", Price4}, {"Ripple", Price5}})

	t.SetStyle(table.StyleLight)

	t.Render()
	t.SetCaption(now1)
	t.SetCaption(date)

	fmt.Println(t.Render())

}
