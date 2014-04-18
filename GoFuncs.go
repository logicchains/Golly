package Golly

import (
	"errors"
	"fmt"
)

type baseType int

const (
	Int baseType = iota
	Float
	Char
	Symbol
	List
	FuncDef
	VarDef
)

type goFuncType int16

const (
	GoAddT goFuncType = iota
	GoSubtractT
	GoMultiplyT
	GoDivideT
	GoIfT
	GoEvalT
)

const (
	FUNCTION_TYPE_NAME    = "Function"
	LIST_TYPE_NAME        = "List"
	ENVIRONMENT_TYPE_NAME = "Environment"
	VAR_TYPE_NAME         = "Var"
)

func CallGoFunc(funcType goFuncType, parameters []ListCell) ([]*ListCell, error) {
	switch funcType {
	case GoAddT:
		res, err := GoAdd(&parameters[0], &parameters[1])
		if err != nil {
			return nil, err
		} else {
			return res, nil
		}
	case GoSubtractT:
		res, err := GoSubtract(&parameters[0], &parameters[1])
		if err != nil {
			return nil, err
		} else {
			return res, nil
		}
	case GoMultiplyT:
		res, err := GoMultiply(&parameters[0], &parameters[1])
		if err != nil {
			return nil, err
		} else {
			return res, nil
		}
	case GoDivideT:
		res, err := GoDivide(&parameters[0], &parameters[1])
		if err != nil {
			return nil, err
		} else {
			return res, nil
		}
	case GoIfT:
		res, err := GoIf(&parameters[0], &parameters[1], &parameters[2])
		if err != nil {
			return nil, err
		} else {
			return res, nil
		}
	case GoEvalT:
		res, err := GoEval(&parameters[0], &parameters[1])
		if err != nil {
			return nil, err
		} else {
			return res, nil
		}

	default:
		err := fmt.Sprintf("Error: attempting to call unhandled builtin function of type number %v.\n", funcType)
		return nil, errors.New(err)
	}
}

