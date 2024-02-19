package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "test_application_hellos",
		Help: "The total number of worlds we have said hello to",
	})
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var headers = ""
	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			headers += name + value
		}
	}
	log.Print(headers)
	opsProcessed.Inc()
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	log.Printf("Starting Server")
	http.HandleFunc("/", homeHandler)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
