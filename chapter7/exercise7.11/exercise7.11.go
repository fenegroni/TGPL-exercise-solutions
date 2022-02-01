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

func (db database) listHandler(resp http.ResponseWriter, _ *http.Request) {
	var items []string
	for i := range db {
		items = append(items, i)
	}
	sort.Strings(items)
	for _, i := range items {
		_, _ = fmt.Fprintf(resp, "%s: %s\n", i, db[i])
	}
}

func (db database) update(resp http.ResponseWriter, req *http.Request, create bool) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if !create {
		_, ok := db[item]
		if !ok {
			http.Error(resp, fmt.Sprintf("no such item: %q\n", item), http.StatusNotFound)
			return
		}
	}
	priceVal, err := strconv.ParseFloat(price, 32)
	if err != nil {
		http.Error(resp, fmt.Sprintf("Could not parse price: %q\n", price), http.StatusBadRequest)
		return
	}
	db[item] = dollars(priceVal)
}

func (db database) updateHandler(create bool) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		db.update(resp, req, create)
	}
}

func Exercise711() {
	db := database{}
	http.DefaultServeMux.HandleFunc("/list", db.listHandler)
	http.DefaultServeMux.HandleFunc("/update", db.updateHandler(false))
	http.DefaultServeMux.HandleFunc("/create", db.updateHandler(true))
}