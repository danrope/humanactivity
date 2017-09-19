package util

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Read sample CSV data.  Results are row-based which emulates the expected input form
func ReadCSV() [][]float64 {
	var data [][]float64
	absPath, _ := filepath.Abs("../../data/onesample.csv")
	dat, err := ioutil.ReadFile(absPath)
	check(err)
	r := csv.NewReader(strings.NewReader(string(dat)))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var numbers []float64

		for _, elem := range record {
			i, err := strconv.ParseFloat(elem, 64)
			if err == nil {
				numbers = append(numbers, i)
			}
		}
		data = append(data, []float64(numbers))
	}

	return data
}
