package checker

import (
	"sync"
	"net/http"
	"fmt"
)

type ServerList []Server

func (servers ServerList) Failed() ServerList {
	var wg sync.WaitGroup
	failedServerList := make(ServerList, 0)
	
	ch := make(chan Server)
	wg.Add(len(servers))

	for _, server := range servers {
		/* Fan In */
		go func(c chan Server, s Server) {
			/* Test if site is up, send failed down the channel */
			resp, err := http.Get(s.SchemaDomain())
			if err != nil || resp.StatusCode != http.StatusOK {
				c <- s
			}
			wg.Done()
		}(ch, server)
	}

	/* Fan Out */
	go func(c chan Server) {
		/*Append failed servers*/
		for {
			s := <-c
			failedServerList = append(failedServerList, s)
		}
	}(ch)

	wg.Wait()
	return failedServerList
}

func (servers ServerList) Print() {
	for _, server := range servers {
		fmt.Println(server.String())
	}
}
