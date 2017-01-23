package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type StateTransition struct {
	ID          string `json:"id,omitempty"`
	SourceID    string `json:"sourceId,omitempty"`
	TargetID    string `json:"targetId,omitempty"`
	Description string `json:"description,omitempty"`
	Owner       string `json:"owner,omitempty"`
}

var stateTransitions []StateTransition

// written with perspective of single transition between states
func GetStateTransitionEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	println("passed state is " + params["sid"] + " & statetransition value is " + params["tid"])
	for _, item := range stateTransitions {
		if item.SourceID == params["sid"] && item.ID == params["tid"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&StateTransition{})
}

func GetStateTransitionsEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	println("passed state value is " + params["id"])
	var transition []StateTransition
	for _, item := range stateTransitions {
		if item.SourceID == params["sid"] {
			transition = append(transition, item)
		}
	}
	json.NewEncoder(w).Encode(transition)
}

func CreateStateTransitionEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var stateTransition StateTransition
	_ = json.NewDecoder(req.Body).Decode(&stateTransition)
	stateTransition.ID = params["tid"]
	//stateTransition.SourceID = params["sid"]

	stateTransitions = append(stateTransitions, stateTransition)
	json.NewEncoder(w).Encode(stateTransitions)
}

func DeleteStateTransitionEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range stateTransitions {
		if params["sid"] == item.SourceID && params["tid"] == item.ID {
			stateTransitions = append(stateTransitions[:index], stateTransitions[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(stateTransitions)
}
