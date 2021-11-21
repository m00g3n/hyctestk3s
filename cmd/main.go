package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var healtz = func(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "OK")
	w.WriteHeader(http.StatusOK)
}

var d6 = func(rnd *rand.Rand) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		result := rnd.Intn(5) + 1
		fmt.Fprintf(w, "%d", result)
		w.WriteHeader(http.StatusOK)
	}
}

func main() {
	// initialize random source
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	http.Handle("/d6", d6(rnd))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
