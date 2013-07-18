// Gone-lib CSV file "one" liners. See ReadAll().
package csv

import (
	"encoding/csv"
	"os"
)

// Read all rows of a CSV file as a 2D array of strings.
func ReadAll(filename string) ([][]string, error) {
	// Open file for reading
	file, ferr := os.Open(filename)
	if ferr != nil {
		return nil, ferr
	}
	defer file.Close()

	// Read file
	reader := csv.NewReader(file)
	records, raErr := reader.ReadAll()
	if raErr != nil {
		return nil, raErr
	}

	return records, nil
}
