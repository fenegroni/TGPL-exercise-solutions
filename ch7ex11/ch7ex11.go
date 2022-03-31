package ch7ex11

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
)

type dollars float32
type database map[string]dollars

func Exercise711() {
	db := database{}
	http.DefaultServeMux.HandleFunc("/list", db.listHandler)
	http.DefaultServeMux.HandleFunc("/create", db.createHandler)
	http.DefaultServeMux.HandleFunc("/update", db.updateHandler)
	http.DefaultServeMux.HandleFunc("/delete", db.deleteHandler)
}

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

func (db database) createHandler(resp http.ResponseWriter, req *http.Request) {
	db.updateHandlerImpl(resp, req, true)
}

func (db database) updateHandler(resp http.ResponseWriter, req *http.Request) {
	db.updateHandlerImpl(resp, req, false)
}

func (db database) deleteHandler(_ http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	delete(db, item)
}

func (db database) updateHandlerImpl(resp http.ResponseWriter, req *http.Request, create bool) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	_, found := db[item]
	if !found && !create {
		http.Error(resp, fmt.Sprintf("no such item: %q\n", item), http.StatusBadRequest)
		return
	} else if found && create {
		http.Error(resp, fmt.Sprintf("item already exists: %q\n", item), http.StatusBadRequest)
		return
	}
	priceVal, err := strconv.ParseFloat(price, 32)
	if err != nil {
		http.Error(resp, fmt.Sprintf("Could not parse price: %q\n", price), http.StatusBadRequest)
		return
	}
	db[item] = dollars(priceVal)
}
