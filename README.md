A Simple state machine representation using go lang.
sample url for creating state machine can be found here

http://www.jsoneditoronline.org/?id=a2aaead519e9a4c02f05623e76bfaefa

steps:
-- pull the code base
-- go build
-- run statemc.exe {or sh}


curl -H "Content-Type: application/json" -X POST -d '{json from above sample url}' http://localhost:12345/statemc/{ownerid}

this will create a state machine with states and transitions.
TODO:
1. validations for valid state transitions
2. validations for mandatory fields etc...
this is still work in progress
