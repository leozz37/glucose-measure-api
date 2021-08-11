package middlewares

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func DownloadFile(fileName, fileDownloadURL string) error {
	out, err := os.Create(fileName + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(fileDownloadURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func UnzipFile(fileName string) {
	gzipfile, err := os.Open(fileName + ".gz")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reader, err := gzip.NewReader(gzipfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer reader.Close()

	newFileName := strings.TrimSuffix(fileName, ".gz")
	writer, err := os.Create(newFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer writer.Close()

	if _, err = io.Copy(writer, reader); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetLastMeasureFromCSV(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := csv.NewReader(f)
	row, err := r.ReadAll()
	if err != nil {
		return "", err
	}

	glucoseLevel := strings.Split(row[0][0], ";")[6]
	glucoseLevel = strings.Replace(glucoseLevel, "\u0000", "", -1)
	return glucoseLevel, nil
}
