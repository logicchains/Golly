package main

import( 
	"fmt"
	"Golly/parser"
	"strconv"
)

type baseType int
const(
	Int baseType = iota
	Float
	Char
	Symbol
	List
	FuncDef
	VarDef
)

type FuncParem struct{
	Name string
	TypeID int64
}

type FunctionDef struct{
	NumParems int16
	Parems []FuncParem
	ReturnVals []FuncParem
}

type singleType struct{
	Inputs []string
	Outputs []string
	Order int8 
}

type TypeDef struct{
	Name string
	Types []singleType
}

type Union struct{
	CurType int16 
	Types []string 
}

type EnvBinding struct{
	Name string
	Binding ListCell
}

type Environment []EnvBinding

func (env Environment) findBinding(name string, recur bool) *EnvBinding{
	for i, _ := range env{
		if env[i].Name == name{
			return &env[i]
		}
	}
	if recur{
		if len(env) == 0{
			panic("Error: root environment is empty!\n")
		}
		if env[0].Binding.TypeName == "environment"{
			if parentEnv, ok := env[0].Binding.Value.(Environment); ok {	
				return parentEnv.findBinding(name, true)
			}else{
				panic("Error: encountered an environment-typed cell that contained no environment! This shouldn't happen; report it as a bug.\n")
			}
		}else{
			return nil
		}
	}
	return nil
}

func (env *Environment) addBinding(recur bool) *EnvBinding{
	if recur{
		if (*env)[0].Binding.TypeName == "environment"{
			if parentEnv, ok := (*env)[0].Binding.Value.(Environment); ok {
				return parentEnv.addBinding(true)
			}else{
				panic("Error: encountered an environment-typed cell that contained no environment! This shouldn't happen; report it as a bug.\n")
			}		
		}else {
			*env = append(*env, EnvBinding{})
			return &((*env)[len(*env)])
		}
	}else{
		*env = append(*env, EnvBinding{})
		return &((*env)[len(*env)])
	}
}

type ListCell struct{
	TypeName string
	Value interface{}
	Mutable bool
}

type CellList struct{
	Cells []ListCell
	Environment []EnvBinding
}

func evalNumToken(num *Parser.Token, lineNum int, caller string)(ListCell){
	newValue := ListCell{}
			if (*num).LitType == Parser.FixNum{
				floatval, err := strconv.ParseFloat((*num).Value, 32)
				if err != nil{
					errMsg := fmt.Sprintf("Error: cannot parse string %v to float in %v at line %v.\n", (*num).Value, caller, lineNum) 
					panic(errMsg)
				}else{
					newValue.Value = floatval
					newValue.TypeName = "float"
				}
			}else if (*num).LitType == Parser.FixNum{
				intval, err := strconv.Atoi((*num).Value)
				if err != nil{
					errMsg := fmt.Sprintf("Error: cannot parse string %v to int in %v at line %v.\n", (*num).Value, caller, lineNum) 
					panic(errMsg)
				}else{
					newValue.Value = intval
					newValue.TypeName = "int"
				}
			}
	return newValue
}

