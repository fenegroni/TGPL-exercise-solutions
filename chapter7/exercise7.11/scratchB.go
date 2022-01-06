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

func (db database) list(response http.ResponseWriter, request *http.Request) {
	for item, price := range db {
		fmt.Fprintf(response, "%s: %s\n", item, price)
	}
}

func (db database) price(response http.ResponseWriter, request *http.Request) {
	item := request.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		http.Error(response, fmt.Sprintf("no such item: %q\n", item), http.StatusNotFound)
		return
	}
	fmt.Fprintf(response, "%s\n", price)
}

func main() {
	db := database{"shoes": 50.00, "socks": 5.00}
	http.DefaultServeMux.HandleFunc("/list", db.list)
	http.DefaultServeMux.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
