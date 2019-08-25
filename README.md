# csv-reader

this repo is custom csv reader sample


## Usage

```golang
type CsvObj struct {
	ID    int `csv:"id"`
	Val1  int `csv:"val1"`
	Val2  int `csv:"val2"`
	Total int `csv:"total"`
}

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
```
