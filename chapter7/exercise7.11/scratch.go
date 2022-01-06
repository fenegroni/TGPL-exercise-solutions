package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32
type database map[string]dollars

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func (db database) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(response, "%s: %s\n", item, price)
		}
	case "/price":
		item := request.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			http.Error(response, fmt.Sprintf("no such item: %q\n", item), http.StatusNotFound)
			return
		}
		fmt.Fprintf(response, "%s\n", price)
	default:
		response.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(response, "no such page: %s\n", request.URL)
	}
}

func main() {
	db := database{"shoes": 50.00, "socks": 5.00}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
