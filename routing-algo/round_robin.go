package routingalgo

import (
	"fmt"
	"load-balancer/models"
	"log"
	"net/http/httputil"
	"net/url"
)

type roundrobin struct {
}

var currentServer int = 0

func GetRoundRobin() *roundrobin {
	return &roundrobin{}
}

func (a *roundrobin) InitServers(arrOfServers []string) {

	log.Printf("intilazing server with roundrobin")
	var lb []*models.Backend

	for _, element := range arrOfServers {
		route, _ := url.Parse(element)

		backend := &models.Backend{
			Proxy: httputil.NewSingleHostReverseProxy(route),
		}

		lb = append(lb, backend)

	}

	numberOfBackendServers = len(lb)
	listOfBackend = models.BackendList{BL: lb}
}

func (a *roundrobin) GetServer() models.Backend {

	currentServer++

	for !listOfBackend.BL[currentServer].Healthy {
		currentServer++
	}

	if currentServer == numberOfBackendServers {
		currentServer = 0
	}
	fmt.Println(currentServer)
	return *listOfBackend.BL[currentServer]

}
