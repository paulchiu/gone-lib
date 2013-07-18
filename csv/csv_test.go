package csv

import (
	"testing"
)

func TestReadAll(t *testing.T) {
	file := "/tmp/sample.csv"
	data, err := ReadAll(file)

	if err != nil {
		t.Error("Could not read CSV file due to: " + err.Error())
	} else {
		t.Log(data)
	}
}

// This example shows how to read all rows from a given CSV file.
func ExampleReadAll() {
	data, err := ReadAll("/Users/example/tmp/mydata.csv")

	if data == nil && err != nil {
		// Could not read file
	}
}
