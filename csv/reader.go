package csv

import (
	"encoding/csv"
	"io"
	"os"
)

type fileReader struct {
	reader  *csv.Reader
	csvFile *os.File
}

func newfileReader(filepath string) (*fileReader, error) {
	csvfile, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(csvfile)

	return &fileReader{
		reader:  reader,
		csvFile: csvfile,
	}, nil
}

func GetRecordsFromCsvFile(filepath string) (records map[string][2]string, err error) {
	fileReader, err := newfileReader(filepath)
	if err != nil {
		return nil, err
	}
	defer fileReader.csvFile.Close()

	records = make(map[string][2]string)
	var record []string
	for record, err = fileReader.reader.Read(); err == nil; record, err = fileReader.reader.Read() {
		records[record[0]] = [2]string{record[1], record[2]}
	}

	if err != nil && err != io.EOF {
		return nil, err
	}
	return records, nil
}
