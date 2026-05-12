package assembler

import (
	"fmt"
	"strconv"
	"strings"
)

func getLabels(tokens []Token) map[string]int32 {
	// returning, a map of
	// identifier -> an address to memory.
	//
	// what about directives?
	idx := 0
	var line [] Token
	labels := make(map[string]int32)
	for _, t := range tokens {
		if t.Kind != NewLine { line = append(line, t); continue }

		if line[0].Kind == Identifier && line[1].Kind == Colon {
		   	labels[line[0].Lexeme] = int32(idx)
			// idx += 1
	   } else {
	   		idx += 1
	   }

		line = line[:0]
	}
	return labels
}

func Analyze(tokens []Token) ([]uint32, error) {
	var line [] Token
	// haveErr := false
	userLabels := getLabels(tokens)
	fmt.Println(userLabels)
	// directives should prob check for them just like labels, and
	// store the action that the program must load into memory once running vm

	hasErrors := false
	var mem []uint32
	for _, t := range tokens {
		if t.Kind != NewLine { line = append(line, t); continue }

		if isLabel(line) { line = line[:0] ; continue }

		op, err := getStructure(line, userLabels)
		if err != nil {
			fmt.Printf("[ERROR LINE %v]: %v\n", line[0].Line, err)
			printCodeLine(line)
			hasErrors = true
		}

		if !hasErrors {
			//mem = append(mem, Encode(op, userLabels))
			mem = append(mem, Encode(op))
		}

		// what is this doing?
		line = line[:0]
	}

	if hasErrors { return nil, fmt.Errorf("\nWont encode due to error") }
	return mem, nil
}

func printCodeLine(line []Token) {
	fmt.Print("-> ")
	// need to add more logic for it to be prettier.
	for _, l := range line {
		fmt.Printf("%v ", l.Lexeme)
	}
	fmt.Println()
}

func getStructure(line []Token, labels map[string]int32) (Op, error) {
	// caller checks for user defined lables
	if len(line) == 0 { return nil, fmt.Errorf("Empty instruction") }

	switch strings.ToLower(line[0].Lexeme) {
	case "halt":
		return checkHalt(line)
	case "mov":
		return checkMov(line, labels)
	case "add":
		return checkAdd(line)
	case "sub":
		return checkSub(line)
	case "cmp":
		return checkCmp(line)
	case "b":
		return checkB(line, labels)
	case "blt":
		return checkBlt(line, labels)
	case "beq":
		return checkBeq(line, labels)
	case "bgt":
		return checkBgt(line, labels)
	case "lsl":
		return checkLsl(line)
	case "lsr":
		return checkLsr(line)
	case "ldr":
		return checkLdr(line, labels)
	case "str":
		return checkStr(line, labels)
	}
	return nil, fmt.Errorf("error getting structure")
}

func isNumber(tok Token) bool {
	// should I also check if I can convert the lexeme to a actual value?
	// that can fit within the thing?
	return tok.Kind == Number || tok.Kind == HexNumber || tok.Kind == BitNumber
}

func checkHalt(line []Token) (Op, error) {
	if len(line) == 1 {
		return Opp{ Op: "halt", }, nil
	}

	return nil, fmt.Errorf("Error on halt instruction")
}

// mov r0, r1
// mov r0, #(number)
// mov r0, #msg // meaning mov r0, the beggining of msg assembly directive
func checkMov(line []Token, labels map[string]int32) (Op, error) {
	// TODO: negative numbers
	switch len(line) {
	case 4:
		if line[1].Kind == Register && line[2].Kind == Comma && line[3].Kind == Register {
			return Oprr{
				Op: "movrr",
				R1: RegisterToI8[lower(line[1].Lexeme)],
				R2: RegisterToI8[lower(line[3].Lexeme)],
			}, nil
		}

	case 5:
		if line[1].Kind == Register && line[2].Kind == Comma && line[3].Kind == Hash && (isNumber(line[4]) || line[4].Kind == Identifier) {
			// TODO: parse immediate or identifier?
			var imm int32
			if isNumber(line[4]) { imm = parseImm(line[4]) } else { panic("parse identifier within mov") }
			return Opri{
				Op: "movri",
				R1: RegisterToI8[lower(line[1].Lexeme)],
				I:  imm,
			}, nil
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
			return Oprrr{
				Op: "addrrr",
				R1: RegisterToI8[lower(line[1].Lexeme)],
				R2: RegisterToI8[lower(line[3].Lexeme)],
				R3: RegisterToI8[lower(line[5].Lexeme)],
			}, nil
		}
	case 7:
		if line[1].Kind == Register && line[2].Kind == Comma && line[3].Kind == Register && line[4].Kind == Comma && line[5].Kind == Hash && isNumber(line[6]) {
		   	// parse immediate
			var imm int32
			if isNumber(line[6]) { imm = parseImm(line[6]) } else { panic("parse identifier within add") }
		   	return Oprri{
		   		Op: "addrri",
		   		R1: RegisterToI8[lower(line[1].Lexeme)],
		   		R2: RegisterToI8[lower(line[3].Lexeme)],
		   		I: imm,
		   	}, nil
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
			return Oprrr{
				Op: "subrrr",
				R1: RegisterToI8[lower(line[1].Lexeme)],
				R2: RegisterToI8[lower(line[3].Lexeme)],
				R3: RegisterToI8[lower(line[5].Lexeme)],
			}, nil
		}
	case 7:
		if  line[1].Kind == Register && line[2].Kind == Comma && line[3].Kind == Register && line[4].Kind == Comma && line[5].Kind == Hash && isNumber(line[6]) {
			var imm int32
			if isNumber(line[6]) { imm = parseImm(line[6]) } else { panic("parse identifier within subrri") }		// parse immediate
			return Oprri{
				Op: "subrri",
				R1: RegisterToI8[lower(line[1].Lexeme)],
				R2: RegisterToI8[lower(line[3].Lexeme)],
				I: imm,
			}, nil
		}
	}

	return nil, fmt.Errorf("invalid sub instruction")
}

