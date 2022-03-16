package main

import (
	"fmt"
	routingalgo "load-balancer/routing-algo"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// var backend_list models.BackendList

var a routingalgo.Algo

func main() {
	fmt.Println("LB Started....")
	rand.Seed(time.Now().UnixNano())

	a = *routingalgo.GetAlgo("roundrobin")
	servers := []string{"http://localhost:8083", "http://localhost:8082"}

	a.InitServers(servers)

	http.HandleFunc("/", test)

	http.ListenAndServe(":8080", nil)
}

func test(w http.ResponseWriter, r *http.Request) {

	t, err := a.GetServer()

	if err != nil {
		log.Println("unable to forward the call")
		w.WriteHeader(500)
		fmt.Fprintf(w, "no backend server up to serve the request")

	} else {
		t.Proxy.ServeHTTP(w, r)
	}

}
