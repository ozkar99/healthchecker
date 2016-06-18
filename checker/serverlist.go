package checker

import (
	"sync"
	"net/http"
	"errors"
)

type ServerList struct {
	Servers []Server
	sync.WaitGroup
	comm chan Server
}

func (sl *ServerList) Init() {
	sl.comm = make(chan Server)
	sl.Add(len(sl.Servers)) //initalize embedded waitgroup
} 

func (serverList ServerList) Failed() []Server {
	failedServerList := make([]Server, 0)
	
	for _, server := range serverList.Servers {
		/* Fan In */
		go func(s Server) {
			/* Test if site is up, send failed down the channel */
			resp, err := http.Get(s.SchemaDomain())
			if err != nil || resp.StatusCode != http.StatusOK {
				if err != nil {
					s.Error = err
				} else {
					s.Error = errors.New("Response Status is not 200 OK")
				}
				serverList.comm <- s
			}
			serverList.Done()
		}(server)
	}

	/* Fan Out */
	go func() {
		/*Append failed servers*/
		for {
			s := <-serverList.comm
			failedServerList = append(failedServerList, s)
		}
	}()

	serverList.Wait()
	return failedServerList
}
