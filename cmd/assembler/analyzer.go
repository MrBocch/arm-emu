package assembler

import (
    "fmt"
    "strings"
)

func getLabels(tokens []Token) map[string]int {
	// what about directives? 
	idx := 0
	var line [] Token
	labels := make(map[string]int)
	for _, t := range tokens {
		if t.Kind != NewLine { line = append(line, t); continue }

		if line[0].Kind == Identifier &&
		   line[1].Kind == Colon {
		   	labels[line[0].Lexeme] = idx
	   } else {
	   		idx += 4
	   }

		line = line[:0]
	}
	return labels
}

func Analyze(tokens []Token) {
	var line [] Token
	// haveErr := false 
	// userLabels := getLabels(tokens)
	// fmt.Println(userLabels)
	// directives should prob check for them just like labels, and
	// store the action that the program must load into memory once running vm 

	for _, t := range tokens {
		if t.Kind != NewLine { line = append(line, t); continue }

		op, err := getStructure(line)
		if err != nil {
			fmt.Printf("[ERROR] %v", line[0].Line)
			fmt.Println(err)
		}
		fmt.Println(op)
		//fmt.Printf("Analyze %v -> %v\n", line, checkSyntax(line))
		// do something for real.
		// if !checkSyntax(line) {
			// printError(line)
			// haveErr = true 
		// } 
		// if !haveErr {
			// encode
			// Encode()
		// }

		
		line = line[:0]
	}

}

func getStructure(line []Token) (Op, error) {
	if len(line) == 0 { return nil, fmt.Errorf("Empty instruction") }
	
	switch strings.ToLower(line[0].Lexeme) {
	case "mov":
		return checkMov(line) 
	case "add":
		return checkAdd(line)
	case "sub":
		return checkSub(line)
	case "halt":
		return checkHalt(line)
	}
	return nil, fmt.Errorf("invalid instruction instruction")
	/*
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
	case "ldr":
		return checkLdr(line)
	case "and":
		return checkAnd(line)
	case "str":
		return checkStr(line)

	default:
		// check for identifier
		return checkIdentifier(line)
	}
	return false 
	*/
}

func isNumber(tok Token) bool {
	// should I also check if I can convert the lexeme to a actual value? 
	return tok.Kind == Number || tok.Kind == HexNumber || tok.Kind == BitNumber
}


// mov r0, r1
// mov r0, #(number)
// mov r0, #msg // meaning mov r0, the beggining of msg assembly directive
func checkMov(line []Token) (Op, error) {
	// TODO: negative numbers
	switch len(line) {
	case 4:
		if line[1].Kind == Register &&
			line[2].Kind == Comma &&
			line[3].Kind == Register {

			return Oprr{ op: "mov", r1: lower(line[1].Lexeme), r2: lower(line[3].Lexeme),}, nil
		}

	case 5:
		if line[1].Kind == Register &&
			line[2].Kind == Comma &&
			line[3].Kind == Hash &&
			(isNumber(line[4]) || line[4].Kind == Identifier) {
			// TODO: parse immediate
			return Opri{ op: "mov", r1: lower(line[1].Lexeme), i:  0,}, nil
		}
	}

	return nil, fmt.Errorf("invalid mov instruction")
}

// add r0, r1, r2
// add r0, r1, #0b11
// add r0, r1, #3
// add r0, r1, #0xA
func checkAdd(line []Token) (Op, error) {
	// TODO negative numbers 
	switch len(line){
	case 6:
		if line[1].Kind == Register && line[2].Kind == Comma && line[3].Kind == Register && line[4].Kind == Comma && line[5].Kind == Register {
			return Oprrr{ op: "add", r1: lower(line[1].Lexeme), r2: lower(line[3].Lexeme), r3: lower(line[5].Lexeme),}, nil
		}
	case 7:
		if line[1].Kind == Register && line[2].Kind == Comma && line[3].Kind == Register && line[4].Kind == Comma && line[5].Kind == Hash && isNumber(line[6]) {
		   	// parse immediate
		   	return Oprri{ op: "add", r1: lower(line[1].Lexeme), r2: lower(line[3].Lexeme), i: 0,}, nil
		   }
	}
	return nil, fmt.Errorf("invalid add instruction")
}

