package service

import (
	"log"
	"net/http"
)

func IsServerUp(serverUrl string) bool {

	resp, err := http.Get(serverUrl + "/health")

	if err != nil {
		log.Printf("unable to reach %s server with error: %s\n", serverUrl, err)
		return false
	}

	if resp.StatusCode == 200 {
		return true
	}

	return false
}
