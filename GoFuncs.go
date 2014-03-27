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

type goFuncType int

const (
	GoAddT goFuncType = iota
	GoSubtractT
	GoMultiplyT
	GoDivideT
)

type GoFunc func (Cell1 *ListCell, Cell2 *ListCell)(*ListCell, error)

func CallGoFunc(funcType goFuncType, parameters []*ListCell)(returnVals []*ListCell, error){
	switch funcType{
	case GoAddT:
		res, err := GoAdd(parameters[0], parameters[1])
		if err{
			return nil, err
		}else{
			return res, nil
		}
	case GoSubtractT:
		res, err := GoSubtract(parameters[0], parameters[1])
		if err{
			return nil, err
		}else{
			return res, nil
		}
	case GoMultiplyT:
		res, err := GoMultiply(parameters[0], parameters[1])
		if err{
			return nil, err
		}else{
			return res, nil
		}
	case GoDivideT:
		res, err := GoDivide(parameters[0], parameters[1])
		if err{
			return nil, err
		}else{
			return res, nil
		}
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
	returnVals append(returnVals, returnVal)
	return &returnVals, nil
}

func GoSubtract(Cell1 *ListCell, Cell2 *ListCell)([]*ListCell, error){
	if Cell1.TypeName != Cell2.TypeName{
		err := fmt.Sprintf("Error: attempting to subtract type %v from type %v, but these types are not compatible.\n", Cell2.TypeName, Cell1.TypeName)
		return nil, errors.New(err)
	}
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
	returnVals append(returnVals, returnVal)
	return returnVals, nil
}

func GoMultiply(Cell1 *ListCell, Cell2 *ListCell)([]*ListCell, error){
	if Cell1.TypeName != Cell2.TypeName{
		err := fmt.Sprintf("Error: attempting to multiply type %v by type %v, but these types are not compatible.\n", Cell1.TypeName, Cell2.TypeName)
		return nil, errors.New(err)
	}
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
	returnVals append(returnVals, returnVal)
	return &returnVals, nil
}

func GoDivide(Cell1 *ListCell, Cell2 *ListCell)([]*ListCell, error){
	if Cell1.TypeName != Cell2.TypeName{
		err := fmt.Sprintf("Error: attempting to divide type %v by type %v, but these types are not compatible.\n", Cell2.TypeName, Cell1.TypeName)
		return nil, errors.New(err)
	}
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
	returnVals append(returnVals, returnVal)
	return &returnVals, nil
}
