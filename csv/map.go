package csv

import (
	"encoding/csv"
	"io"
	"os"
)

type fileData struct {
	reader  *csv.Reader
	csvFile *os.File
}

func newFileData(filepath string) (*fileData, error) {
	csvfile, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(csvfile)

	return &fileData{
		reader:  reader,
		csvFile: csvfile,
	}, nil
}

func MapAllFromFilepath(filepath string) (records map[string][2]string, err error) {
	fileData, err := newFileData(filepath)
	if err != nil {
		return nil, err
	}
	defer fileData.csvFile.Close()

	records = make(map[string][2]string)
	var record []string

	for record, err = fileData.reader.Read(); err == nil; record, err = fileData.reader.Read() {
		records[record[0]] = [2]string{record[1], record[2]}
	}

	if err != nil && err != io.EOF {
		return nil, err
	}
	return records, nil
}
