package handler

import (
	"encoding/json"
	"github.com/mazzoleni-gabriel/transaction-risk/risk"
	"github.com/mazzoleni-gabriel/transaction-risk/risk/service"
	"net/http"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//Decode json into user
	var transactionsRequest risk.TransactionsRequest
	json.NewDecoder(req.Body).Decode(&transactionsRequest)

	response := service.RateRisks(transactionsRequest)

	// Decode user into json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	b, _ := json.Marshal(response)
	_, _ = w.Write(b)
}
