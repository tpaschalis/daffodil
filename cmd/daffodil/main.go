package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/tpaschalis/daffodil"
)

var cfg *daffodil.Config
var df *daffodil.Daffodil

func main() {
	cfg, err := daffodil.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	df, err = daffodil.NewDaffodil(cfg)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", daffodilHandler)
	http.HandleFunc("/dismantle", dismantleHandler)
	http.ListenAndServe(":8080", nil)
}

func daffodilHandler(w http.ResponseWriter, r *http.Request) {
	id, err := df.Next()
	if err != nil {
		w.WriteHeader(http.StatusLocked)
		_, err := fmt.Fprint(w, err.Error())
		if err != nil {
			fmt.Fprintf(w, "fprint returned error in Write to ResponseWriter : %s", err.Error())
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%v", id)
}

func dismantleHandler(w http.ResponseWriter, r *http.Request) {
	strID := r.URL.Query().Get("id")

	if strID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "url parameter 'id' is missing")
	}

	id, err := strconv.ParseUint(strID, 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "provided 'id' could not be parsed into a Uint64 variable")
	}

	ts, node, seq := daffodil.ID(id).Dismantle()
	res, err := json.MarshalIndent(map[string]int64{
		"id":        int64(id),
		"timestamp": ts,
		"node":      node,
		"sequence":  seq,
	}, "", "    ")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "could not generate JSON response")
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(res))
}
