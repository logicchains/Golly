package main

import( 
	"fmt"
	"golly/parser"
)

type baseType int
const(
	Int baseType = iota
	Float
	Char
	Symbol
	List
	FuncDef
)

type FuncParem struct{
	Name string
	TypeID int64
}

type FunctionDef struct{
	Name string
	NumParems int16
	Parems []FuncParem
}

type ListCell struct{
	TypeID int64
	Value interface{}
	Mutable bool
}

type CellList struct{
	Cells []ListCell
	Environment []ListCell
}

func evalList(list Parser.Token){
	firstVal := &list.ListVals[0]
	switch firstVal.Type{
	case Parser.NumToken : 
		errMsg := fmt.Sprintf("Error: attempting to evaluate a literal, %v, at line %v.\n",firstVal.Value, firstVal.LineNum) 
		panic(errMsg)
	} 
}

func main(){
	//types := []string{"Int", "Float", "Char", "Symbol", "List"}
	input := `(+ 3 lenny)`
	res := Parser.Lex(input)
	tokens := Parser.ParseList(res, 0)
	for _, tok := range tokens.ListVals{
		fmt.Println(tok)
	}
	evalList(tokens.ListVals[0])
}
