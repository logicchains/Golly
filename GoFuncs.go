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

const(
	FUNCTION_TYPE_NAME = "Function"
	LIST_TYPE_NAME = "List"
	ENVIRONMENT_TYPE_NAME = "Environment"
	VAR_TYPE_NAME = "Var"
)

type GoFunc func (Cell1 *ListCell, Cell2 *ListCell)(*ListCell, error)

func CallGoFunc(funcType goFuncType, parameters []*ListCell)([]*ListCell, error){
	switch funcType{
	case GoAddT:
		res, err := GoAdd(parameters[0], parameters[1])
		if err != nil{
			return nil, err
		}else{
			return res, nil
		}
	case GoSubtractT:
		res, err := GoSubtract(parameters[0], parameters[1])
		if err != nil{
			return nil, err
		}else{
			return res, nil
		}
	case GoMultiplyT:
		res, err := GoMultiply(parameters[0], parameters[1])
		if err != nil{
			return nil, err
		}else{
			return res, nil
		}
	case GoDivideT:
		res, err := GoDivide(parameters[0], parameters[1])
		if err != nil{
			return nil, err
		}else{
			return res, nil
		}
	case GoIfT:
		res, err := GoIf(parameters[0], parameters[1])
		if err != nil{
			return nil, err
		}else{
			return res, nil
		}
	default:
			err := fmt.Sprintf("Error: attempting to call unhandled builtin function of type number %v.\n", funcType)
			return nil, errors.New(err)
	}
}

