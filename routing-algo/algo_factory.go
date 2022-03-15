package routingalgo

import (
	"load-balancer/models"
)

var listOfBackend models.BackendList
var numberOfBackendServers int

type algoFactory interface {
	InitServers(listOfServers []string)
	GetServer() models.Backend
}

type Algo struct {
	name        string
	InitServers func(listOfServers []string)
	GetServer   func() models.Backend
}

func GetAlgo(algo_name string) *Algo {

	switch algo_name {

	case "random":
		r := GetRandom()
		return &Algo{
			name:        "random",
			InitServers: r.InitServers,
			GetServer:   r.GetServer,
		}

	case "roundrobin":
		r := GetRoundRobin()
		return &Algo{
			name:        "roundrobin",
			InitServers: r.InitServers,
			GetServer:   r.GetServer,
		}

	default:
		r := GetRandom()
		return &Algo{
			name:        "random",
			InitServers: r.InitServers,
			GetServer:   r.GetServer,
		}
	}
}
