package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Teamdetail struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (t Teamdetail) toString() string {
	bytes, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}

func getTeamdetails() []Teamdetail {
	teamdetails := make([]Teamdetail, 3)
	raw, err := ioutil.ReadFile("jsonFile.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(raw, &teamdetails)
	return teamdetails
}

func main() {
	fmt.Println("Reading a text file")
	readTxtFile()
	fmt.Println("Reading a csv file")
	readCsvFile()
	fmt.Println("Reading a json file")
	readjsonFile()
}

func readTxtFile() {
	//fmt.Println("Read a text file")
	data, err := ioutil.ReadFile("txtFile.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}

func readCsvFile() {
	csv_file, ferr := os.Open("csvFile.csv")
	if ferr != nil {
		fmt.Println("File reading error", ferr)
	}
	r := csv.NewReader(csv_file)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("File reading error", err)
		}
		fmt.Println(record)
	}
}

func readjsonFile() {
	teamdetails := getTeamdetails()
	fmt.Println(teamdetails)
	for _, te := range teamdetails {
		fmt.Println(te.toString())
	}
}
