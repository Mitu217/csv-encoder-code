# csv-reader

this repo is custom csv reader sample


## Usage

```golang
func main() {
	r := strings.NewReader(csvData)
	dec := csv.NewDecoder(r,
		// Extra decode params
		DecodeTotal,
		// Validate params
		ValidateTotal
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

func ValidateTotal(objs *[]*csv.CsvObj) error {
	for _, obj := range *objs {
		if obj.Total < 100 {
			return errors.New("invalid total params")
		}
	}
	return nil
}
```
