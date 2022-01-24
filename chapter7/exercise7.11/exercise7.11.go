package exercise7_11

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
)

type dollars float32
type database map[string]dollars

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func (db database) list(resp http.ResponseWriter, _ *http.Request) {
	var items []string
	for i := range db {
		items = append(items, i)
	}
	sort.Strings(items)
	for _, i := range items {
		_, _ = fmt.Fprintf(resp, "%s: %s\n", i, db[i])
	}
}

func (db database) update(resp http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	_, ok := db[item]
	if !ok {
		http.Error(resp, fmt.Sprintf("no such item: %q\n", item), http.StatusNotFound)
		return
	}
	if priceVal, err := strconv.ParseFloat(price, 32); err == nil {
		db[item] = dollars(priceVal)
	}
}

func (db database) create(resp http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	_, ok := db[item]
	if !ok {
		http.Error(resp, fmt.Sprintf("no such item: %q\n", item), http.StatusNotFound)
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
