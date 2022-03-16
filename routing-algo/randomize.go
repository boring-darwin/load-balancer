package routingalgo

import (
	"errors"
	"load-balancer/models"
	"load-balancer/service"
	"log"
	"math/rand"
	"net/http/httputil"
	"net/url"
	"time"
)

type random struct {
}

func GetRandom() *random {

	return &random{}

}

func (a *random) InitServers(arrOfServers []string) {

	var lb []*models.Backend

	for _, element := range arrOfServers {
		route, _ := url.Parse(element)

		// service.isServerUp(element)
		backend := &models.Backend{
			Url:     element,
			Healthy: service.IsServerUp(element),
			Proxy:   httputil.NewSingleHostReverseProxy(route),
		}

		lb = append(lb, backend)

	}

	numberOfBackendServers = len(lb)
	listOfBackend = models.BackendList{BL: lb}
}

func (a *random) GetServer() (models.Backend, error) {

	rand.Seed(time.Now().UnixNano())
	num := 0 + rand.Intn(numberOfBackendServers-0)
	var numberOfDownServer int = 0
	for !listOfBackend.BL[num].Healthy {
		numberOfDownServer++
		num = 0 + rand.Intn(numberOfBackendServers-0)

		if numberOfDownServer == numberOfBackendServers {
			log.Println("No Backend server aviablabe to server the request")
			return *listOfBackend.BL[num], errors.New("no server up to serve")
		}
	}

	return *listOfBackend.BL[num], nil
}
