package csv

import (
	"io"

	"github.com/gocarina/gocsv"
)

// DecodeFunc ...
type DecodeFunc func(objs *[]*CsvObj) error

// Decoder ...
type Decoder struct {
	r  io.Reader
	fs []DecodeFunc
}

// NewDecoder ...
func NewDecoder(r io.Reader, fs ...DecodeFunc) *Decoder {
	return &Decoder{r: r, fs: fs}
}

// Decode ...
func (dec *Decoder) Decode(out *[]*CsvObj) error {
	if err := gocsv.Unmarshal(dec.r, out); err != nil {
		return err
	}
	for _, f := range dec.fs {
		if err := f(out); err != nil {
			return err
		}
	}
	return nil
}
