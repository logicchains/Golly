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

type Union struct{
	CurType int16 
	Types []string 
}

type EnvBinding struct{
	Name string
	Binding interface{}
}

type Environment []EnvBinding

func (env Environment) findBinding(name string, recur bool) *EnvBinding{
	for i, _ := range env{
		if env[i].Name == name{
			return &env[i]
		}
	}
	if recur{
		if parentEnv, ok := env[0].Binding.(Environment); ok {
			return parentEnv.findBinding(name, true)
		} else {
			return nil
		}
	}
	return nil
}

func (env *Environment) addBinding(recur bool) *EnvBinding{
	if recur{
		if parentEnv, ok := (*env)[0].Binding.(Environment); ok {
			return parentEnv.addBinding(true)
		} else {
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

func defGlobal(environ *[]EnvBinding) *EnvBinding{
	if parentEnv, ok := (*environ)[0].Binding.([]EnvBinding); ok {
		return defGlobal(&parentEnv)
	} else {
		*environ = append(*environ, EnvBinding{})
		return &((*environ)[len(*environ)])
	}
}

func evalNumToken(num *Parser.Token, lineNum int, caller string)(ListCell){
	newValue := ListCell{}
			if (*num).NumType == Parser.FixNum{
				floatval, err := strconv.ParseFloat((*num).Value, 32)
				if err != nil{
					errMsg := fmt.Sprintf("Error: cannot parse string %v to float in %v at line %v.\n", (*num).Value, caller, lineNum) 
					panic(errMsg)
				}else{
					newValue.Value = floatval
					newValue.TypeName = "float"
				}
			}else if (*num).NumType == Parser.FixNum{
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
func bindVars(list *Parser.Token, env Environment, lineNum int, global, mut bool, caller string)[]EnvBinding{
	for i := 0; i < len(list.ListVals); i++{
		val := &list.ListVals[i]
		if val.Type != Parser.IdToken{
			errMsg := fmt.Sprintf("Error: attempting to assign to a non-identifier in %v at line %v.\n", caller, lineNum) 
			panic(errMsg)				
		}
		prevBinding := env.findBinding(val.Value, global)
		var newBinding *EnvBinding
		if prevBinding != nil{
			if boundVal, ok := (*prevBinding).Binding.(ListCell); ok {
				if !boundVal.Mutable{
					errMsg := fmt.Sprintf("Error: attempting to assign to an immutable identifier in %v at line %v.\n", caller, lineNum) 
					panic(errMsg)				
				}else{
					newBinding = prevBinding
				}
			} else {	
				errMsg := fmt.Sprintf("Error: malformed environment binding encountered in binding for %v in %v at line %v.\n", val.Value, caller, lineNum) 
				panic(errMsg)				
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
		newValue := ListCell{}
		switch (*nextVal).Type {
		case Parser.NumToken:
			newValue = evalNumToken(nextVal,lineNum,caller)
		case Parser.DefToken:
		case Parser.ListToken:
			newValue.Value = evalListToken(nextVal)
		}
		//newBinding.Mutable = mut

		//var newValType TypeDef
		if (*nextVal).Type == Parser.TypeDefToken{
			//newValType, err := evalType(nextVal)
		}
		//newVal := evalListToken(nextVal)

	}
	return env
}

func evalListToken(list *Parser.Token)(ListCell){
	firstVal := &list.ListVals[0]
//	initCellList := CellList{}
	switch firstVal.Type{
	case Parser.NumToken: 
		errMsg := fmt.Sprintf("Error: attempting to evaluate a literal, %v, at line %v.\n",firstVal.Value, firstVal.LineNum) 
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
