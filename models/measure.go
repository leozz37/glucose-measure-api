package models

import (
	"encoding/csv"
	"os"
	"strings"
)

type Measure struct {
	GlucoseLast        string
	GlucosePenultimate string
	Status             int
}

func GetLastMeasureFromCSV(fileName string) (Measure, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return Measure{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	record, _ := r.Read()
	glucoseLast := strings.Split(string(record[0]), ";")[6]
	glucoseLast = strings.Replace(glucoseLast, "\u0000", "", -1)

	record, _ = r.Read()
	glucosePenultimate := strings.Split(string(record[0]), ";")[6]
	glucosePenultimate = strings.Replace(glucosePenultimate, "\u0000", "", -1)

	measure := Measure{GlucoseLast: glucoseLast, GlucosePenultimate: glucosePenultimate}
	return measure, nil
}
