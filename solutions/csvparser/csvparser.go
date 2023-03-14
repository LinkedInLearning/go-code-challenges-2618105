package csvparser

import (
	"encoding/csv"
	"os"
	"strconv"
)

func Parse(filename string) map[int]string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	output := make(map[int]string)
	for _, r := range records {
		year, err := strconv.Atoi(r[0])
		if err != nil {
			panic(err)
		}

		if val, ok := output[year]; ok {
			output[year] = val + "," + r[1]
		} else {
			output[year] = r[1]
		}

	}

	return output
}
