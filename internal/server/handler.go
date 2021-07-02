package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"waterjugserver/internal/solver"
)

type Input struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	i := Input{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&i)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Wrong Json format", http.StatusBadRequest)
		return
	}

	res, err := solver.Solve(i.X, i.Y, i.Z)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(res)

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(resp)
	if err != nil {
		fmt.Println(err)
	}
}
