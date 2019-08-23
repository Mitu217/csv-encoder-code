package csv

type CsvObj struct {
	ID    int `csv:"id"`
	Val1  int `csv:"val1"`
	Val2  int `csv:"val2"`
	Total int `csv:"total"`
}
