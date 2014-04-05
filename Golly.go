package Golly

import (
	"Golly/parser"
	"fmt"
	"strconv"
)

type FunctionObj struct {
	Type   string
	Parems []string
	Body   []ListCell
	Pure   bool
	GoFunc bool
	FuncType goFuncType
}

func (aFunc *FunctionObj) Call(funcType goFuncType, params []*ListCell)([]*ListCell, error) {
	if aFunc.GoFunc{
		returnedVals, err := CallGoFunc(funcType,  params)
		if err != nil{
			return nil, err
		}else{
			return returnedVals, nil
		}
	}else{
		return nil, nil	
	}
}

func makeSysFunc(funcType goFuncType)ListCell{
	return ListCell{TypeName: "Function", 
		Value: FunctionObj{Type: "GoFunc", Pure: true, GoFunc: true, FuncType: funcType},
		Mutable: false}
}

func CreateSystemFuncs() *SysEnvironment{
	sysBindings := make(map[string]EnvBinding)
	sysBindings["if"] = EnvBinding{Binding: makeSysFunc(GoIfT)}
	sysBindings["+"] = EnvBinding{Binding: makeSysFunc(GoAddT)}
	sysBindings["-"] = EnvBinding{Binding: makeSysFunc(GoSubtractT)}
	sysBindings["*"] = EnvBinding{Binding: makeSysFunc(GoMultiplyT)}
	sysBindings["/"] = EnvBinding{Binding: makeSysFunc(GoDivideT)}
	env := SysEnvironment{Bindings: sysBindings}
	return &env
}

type singleType struct {
	Inputs  []string
	Outputs []string
}

type TypeObj struct {
	Types []singleType
}

func (firstType *TypeObj) EqualTo(secondType *TypeObj) bool {
	return false
}

func TypesEqualP(first *ListCell, second *ListCell, lineNum int, caller *string) bool {
	if firstType, ok := first.Value.(TypeObj); ok {
		if secondType, ok := second.Value.(TypeObj); ok {
			return firstType.EqualTo(&secondType)
		} else {
			errMsg := fmt.Sprintf("Error: cell claiming to be a type actually contains something else, in %v at line %v.\n", caller, lineNum)
			panic(errMsg)
		}
	} else {
		errMsg := fmt.Sprintf("Error: cell claiming to be a type actually contains something else, in %v at line %v.\n", caller, lineNum)
		panic(errMsg)
	}
}

type EnvBinding struct {
	Name    string
	Binding ListCell
}

type SysEnvironment struct {
	Bindings map[string]EnvBinding
}

type Environment struct {
	Bindings []EnvBinding
	Parent   *Environment
	System   *SysEnvironment
}

func (env Environment) findBinding(name string, recur, checkSystem bool) *EnvBinding {
	if checkSystem {
		if binding, ok := env.System.Bindings[name]; ok {
			return &binding
		}
	}
	for i, binding := range env.Bindings {
		if binding.Name == name {
			return &env.Bindings[i]
		}
	}
	if recur {
		if env.Parent == nil {
			return nil
		} else {
			return env.Parent.findBinding(name, true, false)
		}
	}
	return nil
}

func (env *Environment) addBinding(recur bool) *EnvBinding {
	if recur {
		if env.Parent == nil {
			env.Bindings = append(env.Bindings, EnvBinding{})
			return &((env.Bindings)[len(env.Bindings)])
		} else {
			return env.Parent.addBinding(true)
		}
	} else {
		env.Bindings = append(env.Bindings, EnvBinding{})
		return &((env.Bindings)[len(env.Bindings)])
	}
}

type ListCell struct {
	TypeName string
	Value    interface{}
	Mutable  bool
}

type CellList struct {
	Cells []ListCell
	Env   Environment
}

func evalLitToken(num *Parser.Token, lineNum int, caller *string) ListCell {
	newValue := ListCell{}
	switch (*num).LitType {
	case Parser.FloNum:
		floatval, err := strconv.ParseFloat((*num).Value, 32)
		if err != nil {
			errMsg := fmt.Sprintf("Error: cannot parse string %v to float in %v at line %v.\n", (*num).Value, caller, lineNum)
			panic(errMsg)
		} else {
			newValue.Value = floatval
			newValue.TypeName = "float"
		}
	case Parser.FixNum:
		intval, err := strconv.Atoi((*num).Value)
		if err != nil {
			errMsg := fmt.Sprintf("Error: cannot parse string %v to int in %v at line %v.\n", (*num).Value, caller, lineNum)
			panic(errMsg)
		} else {
			newValue.Value = intval
			newValue.TypeName = "int"
		}
	default:
		errMsg := fmt.Sprintf("Error: unhandled literal type for %v in %v at line %v.\n", (*num).Value, caller, lineNum)
		panic(errMsg)
	}
	return newValue
}