// sub r0, r1, r2
// sub r0, r1, #(number)
func checkSub(line []Token) (Op, error) {
	// TODO negative numbers 
	switch len(line) {
	case 6:
		if line[1].Kind == Register && line[2].Kind == Comma && line[3].Kind == Register && line[4].Kind == Comma && line[5].Kind == Register {
			return Oprrr{ op: "sub", r1: lower(line[1].Lexeme), r2: lower(line[3].Lexeme), r3: lower(line[5].Lexeme),}, nil
		}
	case 7:
		if  line[1].Kind == Register && line[2].Kind == Comma && line[3].Kind == Register && line[4].Kind == Comma && line[5].Kind == Hash && isNumber(line[6]) {
			// parse immediate
			return Oprri{ op: "sub", r1: lower(line[1].Lexeme), r2: lower(line[3].Lexeme), i: 0,}, nil
		}
	}

	return nil, fmt.Errorf("invalid sub instruction")
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
func checkHalt(line []Token) (Op, error) {
	if len(line) == 1 && lower(line[0].Lexeme) == "halt" {
		return Opp{ op: "halt", }, nil
	}
	return nil, fmt.Errorf("invalid halt instruction")
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

// bne label
func checkBne(line []Token) bool {
	// TODO check if its an actuall user defined label
	return len(line) == 2 && line[1].Kind == Identifier
}


// and r0, r1, r2
// and r0, r1, #(number)
func checkAnd(line []Token) bool {
	switch len(line){
	case 6:
		return line[1].Kind == Register &&
			   line[2].Kind == Comma &&
			   line[3].Kind == Register &&
			   line[4].Kind == Comma &&
			   line[5].Kind == Register
	case 7:
		return line[1].Kind == Register &&
		   line[2].Kind == Comma &&
		   line[3].Kind == Register &&
		   line[4].Kind == Comma &&
		   line[5].Kind == Hash &&
		   isNumber(line[6])

	}
	return false
}

// address must be divisable by 4
// str r0, .Thing
// str r0, number 
// str r0, [r1]
// str r0, [r1,r2]
// str r0, [r1 + r2]
// str r0, [r1 - r2]
func checkStr(line []Token) bool {
	switch len(line) {
	case 4:
		return line[1].Kind == Register && line[2].Kind == Comma && (isNumber(line[3]) || line[3].Kind == Identifier)
	case 6:
		return line[1].Kind == Register && line[2].Kind == Comma && line[3].Kind == LeftBracket && line[4].Kind == Register && line[5].Kind == RightBracket
	case 8:
		return line[1].Kind == Register &&
			   line[2].Kind == Comma &&
			   line[3].Kind == LeftBracket &&
			   line[4].Kind == Register &&
			   (line[5].Kind == Comma || line[5].Kind == Plus || line[5].Kind == Minus) &&
			   line[6].Kind == Register &&
			   line[7].Kind == RightBracket
	}
	return false 
}

// address must be divisable by 4
// ldr r0, .Thing
// ldr r0, number 
// ldr r0, [r1]
// ldr r0, [r1 , r2]
// ldr r0, [r1 + r2]
// ldr r0, [r1 - r2]
func checkLdr(line []Token) bool {
	switch len(line) {
	case 4:
		return line[1].Kind == Register && line[2].Kind == Comma && (isNumber(line[3]) || line[3].Kind == Identifier)
	// case 5: return line[1].Kind == Register && line[2].Kind == Comma && line[3].Kind == Identifier
	case 6:
		return line[1].Kind == Register && line[2].Kind == Comma && line[3].Kind == LeftBracket && line[4].Kind == Register && line[5].Kind == RightBracket
	case 8:
		return line[1].Kind == Register &&
			   line[2].Kind == Comma &&
			   line[3].Kind == LeftBracket &&
			   line[4].Kind == Register &&
			   (line[5].Kind == Comma || line[5].Kind == Plus || line[5].Kind == Minus) &&
			   line[6].Kind == Register &&
			   line[7].Kind == RightBracket
	}
	return false
}

// userIdentifier:
func checkIdentifier(line []Token) bool {
	// TODO check for user defined/built in identifiers
	// what about assembly directives? 
	return line[0].Kind == Identifier && line[1].Kind == Colon
}


func printError(line []Token) {
	// TODO: how about more helpful error messages? 
	fmt.Printf("[ERROR]\n[LINE: %v] ", line[0].Line)
	for _, t := range line {
		fmt.Printf("%v ", t.Lexeme)
	}
	fmt.Println()
}

func lower(s string) string {
	return strings.ToLower(s)
}
