package service

import (
	"fmt"
	"log"
	"net/http"
)

func forwardProxy(w http.ResponseWriter, r *http.Request) {

	u := r.URL

	if u.Path == "/server/all" {
		getAllServer(w, r)
		return
	}

	t, err := a.GetServer()

	if err != nil {
		log.Println("unable to forward the call")
		w.WriteHeader(500)
		fmt.Fprintf(w, "no backend server available to serve the request")

	} else {
		t.Proxy.ServeHTTP(w, r)
	}

}