func evalIdToken(identifierName *Parser.Token, env *Environment, lineNum int, caller *string) interface{} {
	var newValue interface{}
	valueReferenced := env.findBinding((*identifierName).Value, true, true)
	if valueReferenced == nil {
		errMsg := fmt.Sprintf("Error: attempting to evalute var %v in %v at line %v, but that var is unbound.\n", (*identifierName).Value, caller, lineNum)
		panic(errMsg)
	} else {
		newValue = valueReferenced.Binding
	}
	return newValue
}

func parseType(identifierToBindTo, potentialType *Parser.Token, env *Environment, lineNum int, caller *string) (*ListCell, string) {
	newTypeName := ""
	typeLiteralFound := false
	var newType *ListCell
	switch (*potentialType).Type {
	case Parser.LiteralToken:
		errMsg := fmt.Sprintf("Error: attempting use a numeric literal as the type for %v in %v at line %v.\n", identifierToBindTo.Value, caller, lineNum)
		panic(errMsg)
	case Parser.DefToken:
		errMsg := fmt.Sprintf("Error: attempting use a reserved name as the type for %v in %v at line %v.\n", identifierToBindTo.Value, caller, lineNum)
		panic(errMsg)
	case Parser.IdToken:
		potentialNewTypeValue := env.findBinding(potentialType.Value, true, true)
		if potentialNewTypeValue != nil {
			if potentialNewTypeValue.Binding.TypeName != "type" {
				errMsg := fmt.Sprintf("Error: attempting to assign something that is not a type, but a %v, to %v in %v at line %v.\n", potentialNewTypeValue.Binding.TypeName, identifierToBindTo.Value, caller, lineNum)
				panic(errMsg)
			} else {
				newTypeName = potentialType.Value
				newType = &potentialNewTypeValue.Binding
			}
			//Handle var containing quoted type here
		} else {
			errMsg := fmt.Sprintf("Error: attempting to assign identifier %v to %v in %v at line %v, but %v is unbound.\n", (*potentialType).Value, (*identifierToBindTo).Value, caller, lineNum, (*potentialType).Value)
			panic(errMsg)
		}
	case Parser.ListToken:
		potentialNewType := evalListToken(potentialType)
		if newType.TypeName != "type" {
			errMsg := fmt.Sprintf("Error: attempting to assign something that is not a type, but a %v, to %v in %v at line %v.\n", newType.TypeName, (*identifierToBindTo).Value, caller, lineNum)
			panic(errMsg)
		} else {
			typeLiteralFound = true
			newType = &potentialNewType
		}
		//Handle function returning quoted type here
	case Parser.TypeAnnToken:
		errMsg := fmt.Sprintf("Error: misplaced type annotation marker in %v at line %v.\n", caller, lineNum)
		panic(errMsg)
	default:
		errMsg := fmt.Sprintf("Error: unhandled token type in %v at line %v.\n", caller, lineNum)
		panic(errMsg)
	}
	if typeLiteralFound {
		if _, ok := newType.Value.(TypeObj); !ok {
			errMsg := fmt.Sprintf("Error: cell claiming to be a type actually contains something else, in %v at line %v.\n", caller, lineNum)
			panic(errMsg)
		}
	}
	return newType, newTypeName
}

func parseNewIdentifier(val *Parser.Token, env *Environment, global bool, lineNum int, caller *string) *EnvBinding {
	if val.Type != Parser.IdToken {
		errMsg := fmt.Sprintf("Error: attempting to assign to a non-identifier in %v at line %v.\n", caller, lineNum)
		panic(errMsg)
	}
	prevBinding := env.findBinding(val.Value, false, true)
	var newBinding *EnvBinding
	if prevBinding != nil {
		if !(*prevBinding).Binding.Mutable {
			errMsg := fmt.Sprintf("Error: attempting to assign to an immutable identifier in %v at line %v.\n", caller, lineNum)
			panic(errMsg)
		} else {
			newBinding = prevBinding
		}
	} else {
		newBinding = env.addBinding(global)
	}
	return newBinding
}

func parseIdentifierToBeBound(identifierToBeBoundTo, identifierToBind *Parser.Token, env *Environment, global bool, lineNum int, caller *string) *ListCell {
	newValue := ListCell{TypeName: "undecided"}
	switch (*identifierToBind).Type {
	case Parser.LiteralToken:
		newValue = evalLitToken(identifierToBind, lineNum, caller)
	case Parser.DefToken:
		errMsg := fmt.Sprintf("Error: attempting to assign reserved name %v to %v in %v at line %v.\n", identifierToBind.Value, identifierToBeBoundTo.Value, caller, lineNum)
		panic(errMsg)
	case Parser.IdToken:
		potentialNewValue := env.findBinding(identifierToBind.Value, true, true)
		if potentialNewValue != nil {
			newValue.Value = potentialNewValue
		} else {
			errMsg := fmt.Sprintf("Error: attempting to assign identifier %v to %v in %v at line %v, but %v is unbound.\n", identifierToBind.Value, identifierToBeBoundTo.Value, caller, lineNum, identifierToBind.Value)
			panic(errMsg)
		}
	case Parser.ListToken:
		newValue.Value = evalListToken(identifierToBind)
	case Parser.TypeAnnToken:
		errMsg := fmt.Sprintf("Error: expected identifier to %v in %v at line %v, but got type annotation token \":\".\n", identifierToBeBoundTo.Value, caller, lineNum)
		panic(errMsg)
	default:
		errMsg := fmt.Sprintf("Error: unhandled token type for %v in %v at line %v.\n", identifierToBind.Value, caller, lineNum)
		panic(errMsg)
	}
	return &newValue
}

