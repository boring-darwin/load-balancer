package service

import (
	"encoding/json"
	routingalgo "load-balancer/routing-algo"
	"net/http"
)

func getAllServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(routingalgo.GetAllBackendServerAsList())
}

func removeBackendServer(w http.ResponseWriter, r *http.Request) {

}

func addBackendServer(w http.ResponseWriter, r *http.Request) {

}
