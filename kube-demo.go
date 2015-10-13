package main // import "github.com/ContainerSolutions/kubernetes-demo"

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "?"
		}

		fmt.Fprintf(w, hostname)
	})

	log.Print("Ready...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