func bindVars(list *Parser.Token, env Environment, lineNum int, global, mut bool, caller *string) Environment {
	for i := 0; i < len(list.ListVals); i++ {
		howManyIndicesToJumpForward := 1
		isTypeAnnotated := false
		firstListItem := &list.ListVals[i]
		newBinding := parseNewIdentifier(firstListItem, &env, global, lineNum, caller)
		newBinding.Name = firstListItem.Value
		if i >= len(list.ListVals) {
			errMsg := fmt.Sprintf("Error: nothing to assign to %v in %v at line %v.\n", firstListItem.Value, caller, lineNum)
			panic(errMsg)
		}
		var potentialNewValue *ListCell
		annotatedTypeName := ""
		var annotatedTypeValue *ListCell
		nextListItem := &list.ListVals[i+1]
		if nextListItem.Type == Parser.TypeAnnToken {
			if i+2 >= len(list.ListVals) {
				errMsg := fmt.Sprintf("Error: no type and/or value provided in assignment to %v in %v at line %v.\n", firstListItem.Value, caller, lineNum)
				panic(errMsg)
			} else {
				potentialTypeItem := &list.ListVals[i+2]
				annotatedTypeValue, annotatedTypeName = parseType(firstListItem, potentialTypeItem, &env, lineNum, caller)
				potentialNewValueItem := &list.ListVals[i+3]
				potentialNewValue = parseIdentifierToBeBound(firstListItem, potentialNewValueItem, &env, global, lineNum, caller)
				howManyIndicesToJumpForward = 3
				isTypeAnnotated = true
			}
		} else {
			potentialNewValue = parseIdentifierToBeBound(firstListItem, nextListItem, &env, global, lineNum, caller)
		}
		if isTypeAnnotated {
			if (annotatedTypeName == "" && TypesEqualP(annotatedTypeValue, &env.findBinding(potentialNewValue.TypeName, true, true).Binding, lineNum, caller)) ||
				(potentialNewValue.TypeName == "undecided" && potentialNewValue.TypeName == annotatedTypeName) {
				if annotatedTypeName != "" {
					potentialNewValue.TypeName = annotatedTypeName
				}
			} else {
				errMsg := fmt.Sprintf("Error: attempting to assign type %v to %v in %v at line %v, but it is already of type %v.\n", annotatedTypeName, firstListItem.Value, caller, lineNum, potentialNewValue.TypeName)
				panic(errMsg)
			}
		}
		potentialNewValue.Mutable = mut
		newBinding.Binding = *potentialNewValue
		i += howManyIndicesToJumpForward
	}
	return env
}

func evalListToken(list *Parser.Token) ListCell {
	firstVal := &list.ListVals[0]
	//	initCellList := CellList{}
	initEnvironment := Environment{}
	switch firstVal.Type {
	case Parser.LiteralToken:
		errMsg := fmt.Sprintf("Error: attempting to evaluate a literal, %v, at line %v.\n", firstVal.Value, firstVal.LineNum)
		panic(errMsg)
	case Parser.DefToken:
		defKind := &firstVal.Value
		if len(list.ListVals) < 3 {
			errMsg := fmt.Sprintf("Error: too few arguments to %v at line %v.\n", defKind, firstVal.LineNum)
			panic(errMsg)
		} else if list.ListVals[1].Type != Parser.ListToken {
			errMsg := fmt.Sprintf("Error: first argument (%v) to %v at line %v is not a list.\n", list.ListVals[1].Value, defKind, firstVal.LineNum)
			panic(errMsg)
		} else if list.ListVals[2].Type != Parser.ListToken {
			errMsg := fmt.Sprintf("Error: second argument (%v) to %v  at line %v is not a list.\n", list.ListVals[2].Value, defKind, firstVal.LineNum)
			panic(errMsg)
		} else if len(list.ListVals) > 3 {
			errMsg := fmt.Sprintf("Error: too many arguments to %v at line %v.\n", defKind, firstVal.LineNum)
			panic(errMsg)
		}
		tmpCallerName := "let"
		initEnvironment = bindVars(&list.ListVals[1], initEnvironment, firstVal.LineNum, true, true, &tmpCallerName)
	default:
		errMsg := fmt.Sprintf("Error: unhandled token type for %v at line %v.\n", firstVal.Value, firstVal.LineNum)
		panic(errMsg)
	}
	return ListCell{}
}

func Initialise(input string)ListCell{
	res := Parser.Lex(&input)
	tokens := Parser.ParseList(res, 0)
	for _, tok := range tokens.ListVals {
		fmt.Println(tok)
	}
	return evalListToken(&tokens.ListVals[0])
}
