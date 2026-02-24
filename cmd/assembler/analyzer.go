package assembler

import (
    "fmt"
    "strings"
)

func Analyze(tokens []Token) {
	var line [] Token
	for _, t := range tokens {
		if t.Kind != NewLine { line = append(line, t); continue }

		//fmt.Printf("Analyze %v -> %v\n", line, checkSyntax(line))
		// do something for real.
		if !checkSyntax(line) {
			printError(line)
		} else {
			fmt.Printf("Encoding: %v\n", line)
		}
		line = line[:0]
	}
}

func checkSyntax(line []Token) bool {
	if len(line) == 0 { return false }
	switch strings.ToLower(line[0].Lexeme) {
	case "mov":
		return checkMov(line) 
	case "add":
		return checkAdd(line)
	case "sub":
		return checkSub(line)
	case "halt":
		return checkHalt(line)
	case "cmp":
		return checkCmp(line)
	case "b":
		return checkB(line)
	case "blt":
		return checkBlt(line)
	case "beq":
		return checkBeq(line)
	case "bgt":
		return checkBgt(line)
	case "bne":
		return checkBne(line)

	}
	return false 
}

func isNumber(tok Token) bool {
	// should I also check if I can convert the lexeme to a actual value? 
	return tok.Kind == Number || tok.Kind == HexNumber || tok.Kind == BitNumber
}


// mov r0, r1
// mov r0, #0b101
// mov r0, #1
// mov r0, #0x10
func checkMov(line []Token) bool {
	// TODO negative numbers 
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


// add r0, r1, r2
// add r0, #0b11
// add r0, #3
// add r0, #0xA
func checkAdd(line []Token) bool {
	// TODO negative numbers 
	switch len(line){
	case 6:
		return line[1].Kind == Register &&
			   line[2].Kind == Comma &&
			   line[3].Kind == Register &&
			   line[4].Kind == Comma &&
			   line[5].Kind == Register
	case 5:
		return line[1].Kind == Register &&
			   line[2].Kind == Comma &&
			   line[3].Kind == Hash &&
			   isNumber(line[4])
	default:
		return false
	}
}

// sub r0, r1, r2
// sub r0, #0b01
// sub r0, #2
// sub r0, #0xA
func checkSub(line []Token) bool {
	// TODO negative numbers 
	switch len(line) {
	case 6:
		return line[1].Kind == Register &&
			   line[2].Kind == Comma &&
			   line[3].Kind == Register &&
			   line[4].Kind == Comma &&
			   line[5].Kind == Register
	case 5:
		return line[1].Kind == Register &&
			   line[2].Kind == Comma &&
			   line[3].Kind == Hash &&
			   isNumber(line[4])
	default:
		return false
	}

}

// cmp r0, r1
// cmp r0, #0b10
// cmp r0, #10
// cmp r0, #0xA
func checkCmp(line []Token) bool {
	// TODO negative numbers
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

// halt 
func checkHalt(line []Token) bool {
	return len(line) == 1
}

// b label
func checkB(line []Token) bool {
	// TODO check if its an actuall user defined label
	return len(line) == 2 && line[1].Kind == Identifier
}

// blt label
func checkBlt(line []Token) bool {
	// TODO check if its an actuall user defined label
	return len(line) == 2 && line[1].Kind == Identifier
}

// beq label
func checkBeq(line []Token) bool {
	// TODO check if its an actuall user defined label
	return len(line) == 2 && line[1].Kind == Identifier
}

// beq label
func checkBgt(line []Token) bool {
	// TODO check if its an actuall user defined label
	return len(line) == 2 && line[1].Kind == Identifier
}

// beq label
func checkNeq(line []Token) bool {
	// TODO check if its an actuall user defined label
	return len(line) == 2 && line[1].Kind == Identifier
}

func printError(line []Token) {
	// TODO: how about more helpful error messages? 
	fmt.Printf("[ERROR]\n[LINE: %v] ", line[0].Line)
	for _, t := range line {
		fmt.Printf("%v ", t.Lexeme)
	}
	fmt.Println()
}
