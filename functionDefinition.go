package main

import "strconv"

type TypeDef int64

const (
	Text    TypeDef = 0
	Int     TypeDef = 1
	Decimal TypeDef = 2
	Bool    TypeDef = 3
)

type KeyTypeDefinition struct {
	Key   string
	Type  TypeDef
	Value string
}

type ValueTypeDefinition struct {
	Type    TypeDef
	Applied bool
}

type FunctionDefinition interface {
	Function() int

	TryInjectParameters(params map[string]KeyTypeDefinition, result *string) bool

	DefineFunction()

	TryGetVariableType(name string, result *string) bool
}

type A_FunctionDefinition struct {
	Name               string
	ExpectedParameters map[string]ValueTypeDefinition
	Parameters         map[string]KeyTypeDefinition
	Def                FunctionDefinition
}

func (a A_FunctionDefinition) TryInjectParameters(params map[string]KeyTypeDefinition, r *string) bool {

	if len(params) != len(a.ExpectedParameters) {

		lenExp := len(a.ExpectedParameters)
		lenAct := len(params)

		var lenE int64 = 1
		var lenA int64 = 1

		s := strconv.FormatInt(lenE, lenExp)
		rr := strconv.FormatInt(lenA, lenAct)

		*r = "Parameter count mismatch. The expected parameter count is " + s + " but the actual parameter count is " + rr
		return false
	}

	for k, v := range params {

		val, ok := a.ExpectedParameters[k]

		if !ok {

			*r = "Parameter mismatch. The expected parameter " + k + " does not exist"
			return false
		}

		if val.Type == v.Type {
			a.Parameters[k] = v
		} else {

			*r = "Parameter type mismatch. The expected type of param " + k + " is " + string(a.ExpectedParameters[k].Type) + " but the actual type of param " + k + " is " + string(v.Type)
			return false
		}
	}

	successResult := "Parameters injected successfully"
	r = &successResult
	return true
}

func (a A_FunctionDefinition) TryGetVariableType(name string, result *TypeDef) bool {
	val, ok := a.ExpectedParameters[name]

	if ok {
		*result = val.Type
		return true
	} else {
		*result = 0
		return false
	}
}
