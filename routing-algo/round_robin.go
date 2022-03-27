package routingalgo

import (
	"errors"
	"fmt"
	"load-balancer/health"
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

		health.ServerToBeChecked <- element
		backend := &models.Backend{
			Url:     element,
			Healthy: <-health.ServerHealthResult,
			Proxy:   httputil.NewSingleHostReverseProxy(route),
		}

		lb = append(lb, backend)

	}

	numberOfBackendServers = len(lb)
	listOfBackend = models.BackendList{BL: lb}
}

func (a *roundrobin) GetServer() (models.Backend, error) {

	currentServer++
	var numberOfDownServer int = 0

	if currentServer >= numberOfBackendServers {
		currentServer = 0
	}

	for !listOfBackend.BL[currentServer].Healthy {
		numberOfDownServer++
		currentServer++
		if currentServer >= numberOfBackendServers {
			currentServer = 0
		}
		if numberOfDownServer == numberOfBackendServers {
			log.Println("No Backend Server to serve the request")
			return *listOfBackend.BL[0], errors.New("no server up to serve")
		}
	}

	fmt.Println(currentServer)
	return *listOfBackend.BL[currentServer], nil

}
