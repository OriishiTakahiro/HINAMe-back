package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	filename := flag.String("f", "sample.csv", "Specify CSV file")
	flag.Parse()

	csvDat := readCSV(*filename)[1:]
	shelterRecords := make([]string, len(csvDat))
	for i, record := range csvDat {
		name := regexp.QuoteMeta(record[1])
		shelterRecords[i] = fmt.Sprintf("(\"%s\", %s, %s)", name, record[3], record[4])
	}

	insertShelters := "INSERT INTO hiname.shelters (name, latitude, longitude) VALUES \n" + strings.Join(shelterRecords, ",\n") + ";"
	fmt.Println(insertShelters)
}

func readCSV(filename string) (records [][]string) {
	records = make([][]string, 100)
	fr, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fr.Close()
	r := csv.NewReader(fr)
	records, err = r.ReadAll()
	return
}
