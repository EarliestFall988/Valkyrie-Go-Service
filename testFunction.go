package main

func NewFunction() *A_FunctionDefinition {

	name := "TestFunction"
	expectedParameters := map[string]ValueTypeDefinition{"a": {Type: Int, Applied: false}, "b": {Type: Int, Applied: false}}
	parameters := map[string]KeyTypeDefinition{}

	result := A_FunctionDefinition{Name: name, ExpectedParameters: expectedParameters, Parameters: parameters}

	return &result
}

func (a A_FunctionDefinition) Function() int {
	return 0
}
