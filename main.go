package main

import (
	routingalgo "load-balancer/routing-algo"
	"load-balancer/service"
)

// var backend_list models.BackendList

var a routingalgo.Algo

func main() {

	service.Init()
}
