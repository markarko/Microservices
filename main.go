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

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()

	sm.HandleFunc("/hello", hh.ServeHTTP)
	sm.HandleFunc("/goodbye", gh.ServeHTTP)

	http.ListenAndServe("localhost:9090", sm)
}
