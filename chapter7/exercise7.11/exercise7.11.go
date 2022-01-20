package exercise7_11

import (
	"fmt"
	"net/http"
	"strconv"
)

type dollars float32
type database map[string]dollars

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func (db database) list(response http.ResponseWriter, _ *http.Request) {
	for item, price := range db {
		_, _ = fmt.Fprintf(response, "%s: %s\n", item, price)
	}
}

func (db database) update(response http.ResponseWriter, request *http.Request) {
	item := request.URL.Query().Get("item")
	price := request.URL.Query().Get("price")
	_, ok := db[item]
	if !ok {
		http.Error(response, fmt.Sprintf("no such item: %q\n", item), http.StatusNotFound)
		return
	}
	if priceVal, err := strconv.ParseFloat(price, 32); err == nil {
		db[item] = dollars(priceVal)
	}
}

func Exercise711() {
	db := database{"shoes": 50.00, "socks": 5.00}
	http.DefaultServeMux.HandleFunc("/list", db.list)
	http.DefaultServeMux.HandleFunc("/update", db.update)
}
