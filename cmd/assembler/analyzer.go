package assembler

import (
    "fmt"
    "strings"
)

func Analyze(tokens []Token) {
	var line [] Token
	for _, t := range tokens {
		if t.Kind != NewLine { line = append(line, t); continue }

		fmt.Printf("Analyze %v -> %v\n", line, checkSyntax(line))
		// do something for real.
		if !checkSyntax(line) {
			fmt.Printf("[Error] line %v\n", line[0].Line)
		} 
		line = line[:0]
	}
}

func checkSyntax(line []Token) bool {
	if len(line) == 0 { return false }
	switch strings.ToLower(line[0].Lexeme) {
	case "mov":
		return checkMov(line) 
		 
	}

	return false 
}

func isNumber(tok Token) bool {
	// should I also check if I can convert the lexeme to a actual value? 
	return tok.Kind == Number || tok.Kind == HexNumber || tok.Kind == BitNumber
}

func checkMov(line []Token) bool {
	switch len(line) {
	case 4:
		return line[1].Kind == Register &&
			   line[2].Kind == Comma &&
			   line[3].Kind == Register 
	case 5:
		return line[1].Kind == Register &&
			   line[2].Kind == Comma &&
			   line[3].Kind == Hash &&
			   isNumber(line[4])
	default:
		return false 
	}
}
