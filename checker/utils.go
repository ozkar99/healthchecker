package checker

import (
	"bufio"
	"encoding/csv"
	"os"
)

func FromCSVFile(file *os.File) (ServerList, error) {
	var servers ServerList

	r := csv.NewReader(bufio.NewReader(file))

	records, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	for i, record := range records {
		if i != 0 { // Ignore header
			servers = append(servers, Server{Name: record[0], Domain: record[1]})
		}
	}

	return servers, nil
}