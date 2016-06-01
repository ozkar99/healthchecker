package main

import (
	"fmt"
	"log"
	"os"

	"healthcheck/checker"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Needs the file as argument")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	servers, err := checker.FromCSVFile(file)
	if err != nil {
		log.Fatal(err)
	}

	failedHosts := servers.Failed()

	fmt.Println("Failed Hosts:")
	failedHosts.Print()
}
