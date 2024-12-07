package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

//add handlers

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) add(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; ok {
		msg := fmt.Sprintf("duplicate item: %q", item)
		http.Error(w, msg, http.StatusBadRequest) //400
		return
	}
	p, err := strconv.ParseFloat(price, 32)

	if err != nil {
		msg := fmt.Sprintf("invalid price: %q", price)
		http.Error(w, msg, http.StatusBadRequest) //400
		return
	}
	db[item] = dollars(p)
	fmt.Fprintf(w, "added %s with price %q\n", item, db[item])

}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("item missing: %q\n", item)
		http.Error(w, msg, http.StatusNotFound) //404
		return
	}
	p, err := strconv.ParseFloat(price, 32)

	if err != nil {
		msg := fmt.Sprintf("invalid price: %q\n", price)
		http.Error(w, msg, http.StatusBadRequest) //400
		return
	}
	db[item] = dollars(p)
	fmt.Fprintf(w, "updated %s with price %q\n", item, db[item])

}

func (db database) drop(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("item missing: %q", item)
		http.Error(w, msg, http.StatusNotFound) //404
		return
	}

	delete(db, item)
	fmt.Fprintf(w, "removed %s from db \n", item)

}

func main() {
	db := database{
		"shoes": 50,
		"socks": 5,
	}

	// add routes
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/create", db.add)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/drop", db.drop)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
