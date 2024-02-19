package main

/**
 * inspiration from:
 *    https://github.com/brancz/prometheus-example-app
 *    https://github.com/IBM/Golang
 *    https://github.com/rgerardi/hellogo
 */
import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	worldsHelloed = prometheus.NewCounter(prometheus.CounterOpts{
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
	worldsHelloed.Inc()
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	log.Printf("Starting Server")

	r := prometheus.NewRegistry()
	r.MustRegister(worldsHelloed)

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))

	var server = &http.Server{Addr: ":8080", Handler: mux}

	log.Fatal(server.ListenAndServe())
}
