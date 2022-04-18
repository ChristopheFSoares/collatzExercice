package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/ChristopheFSoares/collatzExercice/internal/collatzConjecture"
)

func main() {

	r := chi.NewRouter()

	r.Get("/collatz/{value}", handler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	valueString := chi.URLParam(r, "value")

	value, err := strconv.Atoi(valueString)
	if err != nil || value < 1 {
		http.Error(w, "Invalid parameter in url", http.StatusBadRequest)
		return
	}

	result, _ := collatzConjecture.Collatz(value)
	w.Write([]byte(fmt.Sprintf("Number of step to reach 1 for the number %v is %v \n", value, result)))
}
