package service

import (
	"fmt"
	routingalgo "load-balancer/routing-algo"
	"math/rand"
	"net/http"
	"time"
)

var a routingalgo.Algo

func Init() {

	fmt.Println("LB Started....")
	rand.Seed(time.Now().UnixNano())

	a = *routingalgo.GetAlgo("roundrobin")
	servers := []string{"http://localhost:8083", "http://localhost:8082"}

	a.InitServers(servers)

	http.HandleFunc("/", forwardProxy)

	http.ListenAndServe(":8080", nil)
}
