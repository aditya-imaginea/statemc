package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func CreateStateMachineEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var statemc StateMachine
	_ = json.NewDecoder(req.Body).Decode(&statemc)
	statemc.OID = params["oid"]
	statemc.ID = CreateRandomStateMachineId(10)
	json.NewEncoder(w).Encode(statemc)
}

func CreateRandomStateMachineId(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
