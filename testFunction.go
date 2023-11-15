package main

func TestFunction() *A_FunctionDefinition {

	name := "TestFunction"
	expectedParameters := map[string]bool{"a": true, "b": true}
	parameters := map[string]KeyTypeDefinition{"a": {Key: "a", Type: "int", Value: ""}, "b": {Key: "b", Type: "int", Value: ""}}

	result := A_FunctionDefinition{Name: name, ExpectedParameters: expectedParameters, Parameters: parameters}

	return &result
}

func (a A_FunctionDefinition) Function() int {
	return 0
}
