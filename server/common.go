package server

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Authorization")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong")
}

func InitServer() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/makeToken", randomNumberHandler)


	log.Fatal(http.ListenAndServe(":8081", nil))
}

func randomNumberHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	seed := getURLParam(r, "seed")
	address := getURLParam(r, "address")
	token := getURLParam(r, "token")

	result := address + seed + token + strconv.Itoa(rand.Intn(10000000))
	fmt.Fprintf(w,result)
}

func getURLParam(r *http.Request, paramName string) string {
	keys, seen := r.URL.Query()[paramName]

	if seen && len(keys) > 0 {
		return keys[0]
	}
	return ""
}