package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", law)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func law(w http.ResponseWriter, r *http.Request) {
	counter := 0
	random := rand.Intn(5) + 1
	endDate := time.Now().Add(time.Duration(random) * time.Second).Unix()
	for time.Now().Unix() < endDate {
		counter += 10
	}
	fmt.Fprintln(w, "Don’t worry if it doesn’t work right. If everything did, you’d be out of a job. - (Mosher’s Law of Software Engineering)")
}
