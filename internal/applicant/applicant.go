package main

import (
	"encoding/json"
	"fmt"
	"github.com/golkity/calc_go/pkg/calc"
	"net/http"
)

type Response struct {
	Result float64 `json:"result"`
	Error  string  `json:"error"`
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	expression := r.URL.Query().Get("expression")
	if expression == "" {
		http.Error(w, "Expression is required", http.StatusBadRequest)
		return
	}

	result, err := calc.Calc(expression)

	resp := Response{}
	if err != nil {
		resp.Error = err.Error()
	} else {
		resp.Result = result
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
	}
}

func application() {
	http.HandleFunc("/calc", calculateHandler)
	fmt.Print("Server is running on http://localhost8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Print(err)
	}
}
