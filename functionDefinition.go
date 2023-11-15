package main

type KeyTypeDefinition struct {
	Key   string
	Type  string
	Value string
}

type FunctionDefinition interface {
	Function() int

	TryInjectParameters(params map[string]KeyTypeDefinition, result *string) bool

	DefineFunction()

	TryGetVariableType(name string, result *string) bool
}

type A_FunctionDefinition struct {
	Name               string
	ExpectedParameters map[string]bool
	Parameters         map[string]KeyTypeDefinition
	Def                FunctionDefinition
}

func (a A_FunctionDefinition) TryInjectParameters(params map[string]KeyTypeDefinition, result *string) bool {
	return false
}

func (a A_FunctionDefinition) TryGetVariableType(name string, result *string) bool {
	return false
}
