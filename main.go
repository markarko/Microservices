package main

import (
	//"fmt"
	//"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/markarko/Microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	h := handlers.NewHello(l)
	sm := http.NewServeMux()

	sm.HandleFunc("/", h.ServeHTTP)

	http.ListenAndServe("localhost:8080", sm)
}
