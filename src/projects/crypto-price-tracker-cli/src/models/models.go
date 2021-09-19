package models

import (
	"fmt"
	"github.com/alexeyco/simpletable"
)

type ApiResponse []struct{
	Name string `json:"name"`
	Price string `json:"price"`
	Symbol string `json:"symbol"`
	Status string `json:"status"`
	PriceTimestamp string `json:"price_timestamp"`
	CirculatingSupply string `json:"circulating_supply"`
	Rank string `json:"rank"`
	High string `json:"high"`
}

func (resp ApiResponse) ResponseTable() string{
	// Initialize the table
	table:= simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "S/N"},
			{Align: simpletable.AlignCenter, Text: "Name"},
			{Align: simpletable.AlignCenter, Text: "Symbol"},
			{Align: simpletable.AlignCenter, Text: "Price"},
			{Align: simpletable.AlignCenter, Text: "Status"},
			{Align: simpletable.AlignCenter, Text: "Price Timestamp"},
			{Align: simpletable.AlignCenter, Text: "Circulating Supply"},
			{Align: simpletable.AlignCenter, Text: "Rank"},
			{Align: simpletable.AlignCenter, Text: "High"},
		},
	}

	i:= 0
	for _,row := range resp{
		// Create a row
		r:= []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d",i)},
			{Align: simpletable.AlignCenter, Text: row.Name},
			{Align: simpletable.AlignCenter, Text: row.Symbol},
			{Align: simpletable.AlignCenter, Text: row.Price},
			{Align: simpletable.AlignCenter, Text: row.Status},
			{Align: simpletable.AlignCenter, Text: row.PriceTimestamp},
			{Align: simpletable.AlignCenter, Text: row.CirculatingSupply},
			{Align: simpletable.AlignCenter, Text: row.Rank},
			{Align: simpletable.AlignCenter, Text: row.High},
		}

		// Append the row to the table row cells
		table.Body.Cells = append(table.Body.Cells, r)
		i++
	}

	return table.String()
}