// cmp r0, r1
// cmp r0, #0b10
// cmp r0, #10
// cmp r0, #0xA
func checkCmp(line []Token) (Op, error) {
	// TODO negative numbers
	switch len(line) {
	case 4:
		if line[1].Kind == Register && line[2].Kind == Comma && line[3].Kind == Register {
			return Oprr{
				Op: "cmprr",
				R1: RegisterToI8[lower(line[1].Lexeme)],
				R2: RegisterToI8[lower(line[3].Lexeme)],
			}, nil
		}
	case 5:
		if line[1].Kind == Register && line[2].Kind == Comma && line[3].Kind == Hash && isNumber(line[4]) {
			imm := parseImm(line[4])
			return Opri{
				Op: "cmpri",
				R1: RegisterToI8[lower(line[1].Lexeme)],
				I: imm,
			}, nil
		}
	}
	return nil, fmt.Errorf("invalid cmp instruction")
}

func checkB(line []Token, labels map[string]int32) (Op, error) {
	if len(line) != 2 { return nil, fmt.Errorf("error on B instruction") }

	if line[1].Kind != Identifier { panic("how get here?") }
	_, ok := labels[line[1].Lexeme]
	if !ok { panic("This shouldn't be possible") }

	return Opl {
		Op: "bl",
		I : labels[line[1].Lexeme],
	}, nil
}

// blt label
func checkBlt(line []Token, labels map[string]int32) (Op, error) {
	if len(line) != 2 { return nil, fmt.Errorf("Error on blt instruction") }

	if line[1].Kind != Identifier { panic("how get here?") }
	_, ok := labels[line[1].Lexeme]
	if !ok { panic("This shouldn't be possible") }

	return Opl{
		Op: "bltl",
		I : labels[line[1].Lexeme],
	}, nil
}

// beq label
func checkBeq(line []Token, labels map[string]int32) (Op, error) {
	switch len(line) {
	case 2:
		if line[1].Kind != Identifier { panic("how get here?") }
		_, ok := labels[line[1].Lexeme]
		if !ok { panic("This shouldn't be possible") }
		return Opl {
			Op: "beql",
			I : labels[line[1].Lexeme],
		}, nil

	}

	return nil, fmt.Errorf("Error on beq instruction")
}

func checkBgt(line []Token, labels map[string]int32) (Op, error) {
	switch len(line){
	case 2:
		if line[1].Kind != Identifier { panic("how we get here?") }
		_, ok := labels[line[1].Lexeme]
		if !ok { panic("This shouldn't be possible") }
		return Opl {
			Op: "bgtl",
			I : labels[line[1].Lexeme],
		}, nil
	}
	return nil, fmt.Errorf("Error on bgt instruction")
}

