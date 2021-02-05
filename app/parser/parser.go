package parser

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Order definisce la struttura per il JSON
type Order struct {
	ID         int    `json:"id"`
	Item       string `json:"item"`
	User       string `json:"user"`
	Qty        int    `json:"quantity"`
	City       string `json:"city"`
	Department string `json:"department"`
	Price      int    `json:"price"`
}

func readCSV(file string) [][]string {
	csvFile, err := os.Open(file)
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

func parse(file string) []byte {
	var obj Order
	var report []Order

	records := readCSV(file)
	for _, rec := range records {
		obj.ID, _ = strconv.Atoi(rec[0])
		obj.Item = rec[1]
		obj.User = rec[2]
		obj.Qty, _ = strconv.Atoi(rec[3])
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

func print(report []byte) {
	jsonFile, err := os.Create("sample.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(report)
	jsonFile.Close()
}
