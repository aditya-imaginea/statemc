package main

type StateMachine struct {
	ID          string `json:"smid,omitempty"`
	OID         string `json:"ownerId,omitempty"`
	States      []State
	Transitions []StateTransition
}

type State struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type StateTransition struct {
	ID          string `json:"id,omitempty"`
	SourceID    string `json:"sourceId,omitempty"`
	TargetID    string `json:"targetId,omitempty"`
	Description string `json:"description,omitempty"`
	Owner       string `json:"owner,omitempty"`
}
