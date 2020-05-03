package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tpaschalis/daffodil"
)

func TestDaffodilHandler(t *testing.T) {
	cfg, err := daffodil.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	df, err = daffodil.NewDaffodil(cfg)
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)

	http.HandlerFunc(daffodilHandler).ServeHTTP(rr, req)
	assert.Equalf(t, http.StatusOK, rr.Code, "expected %d status code, got %d instead", http.StatusOK, rr.Code)

	assert.NotEmpty(t, rr.Body.String())
	strID := rr.Body.String()
	id, err := strconv.ParseUint(strID, 10, 64)
	assert.Nil(t, err)
	assert.NotEmpty(t, id)
}

func TestDismantleHandler(t *testing.T) {
	cfg, err := daffodil.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	df, err = daffodil.NewDaffodil(cfg)
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/dismantle?id=17801782472864612", nil)

	http.HandlerFunc(dismantleHandler).ServeHTTP(rr, req)
	assert.Equalf(t, http.StatusOK, rr.Code, "expected %d status code, got %d instead", http.StatusOK, rr.Code)

	response := rr.Body.String()

	assert.NotEmpty(t, response)

	expected := string(`{"id": 17801782472864612, "node": 43507, "sequence": 100, "timestamp": 1061068920}`)
	assert.JSONEq(t, expected, response, "dismantleHandler provided incorrect response")
}
