package main

type StateMc struct {
	ProcessOwner     string `json:"processowner,omitempty"`
	States           []State
	StateTransitions []StateTransition
}
