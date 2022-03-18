package routingalgo

import (
	"load-balancer/models"
)

var listOfBackend models.BackendList
var numberOfBackendServers int

type algoFactory interface {
	InitServers(listOfServers []string)
	GetServer() (models.Backend, error)
}

type Algo struct {
	name        string
	InitServers func(listOfServers []string)
	GetServer   func() (models.Backend, error)
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

func GetAllBackendServerAsList() models.AllBackendServerResponse {

	resp := models.AllBackendServerResponse{}

	for _, r := range listOfBackend.BL {
		i := models.Info{
			Url:     r.Url,
			Healthy: r.Healthy,
		}
		resp.Servers = append(resp.Servers, i)
	}

	return resp
}