// lsl r0, r0, #1
// lsl r0, r0, r1
func checkLsl(line []Token) (Op, error) {
	if len(line) != 6 && len(line) != 7 { return nil, fmt.Errorf("wrong format for lsl") }
	if len(line) == 6 {
		if line[1].Kind != Register && line[2].Kind != Comma && line[3].Kind != Register && line[4].Kind != Comma && line[5].Kind == Register {
			return nil, fmt.Errorf("format for (lsl): lsl rd, rs1, rs2")
		}
		return Oprrr {
			Op: "lslrrr",
			R1: RegisterToI8[lower(line[1].Lexeme)],
			R2: RegisterToI8[lower(line[3].Lexeme)],
			R3: RegisterToI8[lower(line[5].Lexeme)],
		}, nil
	}
	// len line == 7
	if  line[1].Kind == Register && line[2].Kind == Comma && line[3].Kind == Register && line[4].Kind == Comma && line[5].Kind == Hash && isNumber(line[6]) {
		var imm int32
		if isNumber(line[6]) { imm = parseImm(line[6]) } else { panic("parse identifier within lsl rd, rs1, #num") }		// parse immediate
		return Oprri{
			Op: "lslrri",
			R1: RegisterToI8[lower(line[1].Lexeme)],
			R2: RegisterToI8[lower(line[3].Lexeme)],
			I: imm,
		}, nil
	}
	return nil, fmt.Errorf("This should be reached?")
}
// lsr r0, r0, #1
// lsr r0, r0, r1
func checkLsr(line []Token) (Op, error) {
	if len(line) != 6 && len(line) != 7 { return nil, fmt.Errorf("wrong format for lsr") }
	if len(line) == 6 {
		if line[1].Kind != Register && line[2].Kind != Comma && line[3].Kind != Register && line[4].Kind != Comma && line[5].Kind == Register {
			return nil, fmt.Errorf("format for (lsr): lsl rd, rs1, rs2")
		}
		return Oprrr {
			Op: "lsrrrr",
			R1: RegisterToI8[lower(line[1].Lexeme)],
			R2: RegisterToI8[lower(line[3].Lexeme)],
			R3: RegisterToI8[lower(line[5].Lexeme)],
		}, nil
	}
	// len line == 7
	if  line[1].Kind == Register && line[2].Kind == Comma && line[3].Kind == Register && line[4].Kind == Comma && line[5].Kind == Hash && isNumber(line[6]) {
		var imm int32
		if isNumber(line[6]) { imm = parseImm(line[6]) } else { panic("parse identifier within lsr rd, rs1, #num") }		// parse immediate
		return Oprri{
			Op: "lsrrri",
			R1: RegisterToI8[lower(line[1].Lexeme)],
			R2: RegisterToI8[lower(line[3].Lexeme)],
			I: imm,
		}, nil
	}
	return nil, fmt.Errorf("This should be reached?")
}

// Direct addressing
// (ldr|str) rd, (bit|dec|hex| label)
// Indirect addressing
// (ldr|str) rd, [reg]
func checkLdr(line []Token, labels map[string]int32) (Op, error) {
	if len(line) != 4 && len(line) != 6 { return nil, fmt.Errorf("Error con formating ldr") }
	// direct addressing
	if line[1].Kind == Register && line[2].Kind == Comma {
		r1 := RegisterToI8[lower(line[1].Lexeme)]
		var i int32
		switch line[3].Kind {
		case BitNumber, Number, HexNumber:
			i = parseImm(line[3])
		case Identifier:
			v, in := labels[line[3].Lexeme]
			if in == false { return nil, fmt.Errorf("Label not recognized") }
			i = v
		}
		return Opri {
			Op: "ldrDirect",
			R1: r1,
			I: i,
		}, nil
	}

	// indirect addressing
	if line[1].Kind == Register && line[2].Kind == Comma && line[3].Kind == LeftBracket && line[4].Kind == Register && line[5].Kind == RightBracket {
		r1 := RegisterToI8[lower(line[1].Lexeme)]
		r2 := RegisterToI8[lower(line[4].Lexeme)]
		return Oprr {
			Op: "ldrIndirect",
			R1: r1,
			R2: r2,
		}, nil
	}
	return nil, fmt.Errorf("Error on structure")
}

// Direct addressing
// (ldr|str) rd, (bit|dec|hex| label)
func checkStr(line []Token, labels map[string]int32) (Op, error) {
	if len(line) != 4 && len(line) != 6 { return nil, fmt.Errorf("Error con formating str")}
	// direct addressing
	if len(line) == 4 && line[1].Kind == Register && line[2].Kind == Comma {
		r1 := RegisterToI8[lower(line[1].Lexeme)]
		var i int32
		switch line[3].Kind {
		case BitNumber, Number, HexNumber:
			i = parseImm(line[3])
		case Identifier:
			v, in := labels[line[3].Lexeme]
			if in == false { return nil, fmt.Errorf("Label not recognized") }
			i = v
		}
		return Opri {
			Op: "strDirect",
			R1: r1,
			I: i,
		}, nil
	}
	// indirect addressing
	if len(line) == 6 && line[1].Kind == Register && line[2].Kind == Comma && line[3].Kind == LeftBracket && line[4].Kind == Register && line[5].Kind == RightBracket {
		r1 := RegisterToI8[lower(line[1].Lexeme)]
		r2 := RegisterToI8[lower(line[4].Lexeme)]
		return Oprr {
			Op: "strIndirect",
			R1: r1,
			R2: r2,
		}, nil
	}
	return nil, fmt.Errorf("Error on structure")
}
/*
// address must be divisable by 4, (forgor about it.)
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
*/

func lower(s string) string {
	return strings.ToLower(s)
}

func parseImm(tok Token) int32 {
	var base int
	var value string

	switch tok.Kind {
	case BitNumber:
		base = 2
		value = tok.Lexeme[2:] // strip "0b"
	case HexNumber:
		base = 16
		value = tok.Lexeme[2:] // strip "0x"
	case Number:
		base = 10
		value = tok.Lexeme
	default:
		panic("unexpected token type for immediate (label?)")
	}

	n, err := strconv.ParseInt(value, base, 32)
	if err != nil {
		panic("invalid immediate: " + tok.Lexeme)
	}

	return int32(n)
}

func isLabel(line []Token) bool {
	return len(line) == 2 && line[0].Kind == Identifier && line[1].Kind == Colon
}