func GoAdd(Cell1 *ListCell, Cell2 *ListCell) ([]*ListCell, error) {
	if Cell1.TypeName != Cell2.TypeName {
		err := fmt.Sprintf("Error: attempting to add type %v to type %v, but these types are not compatible.\n", Cell1.TypeName, Cell2.TypeName)
		return nil, errors.New(err)
	}
	returnVals := make([]*ListCell, 0, 1)
	returnVal := ListCell{TypeName: Cell1.TypeName, Mutable: Cell1.Mutable}
	if val1, ok1 := Cell1.Value.(int); ok1 {
		if val2, ok2 := Cell2.Value.(int); ok2 {
			returnVal.Value = val1 + val2
		} else {
			err := fmt.Sprintf("Error: attempting to add type %v to type %v, but the first really was an int and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int64); ok1 {
		if val2, ok2 := Cell2.Value.(int64); ok2 {
			returnVal.Value = val1 + val2
		} else {
			err := fmt.Sprintf("Error: attempting to add type %v to type %v, but the first really was an int64 and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int32); ok1 {
		if val2, ok2 := Cell2.Value.(int32); ok2 {
			returnVal.Value = val1 + val2
		} else {
			err := fmt.Sprintf("Error: attempting to add type %v to type %v, but the first really was an int32 and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int16); ok1 {
		if val2, ok2 := Cell2.Value.(int16); ok2 {
			returnVal.Value = val1 + val2
		} else {
			err := fmt.Sprintf("Error: attempting to add type %v to type %v, but the first really was an int16 and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(float64); ok1 {
		if val2, ok2 := Cell2.Value.(float64); ok2 {
			returnVal.Value = val1 + val2
		} else {
			err := fmt.Sprintf("Error: attempting to add type %v to type %v, but the first really was a float64 and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(float32); ok1 {
		if val2, ok2 := Cell2.Value.(float32); ok2 {
			returnVal.Value = val1 + val2
		} else {
			err := fmt.Sprintf("Error: attempting to add type %v to type %v, but the first really was a float32 and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	returnVals = append(returnVals, &returnVal)
	return returnVals, nil
}

func GoSubtract(Cell1 *ListCell, Cell2 *ListCell) ([]*ListCell, error) {
	if Cell1.TypeName != Cell2.TypeName {
		err := fmt.Sprintf("Error: attempting to subtract type %v from type %v, but these types are not compatible.\n", Cell2.TypeName, Cell1.TypeName)
		return nil, errors.New(err)
	}
	returnVals := make([]*ListCell, 0, 1)
	returnVal := ListCell{TypeName: Cell1.TypeName, Mutable: Cell1.Mutable}
	if val1, ok1 := Cell1.Value.(int); ok1 {
		if val2, ok2 := Cell2.Value.(int); ok2 {
			returnVal.Value = val1 - val2
		} else {
			err := fmt.Sprintf("Error: attempting to subtract type %v from type %v, but the first really was an int and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int64); ok1 {
		if val2, ok2 := Cell2.Value.(int64); ok2 {
			returnVal.Value = val1 - val2
		} else {
			err := fmt.Sprintf("Error: attempting to subtract type %v from type %v, but the first really was an int64 and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int32); ok1 {
		if val2, ok2 := Cell2.Value.(int32); ok2 {
			returnVal.Value = val1 - val2
		} else {
			err := fmt.Sprintf("Error: attempting to subtract type %v from type %v, but the first really was an int32 and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int16); ok1 {
		if val2, ok2 := Cell2.Value.(int16); ok2 {
			returnVal.Value = val1 - val2
		} else {
			err := fmt.Sprintf("Error: attempting to subtract type %v from type %v, but the first really was an int16 and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(float64); ok1 {
		if val2, ok2 := Cell2.Value.(float64); ok2 {
			returnVal.Value = val1 - val2
		} else {
			err := fmt.Sprintf("Error: attempting to subtract type %v from type %v, but the first really was a float64 and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(float32); ok1 {
		if val2, ok2 := Cell2.Value.(float32); ok2 {
			returnVal.Value = val1 - val2
		} else {
			err := fmt.Sprintf("Error: attempting to subtract type %v from type %v, but the first really was a float32 and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	returnVals = append(returnVals, &returnVal)
	return returnVals, nil
}

func GoMultiply(Cell1 *ListCell, Cell2 *ListCell) ([]*ListCell, error) {
	if Cell1.TypeName != Cell2.TypeName {
		err := fmt.Sprintf("Error: attempting to multiply type %v by type %v, but these types are not compatible.\n", Cell1.TypeName, Cell2.TypeName)
		return nil, errors.New(err)
	}
	returnVals := make([]*ListCell, 0, 1)
	returnVal := ListCell{TypeName: Cell1.TypeName, Mutable: Cell1.Mutable}
	if val1, ok1 := Cell1.Value.(int); ok1 {
		if val2, ok2 := Cell2.Value.(int); ok2 {
			returnVal.Value = val1 * val2
		} else {
			err := fmt.Sprintf("Error: attempting to multiply type %v by type %v, but the first really was an int and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int64); ok1 {
		if val2, ok2 := Cell2.Value.(int64); ok2 {
			returnVal.Value = val1 * val2
		} else {
			err := fmt.Sprintf("Error: attempting to multiply type %v by type %v, but the first really was an int64 and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int32); ok1 {
		if val2, ok2 := Cell2.Value.(int32); ok2 {
			returnVal.Value = val1 * val2
		} else {
			err := fmt.Sprintf("Error: attempting to multiply type %v by type %v, but the first really was an int32 and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int16); ok1 {
		if val2, ok2 := Cell2.Value.(int16); ok2 {
			returnVal.Value = val1 * val2
		} else {
			err := fmt.Sprintf("Error: attempting to multiply type %v by type %v, but the first really was an int16 and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(float64); ok1 {
		if val2, ok2 := Cell2.Value.(float64); ok2 {
			returnVal.Value = val1 * val2
		} else {
			err := fmt.Sprintf("Error: attempting to multiply type %v by type %v, but the first really was a float64 and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(float32); ok1 {
		if val2, ok2 := Cell2.Value.(float32); ok2 {
			returnVal.Value = val1 * val2
		} else {
			err := fmt.Sprintf("Error: attempting to multiply type %v by type %v, but the first really was a float32 and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	returnVals = append(returnVals, &returnVal)
	return returnVals, nil
}

func GoDivide(Cell1 *ListCell, Cell2 *ListCell) ([]*ListCell, error) {
	if Cell1.TypeName != Cell2.TypeName {
		err := fmt.Sprintf("Error: attempting to divide type %v by type %v, but these types are not compatible.\n", Cell2.TypeName, Cell1.TypeName)
		return nil, errors.New(err)
	}
	returnVals := make([]*ListCell, 0, 1)
	returnVal := ListCell{TypeName: Cell1.TypeName, Mutable: Cell1.Mutable}
	if val1, ok1 := Cell1.Value.(int); ok1 {
		if val2, ok2 := Cell2.Value.(int); ok2 {
			returnVal.Value = val1 / val2
		} else {
			err := fmt.Sprintf("Error: attempting to divide type %v by type %v, but the first really was an int and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int64); ok1 {
		if val2, ok2 := Cell2.Value.(int64); ok2 {
			returnVal.Value = val1 / val2
		} else {
			err := fmt.Sprintf("Error: attempting to divide type %v by type %v, but the first really was an int64 and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int32); ok1 {
		if val2, ok2 := Cell2.Value.(int32); ok2 {
			returnVal.Value = val1 / val2
		} else {
			err := fmt.Sprintf("Error: attempting to divide type %v by type %v, but the first really was an int32 and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int16); ok1 {
		if val2, ok2 := Cell2.Value.(int16); ok2 {
			returnVal.Value = val1 / val2
		} else {
			err := fmt.Sprintf("Error: attempting to divide type %v by type %v, but the first really was an int16 and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(float64); ok1 {
		if val2, ok2 := Cell2.Value.(float64); ok2 {
			returnVal.Value = val1 / val2
		} else {
			err := fmt.Sprintf("Error: attempting to divide type %v by type %v, but the first really was a float64 and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(float32); ok1 {
		if val2, ok2 := Cell2.Value.(float32); ok2 {
			returnVal.Value = val1 / val2
		} else {
			err := fmt.Sprintf("Error: attempting to divide type %v by type %v, but the first really was a float32 and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	returnVals = append(returnVals, &returnVal)
	return returnVals, nil
}

func GoIf(Cell1 *ListCell, Cell2 *ListCell, Cell3 *ListCell) ([]*ListCell, error) {
	returnVals := make([]*ListCell, 0, 1)
	returnVal := new(ListCell)
	if condVal, ok := Cell1.Value.(bool); ok {
		if condVal {
			returnVal = Cell2
		} else {
			returnVal = Cell3
		}

	} else {
		err := fmt.Sprintf("Error: expected bool as first argument to if builtin but got something else.\n", Cell1.TypeName)
		return nil, errors.New(err)
	}
	returnVals = append(returnVals, returnVal)
	return returnVals, nil
}

func EvalPrim(list []ListCell, env Environment) ([]*ListCell, error) {
	if varName, ok := list[0].Value.(string); ok {
		binding := env.findBinding(varName, true, true)
		if binding == nil {
			err := fmt.Sprintf("Error: var in first cell in list passed to eval builtin, %v, is not bound.\n", varName)
			return nil, errors.New(err)
		}
		if funct, ok := binding.Binding.Value.(FunctionObj); ok {
			res, err := funct.Call(list[1:], env)
			if err != nil {
				return nil, err
			} else {
				return res, nil
			}
		} else {
			err := fmt.Sprintf("Error: expected var in first cell of list passed to eval builtin, but it was actually not a var.\n")
			return nil, errors.New(err)
		}
	} else {
		err := fmt.Sprintf("Error: expected var in first cell of list passed to eval builtin, but it was actually not a var.\n")
		return nil, errors.New(err)
	}
	return nil, nil
}

func GoEval(Cell1 *ListCell, Cell2 *ListCell) ([]*ListCell, error) {
	returnVals := make([]*ListCell, 0, 0)
	if list, ok := Cell1.Value.([]ListCell); ok {
		if env, ok2 := Cell1.Value.(Environment); ok2 {
			returnValShad, err := EvalPrim(list, env)
			if err != nil {
				return nil, err
			}
			returnVals = returnValShad
		} else {
			err := fmt.Sprintf("Error: expected environment as second argument to eval builtin but got something else.\n")
			return nil, errors.New(err)
		}
	} else {
		err := fmt.Sprintf("Error: expected list as first argument to eval builtin but got something else.\n")
		return nil, errors.New(err)
	}
	return returnVals, nil
}

func Eval(list []ListCell, env Environment) ([]*ListCell, error) {
	returnVals, err := EvalPrim(list, env)
	if err != nil {
		return nil, err
	}else{
		return returnVals, nil
	}
}

func parseText(input string) (CellList, error){
	for i := 0; i < len(input); i++{
		if input[i] == ' '{
			continue
		}else if input[i] == '('{
			parseList(input[i:])
		}
	}
}

func parseList(input string) (CellList, error){
	list := CellList{}
	listCells := make([]ListCell,0,10)
	for i := 0; i < len(input); i++{
		switch input[i]{
		case ' ':
			continue
		case '"':
			str, err := parseStringLit(input[i:])
			if err != nil{
				return CellList{}, err
			}else{
				listCells = append(listCells, ListCell{Mutable: true, Value: str})
			}
		}
	}
	list.Cells = listCells
	return list, nil
}

func parseStringLit(input string) (string, error){
	for i := 1; i < len(input); i++{
		if input[i] == '"' && input [i-1] != '\\'{
			return fmt.Sprint(input), nil
		}
	}
	err := fmt.Sprintf("Error: unterminated string encountered.\n")
	return "", errors.New(err)
}
