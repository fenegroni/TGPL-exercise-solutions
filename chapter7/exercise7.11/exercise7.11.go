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

func (db database) updateHandlerImpl(resp http.ResponseWriter, req *http.Request, create bool) {
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

func (db database) createHandler(resp http.ResponseWriter, req *http.Request) {
	db.updateHandlerImpl(resp, req, true)
}

func (db database) updateHandler(resp http.ResponseWriter, req *http.Request) {
	db.updateHandlerImpl(resp, req, false)
}

func (db database) deleteHandler(resp http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	delete(db, item)
}

func Exercise711() {
	db := database{}
	http.DefaultServeMux.HandleFunc("/list", db.listHandler)
	http.DefaultServeMux.HandleFunc("/create", db.createHandler)
	http.DefaultServeMux.HandleFunc("/update", db.updateHandler)
	http.DefaultServeMux.HandleFunc("/delete", db.deleteHandler)
}
