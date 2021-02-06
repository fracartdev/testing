package parserservice

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Order definisce la struttura per il JSON
type Order struct {
	ID         string `json:"id"`
	Item       string `json:"item"`
	User       string `json:"user"`
	Quantity   int    `json:"quantity"`
	City       string `json:"city"`
	Department string `json:"department"`
	Price      int    `json:"price"`
}

func readCSV() [][]string {
	csvFile, err := os.Open("app/parser/sample-csv/sample.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return records
}

// Parse dato un path converte il CSV in JSON
func Parse() []byte {
	var obj Order
	var report []Order

	records := readCSV()
	for _, rec := range records {
		obj.ID = rec[0]
		obj.Item = rec[1]
		obj.User = rec[2]
		obj.Quantity, _ = strconv.Atoi(rec[3])
		obj.City = rec[7]
		obj.Department = rec[8]
		obj.Price, _ = strconv.Atoi(rec[9])
		report = append(report, obj)
	}

	jsonData, err := json.Marshal(report)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return jsonData
}

// Print dato un JSON ne stampa il contenuto in un file
func Print(report []byte) {
	jsonFile, err := os.Create("sample.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(report)
	jsonFile.Close()
}
