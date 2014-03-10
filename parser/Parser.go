package Parser

import( 
	"errors"
	"fmt"
	"strings"
	"unicode"
)
const NEW_LINE = "\n"

type tokenType int
const(
	NullToken  tokenType = iota
	ListToken
	IdToken
	DefToken
	NumToken
	TypeDefToken
)

type numType int
const(
	FixNum numType = iota
	FloNum
)

type Token struct{
	Type tokenType
	NumType numType
	Value string
	ListVals []Token
	LineNum int
}

func Lex(input *string) []string{
	*input = strings.Replace(*input, "(", " ( ", -1)
	*input = strings.Replace(*input, ")", " ) ",-1)
	*input = strings.Replace(*input, "\n\r", "\n",-1)
	*input = strings.Replace(*input, "\r", "\n",-1)
	*input = strings.Replace(*input, "\n", " " + NEW_LINE + " ",-1)
	return (strings.Split(*input, " "))
}

func findMatchingParenDist(lexemes []string)(int,error){
	netParens := 1
	for i, lexeme := range lexemes{
		if lexeme == "("{
			netParens += 1
		}else if lexeme == ")"{
			netParens -= 1
		}
		if netParens == 0{
			return i, nil
		}
	}
	return -1,errors.New("Failed to find matching parenthesis!")
}

func numToToken(number string)(Token,error){
	numDots := 0
	for _, dig := range number{
		if !(unicode.IsDigit(dig)) && dig != '.'{
			return Token{Type: NullToken}, errors.New("Number is malformed; contains a non-digit!")
		} 
		if dig == '.'{
			numDots+= 1;
			if numDots > 1{
				return Token{Type: NullToken}, errors.New("Number is malformed; contains too many periods!")
			}
		}
	}
	if numDots > 1{
		return Token{Type: NumToken, NumType: FloNum, Value: number}, nil
	}else{
		return Token{Type: NumToken, NumType: FixNum, Value: number}, nil
	}
}

func strToToken(id string)(Token,error){
	for _, _ = range id{
	}
	strings.TrimSpace(id)
	if id == "let" || id == "letm" || id == "def" || id == "defm"{
		return Token{Type: DefToken, Value: id}, nil
	}else if id == ":"{
		return Token{Type: TypeDefToken, Value: id}, nil
	}else {
		return Token{Type: IdToken, Value: id}, nil
	}
}

func ParseList(lexemes []string, initLine int)Token{
	fmt.Printf("Parsing list: %v\n", lexemes)
	list := Token{Type: ListToken, ListVals: make([]Token,0,100)}
	lineNum := initLine
	for i := 0; i < len(lexemes); i++{
		lexeme := lexemes[i]
	//	fmt.Println(lexeme)
		newToken := Token{Type: NullToken}
		if lexeme == NEW_LINE {
			lineNum++
		}else if lexeme == "("{
			nextParemDist, err := findMatchingParenDist(lexemes[i+1:])
			if err != nil{
				fmt.Printf("Could not find matching right parenthesis for left parenthesis at line %v\n",
					len(strings.Split(strings.Join(lexemes,""),"\n") )) 
				panic("Unable to continue.\n")
			}
			newToken = ParseList(lexemes[i+1:nextParemDist+i+1], lineNum)
			list.ListVals = append(list.ListVals, newToken)
			i += nextParemDist+1
			continue
		}else{
			runes := []rune(lexeme)
			if len(runes) > 0{
				if unicode.IsDigit(runes[0]){
					token, err := numToToken(lexeme)
					if err != nil{
						fmt.Printf("Error: malformed literal at line %v; %v\n",
							len(strings.Split(strings.Join(lexemes,""),"\n")),err) 
						panic("Unable to continue.\n")
					}
					newToken = token
				}else{
					token, err := strToToken(lexeme)
					if err != nil{
						fmt.Printf("Error: malformed identifier at line %v; %v\n",
							len(strings.Split(strings.Join(lexemes,""),"\n")),err) 
						panic("Unable to continue.\n")
					}
					newToken = token
				}
			}
		}
		if newToken.Type != NullToken{
			newToken.LineNum = lineNum
			list.ListVals = append(list.ListVals, newToken)
		}
	}
	return list
}
