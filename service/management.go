package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"load-balancer/models"
	routingalgo "load-balancer/routing-algo"
	"log"
	"net/http"
)

func getAllServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(routingalgo.GetAllBackendServerAsList())
}

//TODO
func removeBackendServer(w http.ResponseWriter, r *http.Request) {

}

func addBackendServer(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Method not allowed")
	} else {
		read := r.Body
		b, err := ioutil.ReadAll(read)

		if err != nil {
			log.Println(err.Error())
		}

		var server models.ServerDeatailsRequest
		err = json.Unmarshal(b, &server)

		if err != nil {
			log.Printf("unable to unmashall the server list body with error : %s\n", err.Error())
		}

		routingalgo.AddNewBackendServer(server.Servers)

	}
	w.WriteHeader(201)
}
