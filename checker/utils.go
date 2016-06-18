package checker

import (
	"bufio"
	"encoding/csv"
	"os"
	"fmt"
)

func FromCSVFile(file *os.File) (ServerList, error) {
	var servers ServerList

	r := csv.NewReader(bufio.NewReader(file))

	records, err := r.ReadAll()

	if err != nil {
		return servers, err
	}

	for i, record := range records {
		if i != 0 { // Ignore header
			servers.Servers = append(servers.Servers, Server{Name: record[0], Domain: record[1]})
		}
	}

	servers.Init()
	return servers, nil
}

func PrintServers(sl []Server) {
	for _, server := range sl {
		fmt.Println(server.String())
	}
}