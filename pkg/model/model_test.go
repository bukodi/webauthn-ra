package model

import (
	"encoding/json"
	"testing"
)

type Versions struct {
	Hash string
}

type FooRecord struct {
	Field1       string
	Field2       string
	HiddenField3 string
}

type FooDTO struct {
	Hash string
	FooRecord
	AddedField4  string
	HiddenField3 string
}

func TestCica(t *testing.T) {

}

func TestMarshall(t *testing.T) {

	rec := &FooRecord{
		Field1:       "value1",
		Field2:       "value2",
		HiddenField3: "value3",
	}

	{
		b, _ := json.MarshalIndent(rec, "", "  ")
		t.Logf("FooRecord: %s\n", string(b))
	}

	dto := &FooDTO{
		FooRecord:   *rec,
		Hash:        "hashValue",
		AddedField4: "value4",
	}

	{
		b, _ := json.MarshalIndent(dto, "", "  ")
		t.Logf("FooDTO: %s\n", string(b))
	}

	{
		inlineDTO := &struct {
			FooRecord
			Value1X string
			Value2X string
		}{
			FooRecord: *rec,
			Value1X:   "xvalue1",
			Value2X:   "xvalue2",
		}
		b, _ := json.MarshalIndent(inlineDTO, "", "  ")
		t.Logf("FooDTO: %s\n", string(b))
	}

}
