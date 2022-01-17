package exercise7_11

import (
	"fmt"
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

func Exercise7_11() {
	db := database{"shoes": 50.00, "socks": 5.00}
	http.DefaultServeMux.HandleFunc("/list", db.list)
}
