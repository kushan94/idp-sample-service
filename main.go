// A small web service that shows which build is running. It exists to prove the
// platform's path from this repo to the cluster: build, scan, sign, deploy.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// The deploy workflow sets nothing here; the commit shows up in the image tag.
	host, _ := os.Hostname()

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "Hello from idp-sample-service, served by %s\n", host)
	})

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
