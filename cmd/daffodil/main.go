package main

import (
	"fmt"
	"net/http"

	"github.com/tpaschalis/daffodil"
)

var cfg *daffodil.Config
var df *daffodil.Daffodil

func main() {
	cfg, _ = daffodil.NewConfig()
	df, _ = daffodil.NewDaffodil(cfg)

	http.HandleFunc("/", daffodilHandler)
	http.ListenAndServe(":8080", nil)

}

func daffodilHandler(w http.ResponseWriter, r *http.Request) {
	id, err := df.Next()
	if err != nil {
		w.WriteHeader(http.StatusLocked)
		_, err := fmt.Fprint(w, err.Error())
		if err != nil {
			fmt.Fprint(w, "Fprint returned error in Write to ResponseWriter : %s", err.Error())
		}
	}
	fmt.Fprintf(w, "%v", id)
}