func GoAdd(Cell1 *ListCell, Cell2 *ListCell)([]*ListCell, error){
	if Cell1.TypeName != Cell2.TypeName{
		err := fmt.Sprintf("Error: attempting to add type %v to type %v, but these types are not compatible.\n", Cell1.TypeName, Cell2.TypeName)
		return nil, errors.New(err)
	}
	returnVals := make([]*ListCell, 0, 1)
	returnVal := ListCell{TypeName: Cell1.TypeName, Mutable: Cell1.Mutable}
	if val1, ok1 := Cell1.Value.(int); ok1 {
		if val2, ok2 := Cell2.Value.(int); ok2 {
			returnVal.Value = val1 + val2
		}else{
			err := fmt.Sprintf("Error: attempting to add type %v to type %v, but the first really was an int and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int64); ok1 {
		if val2, ok2 := Cell2.Value.(int64); ok2 {
			returnVal.Value = val1 + val2
		}else{
			err := fmt.Sprintf("Error: attempting to add type %v to type %v, but the first really was an int64 and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int32); ok1 {
		if val2, ok2 := Cell2.Value.(int32); ok2 {
			returnVal.Value = val1 + val2
		}else{
			err := fmt.Sprintf("Error: attempting to add type %v to type %v, but the first really was an int32 and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int16); ok1 {
		if val2, ok2 := Cell2.Value.(int16); ok2 {
			returnVal.Value = val1 + val2
		}else{
			err := fmt.Sprintf("Error: attempting to add type %v to type %v, but the first really was an int16 and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(float64); ok1 {
		if val2, ok2 := Cell2.Value.(float64); ok2 {
			returnVal.Value = val1 + val2
		}else{
			err := fmt.Sprintf("Error: attempting to add type %v to type %v, but the first really was a float64 and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(float32); ok1 {
		if val2, ok2 := Cell2.Value.(float32); ok2 {
			returnVal.Value = val1 + val2
		}else{
			err := fmt.Sprintf("Error: attempting to add type %v to type %v, but the first really was a float32 and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	returnVals = append(returnVals, &returnVal)
	return returnVals, nil
}

func GoSubtract(Cell1 *ListCell, Cell2 *ListCell)([]*ListCell, error){
	if Cell1.TypeName != Cell2.TypeName{
		err := fmt.Sprintf("Error: attempting to subtract type %v from type %v, but these types are not compatible.\n", Cell2.TypeName, Cell1.TypeName)
		return nil, errors.New(err)
	}
	returnVals := make([]*ListCell, 0, 1)
	returnVal := ListCell{TypeName: Cell1.TypeName, Mutable: Cell1.Mutable}
	if val1, ok1 := Cell1.Value.(int); ok1 {
		if val2, ok2 := Cell2.Value.(int); ok2 {
			returnVal.Value = val1 - val2
		}else{
			err := fmt.Sprintf("Error: attempting to subtract type %v from type %v, but the first really was an int and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int64); ok1 {
		if val2, ok2 := Cell2.Value.(int64); ok2 {
			returnVal.Value = val1 - val2
		}else{
			err := fmt.Sprintf("Error: attempting to subtract type %v from type %v, but the first really was an int64 and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int32); ok1 {
		if val2, ok2 := Cell2.Value.(int32); ok2 {
			returnVal.Value = val1 - val2
		}else{
			err := fmt.Sprintf("Error: attempting to subtract type %v from type %v, but the first really was an int32 and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int16); ok1 {
		if val2, ok2 := Cell2.Value.(int16); ok2 {
			returnVal.Value = val1 - val2
		}else{
			err := fmt.Sprintf("Error: attempting to subtract type %v from type %v, but the first really was an int16 and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(float64); ok1 {
		if val2, ok2 := Cell2.Value.(float64); ok2 {
			returnVal.Value = val1 - val2
		}else{
			err := fmt.Sprintf("Error: attempting to subtract type %v from type %v, but the first really was a float64 and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(float32); ok1 {
		if val2, ok2 := Cell2.Value.(float32); ok2 {
			returnVal.Value = val1 - val2
		}else{
			err := fmt.Sprintf("Error: attempting to subtract type %v from type %v, but the first really was a float32 and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	returnVals = append(returnVals, &returnVal)
	return returnVals, nil
}

func GoMultiply(Cell1 *ListCell, Cell2 *ListCell)([]*ListCell, error){
	if Cell1.TypeName != Cell2.TypeName{
		err := fmt.Sprintf("Error: attempting to multiply type %v by type %v, but these types are not compatible.\n", Cell1.TypeName, Cell2.TypeName)
		return nil, errors.New(err)
	}
	returnVals := make([]*ListCell, 0, 1)
	returnVal := ListCell{TypeName: Cell1.TypeName, Mutable: Cell1.Mutable}
	if val1, ok1 := Cell1.Value.(int); ok1 {
		if val2, ok2 := Cell2.Value.(int); ok2 {
			returnVal.Value = val1 * val2
		}else{
			err := fmt.Sprintf("Error: attempting to multiply type %v by type %v, but the first really was an int and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int64); ok1 {
		if val2, ok2 := Cell2.Value.(int64); ok2 {
			returnVal.Value = val1 * val2
		}else{
			err := fmt.Sprintf("Error: attempting to multiply type %v by type %v, but the first really was an int64 and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int32); ok1 {
		if val2, ok2 := Cell2.Value.(int32); ok2 {
			returnVal.Value = val1 * val2
		}else{
			err := fmt.Sprintf("Error: attempting to multiply type %v by type %v, but the first really was an int32 and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int16); ok1 {
		if val2, ok2 := Cell2.Value.(int16); ok2 {
			returnVal.Value = val1 * val2
		}else{
			err := fmt.Sprintf("Error: attempting to multiply type %v by type %v, but the first really was an int16 and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(float64); ok1 {
		if val2, ok2 := Cell2.Value.(float64); ok2 {
			returnVal.Value = val1 * val2
		}else{
			err := fmt.Sprintf("Error: attempting to multiply type %v by type %v, but the first really was a float64 and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(float32); ok1 {
		if val2, ok2 := Cell2.Value.(float32); ok2 {
			returnVal.Value = val1 * val2
		}else{
			err := fmt.Sprintf("Error: attempting to multiply type %v by type %v, but the first really was a float32 and the second wasn't.\n", Cell1.TypeName, Cell2.TypeName)
			return nil, errors.New(err)
		}
	}
	returnVals = append(returnVals, &returnVal)
	return returnVals, nil
}

func GoDivide(Cell1 *ListCell, Cell2 *ListCell)([]*ListCell, error){
	if Cell1.TypeName != Cell2.TypeName{
		err := fmt.Sprintf("Error: attempting to divide type %v by type %v, but these types are not compatible.\n", Cell2.TypeName, Cell1.TypeName)
		return nil, errors.New(err)
	}
	returnVals := make([]*ListCell, 0, 1)
	returnVal := ListCell{TypeName: Cell1.TypeName, Mutable: Cell1.Mutable}
	if val1, ok1 := Cell1.Value.(int); ok1 {
		if val2, ok2 := Cell2.Value.(int); ok2 {
			returnVal.Value = val1 / val2
		}else{
			err := fmt.Sprintf("Error: attempting to divide type %v by type %v, but the first really was an int and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int64); ok1 {
		if val2, ok2 := Cell2.Value.(int64); ok2 {
			returnVal.Value = val1 / val2
		}else{
			err := fmt.Sprintf("Error: attempting to divide type %v by type %v, but the first really was an int64 and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int32); ok1 {
		if val2, ok2 := Cell2.Value.(int32); ok2 {
			returnVal.Value = val1 / val2
		}else{
			err := fmt.Sprintf("Error: attempting to divide type %v by type %v, but the first really was an int32 and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(int16); ok1 {
		if val2, ok2 := Cell2.Value.(int16); ok2 {
			returnVal.Value = val1 / val2
		}else{
			err := fmt.Sprintf("Error: attempting to divide type %v by type %v, but the first really was an int16 and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(float64); ok1 {
		if val2, ok2 := Cell2.Value.(float64); ok2 {
			returnVal.Value = val1 / val2
		}else{
			err := fmt.Sprintf("Error: attempting to divide type %v by type %v, but the first really was a float64 and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	if val1, ok1 := Cell1.Value.(float32); ok1 {
		if val2, ok2 := Cell2.Value.(float32); ok2 {
			returnVal.Value = val1 / val2
		}else{
			err := fmt.Sprintf("Error: attempting to divide type %v by type %v, but the first really was a float32 and the second wasn't.\n", Cell2.TypeName, Cell1.TypeName)
			return nil, errors.New(err)
		}
	}
	returnVals = append(returnVals, &returnVal)
	return returnVals, nil
}

func GoIf(Cell1 *ListCell, Cell2 *ListCell)([]*ListCell, error){
	returnVals := make([]*ListCell, 0, 1)
	returnVal := new(ListCell)
	if Cell1.TypeName != "bool"{
		err := fmt.Sprintf("Error: expected a boolean value as first argument to if builtin, but got a %v.\n", Cell1.TypeName)
		return nil, errors.New(err)
	}
	if condVal, ok := Cell1.Value.(bool); ok {
		if listVal, ok2 := Cell2.Value.(CellList); ok2{			
			if condVal{
				returnVal = &listVal.Cells[0]
			}else{
				returnVal = &listVal.Cells[1]
			}
		}else{
			err := fmt.Sprintf("Error: second internal argument to if builtin was not actually a list.\n", Cell1.TypeName)
			return nil, errors.New(err)
		}
	}else{
		err := fmt.Sprintf("Error: first argument to if builtin appeared to be a bool but actually wasn't.\n", Cell1.TypeName)
		return nil, errors.New(err)
	}
	returnVals = append(returnVals, returnVal)
	return returnVals, nil
}

func EvalPrim(list []ListCell, env Environment)(*ListCell, error){
	if list[0].TypeName != FUNCTION_TYPE_NAME{
		err := fmt.Sprintf("Error: expected a function as first cell in list passed to eval builtin, but got a %v.\n", list[0].TypeName)
		return nil, errors.New(err)
	}else if list[0].TypeName == VAR_TYPE_NAME{
		if varName, ok := list[0].Value.(string); ok {
			binding := env.findBinding(varName, true, true)
			if binding == nil{
				err := fmt.Sprintf("Error: var in first cell in list passed to eval builtin, %v, is not bound.\n", varName)
				return nil, errors.New(err)
			}
		}else{
			err := fmt.Sprintf("Error: expected var in first cell of list passed to eval builtin, but it was actually not a var.\n")
			return nil, errors.New(err)
		}
		
	}
	return nil, nil
}

func GoEval(Cell1 *ListCell, Cell2 *ListCell)([]*ListCell, error){
	returnVals := make([]*ListCell, 0, 1)
	returnVal := new(ListCell)
	if Cell1.TypeName != FUNCTION_TYPE_NAME || Cell2.TypeName != ENVIRONMENT_TYPE_NAME{
		err := fmt.Sprintf("Error: expected a list and an environment as arguments to eval builtin, but got a %v and a %v.\n", Cell1.TypeName, Cell2.TypeName)
		return nil, errors.New(err)
	}
	if list, ok := Cell1.Value.([]ListCell); ok {
		if env, ok2 := Cell1.Value.(Environment); ok2 {
			returnValShad, err := EvalPrim(list,env)
			if err != nil{
				return nil, err
			}
			returnVal = returnValShad
		}else{
			err := fmt.Sprintf("Error: second argument to eval builtin claimed to an environment but wasn't.\n")
			return nil, errors.New(err)		
		}		
	}else{
		err := fmt.Sprintf("Error: first argument to eval builtin claimed to be a list but wasn't.\n")
		return nil, errors.New(err)
	}
	returnVals = append(returnVals, returnVal)
	return returnVals, nil
}