func evalIdToken(identifierName *Parser.Token, env *Environment, lineNum int, caller string)(interface{}){
	var newValue interface{}
	valueReferenced := env.findBinding((*identifierName).Value, true)
			if valueReferenced == nil{
				errMsg := fmt.Sprintf("Error: attempting to evalute var %v in %v at line %v, but that var is unbound.\n", (*identifierName).Value, caller, lineNum) 
				panic(errMsg)
			}else{
				newValue = valueReferenced.Binding
			}
	return newValue
}
func bindVars(list *Parser.Token, env Environment, lineNum int, global, mut bool, caller string)Environment{
	for i := 0; i < len(list.ListVals); i++{
		val := &list.ListVals[i]
		if val.Type != Parser.IdToken{
			errMsg := fmt.Sprintf("Error: attempting to assign to a non-identifier in %v at line %v.\n", caller, lineNum) 
			panic(errMsg)				
		}
		prevBinding := env.findBinding(val.Value, global)
		var newBinding *EnvBinding
		if prevBinding != nil{
			if !(*prevBinding).Binding.Mutable{
				errMsg := fmt.Sprintf("Error: attempting to assign to an immutable identifier in %v at line %v.\n", caller, lineNum) 
				panic(errMsg)				
			}else{
				newBinding = prevBinding
			}
		}else{
			newBinding = env.addBinding(global)
		}
		if i >= len(list.ListVals){
			errMsg := fmt.Sprintf("Error: nothing to assign to %v in %v at line %v.\n", val.Value, caller, lineNum) 
			panic(errMsg)
		}
		nextVal := &list.ListVals[i+1]
		newBinding.Name = val.Value
		newValue := ListCell{TypeName: "undecided", Mutable: mut}
		typeNameAnnotated := ""
		switch (*nextVal).Type {
		case Parser.LiteralToken:
			newValue = evalNumToken(nextVal,lineNum,caller)
		case Parser.DefToken:
			errMsg := fmt.Sprintf("Error: attempting to assign reserved name %v to %v in %v at line %v.\n", (*nextVal).Value, val.Value, caller, lineNum) 
			panic(errMsg)
		case Parser.ListToken:
			newValue.Value = evalListToken(nextVal)
		case Parser.TypeAnnToken:
			if i >= len(list.ListVals)+1{
				errMsg := fmt.Sprintf("Error: no type provided in assignment to %v in %v at line %v.\n", val.Value, caller, lineNum) 
				panic(errMsg)
			}
			nextValType := &list.ListVals[i+2]
			newValueType := ListCell{}
			newNameFound := false
			switch (*nextValType).Type {
			case Parser.LiteralToken:
				errMsg := fmt.Sprintf("Error: attempting use a numeric literal as the type for %v in %v at line %v.\n", val.Value, caller, lineNum) 
				panic(errMsg)
			case Parser.DefToken:
				errMsg := fmt.Sprintf("Error: attempting use a reserved name as the type for %v in %v at line %v.\n", val.Value, caller, lineNum) 
				panic(errMsg)
			case Parser.ListToken:
				newValueType = evalListToken(nextVal)
				if newValueType.TypeName != "type"{
					errMsg := fmt.Sprintf("Error: attempting to assign something that is not a type, but a %v, to %v in %v at line %v.\n", newValueType.TypeName, val.Value, caller, lineNum)
					panic(errMsg)
				}
				newNameFound = true
			case Parser.TypeAnnToken:
				errMsg := fmt.Sprintf("Error: misplaced type annotation marker in %v at line %v.\n", val.Value, caller, lineNum) 
				panic(errMsg)
			}
			typeName := nextValType.Value
			if newNameFound{
				if foundTypeActual, ok := newValueType.Value.(TypeDef); ok {
					typeName = foundTypeActual.Name
				}else{
					errMsg := fmt.Sprintf("Error: cell claiming to be a type actually contains something else, in %v at line %v.\n", caller, lineNum) 
					panic(errMsg)
				}
			}
			namesBinding := env.findBinding(typeName, true)
			if namesBinding.Binding.TypeName == "type"{
				typeNameAnnotated = typeName	
			}else{
				errMsg := fmt.Sprintf("Error: attempting to assign type %v to %v in %v at line %v, but that type is not bound.\n", typeName, val.Value, caller, lineNum) 
				panic(errMsg)
			}
		}
		if newValue.TypeName != "undecided" && newValue.TypeName != typeNameAnnotated{
			errMsg := fmt.Sprintf("Error: attempting to assign type %v to %v in %v at line %v, but it is already of type %v.\n", typeNameAnnotated, val.Value, caller, lineNum, newValue.TypeName) 
			panic(errMsg)
		}else if typeNameAnnotated != ""{
			newValue.TypeName = typeNameAnnotated
		}
		newBinding.Binding = newValue
	}
	return env
}

func evalListToken(list *Parser.Token)(ListCell){
	firstVal := &list.ListVals[0]
//	initCellList := CellList{}
	initEnvironment := Environment{}
	switch firstVal.Type{
	case Parser.LiteralToken: 
		errMsg := fmt.Sprintf("Error: attempting to evaluate a literal, %v, at line %v.\n", firstVal.Value, firstVal.LineNum) 
		panic(errMsg)
	case Parser.DefToken:
		defKind := &firstVal.Value
		if len(list.ListVals) < 3{
			errMsg := fmt.Sprintf("Error: too few arguments to %v at line %v.\n", defKind, firstVal.LineNum) 
			panic(errMsg)				
		}else if list.ListVals[1].Type != Parser.ListToken{
			errMsg := fmt.Sprintf("Error: first argument (%v) to %v at line %v is not a list.\n",list.ListVals[1].Value, defKind, firstVal.LineNum) 
			panic(errMsg)				
		}else if list.ListVals[2].Type != Parser.ListToken{
			errMsg := fmt.Sprintf("Error: second argument (%v) to %v  at line %v is not a list.\n",list.ListVals[2].Value, defKind, firstVal.LineNum) 
			panic(errMsg)				
		}else if len(list.ListVals) > 3{
			errMsg := fmt.Sprintf("Error: too many arguments to %v at line %v.\n",defKind, firstVal.LineNum) 
			panic(errMsg)				
		}
		initEnvironment = bindVars(&list.ListVals[1], initEnvironment, firstVal.LineNum, true, true, "let")
	
			
	}
	return ListCell{} 
}

func main(){
	//types := []string{"Int", "Float", "Char", "Symbol", "List"}
	input := `(let (a 1) (b 2))`
	res := Parser.Lex(&input)
	tokens := Parser.ParseList(res, 0)
	for _, tok := range tokens.ListVals{
		fmt.Println(tok)
	}
	evalListToken(&tokens.ListVals[0])
}
