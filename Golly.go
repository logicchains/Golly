package main

import( 
	"fmt"
	"Golly/parser"
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

type TypeDef struct{
	Name string
}

type EnvCell struct{
	TypeID int64
	Name string
	Value interface{}
	Mutable bool
}

type ListCell struct{
	TypeID int64
	Value interface{}
	Mutable bool
}

type CellList struct{
	Cells []ListCell
	Environment []EnvCell
}

func findBinding(name string, environ []EnvCell, recur bool) *EnvCell{
	for i, cell := range environ{
		if cell.Name == name{
			return &environ[i]
		}
	}
	if recur{
		if parentEnv, ok := environ[0].Value.([]EnvCell); ok {
			return findBinding(name, parentEnv, true)
		} else {
			return nil
		}
	}
	return nil
}

func defGlobal(environ *[]EnvCell) *EnvCell{
	if parentEnv, ok := (*environ)[0].Value.([]EnvCell); ok {
		return defGlobal(&parentEnv)
	} else {
		*environ = append(*environ, EnvCell{})
		return &((*environ)[len(*environ)])
	}
}

func evalType(token *Parser.Token) (TypeDef, error){
	if token.Type != Parser.IdToken{

	}
	return TypeDef{}, nil
}

func bindVars(list *Parser.Token, environ []EnvCell, lineNum int, global, mut bool, caller string)[]EnvCell{
	for i := 0; i < len(list.ListVals); i++{
		val := &list.ListVals[i]
		if val.Type != Parser.IdToken{
			errMsg := fmt.Sprintf("Error: attempting to assign to a non-identifier in %v at line %v.\n", caller, lineNum) 
			panic(errMsg)				
		}
		prevBinding := findBinding(val.Value, environ, global)
		var newBinding *EnvCell
		if prevBinding != nil{
			if !prevBinding.Mutable{
				errMsg := fmt.Sprintf("Error: attempting to assign to an immutable identifier in %v at line %v.\n", caller, lineNum) 
				panic(errMsg)				
			}else{
				newBinding = prevBinding
			}
		}else{
			if global{
				newBinding = defGlobal(&environ)
			}else{
				environ = append(environ, EnvCell{}) 
				newBinding = &(environ[len(environ)])
			}
		}
		newBinding.Name = val.Value
		newBinding.Mutable = mut
		nextVal := &list.ListVals[i+1]
		var newValType TypeDef
		if (*nextVal).Type == Parser.TypeDefToken{
			newValType, err := evalType(nextVal)
		}
		newVal := evalListToken(nextVal)

	}
	return environ
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
