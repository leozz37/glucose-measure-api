package models

import (
	"encoding/csv"
	"os"
	"strings"
)

type Measure struct {
	GlucoseLevel string
	Date         string
	Status       int
}

func GetLastMeasureFromCSV(fileName string) (Measure, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return Measure{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	row, err := r.ReadAll()
	if err != nil {
		return Measure{}, err
	}

	glucoseLevel := strings.Split(row[0][0], ";")[6]
	glucoseLevel = strings.Replace(glucoseLevel, "\u0000", "", -1)

	measureDate := strings.Split(row[0][0], ";")[1]
	measureDate = strings.Replace(measureDate, "\u0000", "", -1)
	measureDate = strings.Replace(measureDate, ".", ":", -1)

	measure := Measure{GlucoseLevel: glucoseLevel, Date: measureDate}
	return measure, nil
}
