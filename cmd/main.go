package main

import (
	"github.com/mazzoleni-gabriel/transaction-risk/risk/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/rate-risks", handler.Handler)

	http.ListenAndServe(":8090", nil)
}