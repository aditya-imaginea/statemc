package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var states []State

func GetStateEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range states {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&State{})
}

func GetStatesEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(states)
}

func CreateStateEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var state State
	_ = json.NewDecoder(req.Body).Decode(&state)
	state.ID = params["id"]

	states = append(states, state)
	json.NewEncoder(w).Encode(states)
}

func DeleteStateEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range states {
		if params["id"] == item.ID {
			states = append(states[:index], states[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(states)
}

func main() {
	states = append(states, State{ID: "1", Name: "compliance", Description: " state used to describe compliance"})
	states = append(states, State{ID: "2", Name: "governance", Description: " state used to describe governance"})
	states = append(states, State{ID: "3", Name: "risk"})
	states = append(states, State{ID: "4", Name: "escalation"})

	router := mux.NewRouter()
	//Routes for states
	router.HandleFunc("/state", GetStatesEndpoint).Methods("GET")
	router.HandleFunc("/state/{id}", GetStateEndpoint).Methods("GET")
	router.HandleFunc("/state/{id}", CreateStateEndpoint).Methods("POST")
	router.HandleFunc("/state/{id}", DeleteStateEndpoint).Methods("DELETE")
	//Routes for state transitions
	router.HandleFunc("/state/{sid}/transition", GetStateTransitionsEndpoint).Methods("GET")
	router.HandleFunc("/state/{sid}/transition/{tid}", GetStateTransitionEndpoint).Methods("GET")
	router.HandleFunc("/state/{sid}/transition/{tid}", CreateStateTransitionEndpoint).Methods("POST")
	router.HandleFunc("/state/{sid}/transition/{tid}", DeleteStateTransitionEndpoint).Methods("DELETE")
	// Routes for creating statemachine
	//	router.HandleFunc("/statemc", GetStatesMcEndpoint).Methods("GET")
	//router.HandleFunc("/statemc/{smid}", GetStateMcEndpoint).Methods("GET")
	//router.HandleFunc("/statemc/{smid}/states", AppendStatesEndpoint).Methods("POST")
	//router.HandleFunc("/statemc/{smid}/states", GetStatesInSmEndpoint).Methods("GET")
	//router.HandleFunc("/statemc/{smid}/owner", GetStateMcOwner).Methods("GET")
	//router.HandleFunc("/statemc/owner", GetAllStateMcOwners).Methods("GET")
	router.HandleFunc("/statemc/{oid}", CreateStateMachineEndpoint).Methods("POST")
	//router.HandleFunc("/statemc/{smid}/transitions", AppendTransitionsEndpoint).Methods("POST")
	//router.HandleFunc("/statemc/{smid}/transitions", GetTransitionsinSmEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))

}
