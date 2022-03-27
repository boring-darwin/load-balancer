package health

import (
	"log"
	"net/http"
)

var ServerToBeChecked = make(chan string)
var ServerHealthResult = make(chan bool)

func IsServerUp() {

	log.Println("Health check routine initalized")
	go func() {
		for {

			select {

			case server := <-ServerToBeChecked:

				resp, err := http.Get(server + "/health")

				if err != nil {
					log.Printf("unable to reach %s server with error: %s\n", server, err)
					ServerHealthResult <- false
				} else if resp.StatusCode == 200 {
					log.Printf("able to reach %s server\n", server)
					ServerHealthResult <- true
				}

			default:
				// fmt.Println("waiting")

			}

		}
	}()
}
