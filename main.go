package main

import (
	"log"
	"strings"

	"github.com/Mitu217/csv_reader/csv"
)

const csvData = `
id,val1,val2
1,100,200
2,102,202
`

func main() {
	r := strings.NewReader(csvData)
	dec := csv.NewDecoder(r,
		DecodeTotal,
	)

	var objs []*csv.CsvObj
	if err := dec.Decode(&objs); err != nil {
		log.Fatal(err)
		return
	}
	for _, obj := range objs {
		log.Println(obj)
	}
}

func DecodeTotal(objs *[]*csv.CsvObj) error {
	for _, obj := range *objs {
		obj.Total = obj.Val1 + obj.Val2
	}
	return nil
}
