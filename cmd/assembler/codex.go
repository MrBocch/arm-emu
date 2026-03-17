package assembler

import (
	"fmt"
	"strconv"
)

// go is shit for not having sum types
type Op interface {
	isOp()
}

// halt 
type Opp struct {
	Op string
}
func (Opp) isOp() {}
func (o Opp) String() string {
	return fmt.Sprintf("%s", o.Op)
}

// OP, R, R
type Oprr struct { 
	Op string 
	R1 string 
	R2 string 
}
func (Oprr) isOp() {}
func (o Oprr) String() string {
	return fmt.Sprintf("%s %s, %s", o.Op, o.R1, o.R2)
}


// OP R, I
type Opri struct {
	Op string
	R1 string
	I  int32
}
func (Opri) isOp() {}
func (o Opri) String() string {
	return fmt.Sprintf("%s %s, #%d", o.Op, o.R1, o.I)
}


// OP R, R, I
type Oprri struct {
	Op string
	R1 string
	R2 string
	I  int32
}
func (Oprri) isOp() {}
func (o Oprri) String() string {
	return fmt.Sprintf("%s %s, %s, #%d", o.Op, o.R1, o.R2, o.I)
}

// OP R, R, R
type Oprrr struct {
	Op string
	R1 string
	R2 string
	R3 string
}
func (Oprrr) isOp() {}
func (o Oprrr) String() string {
	return fmt.Sprintf("%s %s, %s, %s", o.Op, o.R1, o.R2, o.R3)
}


func opToS(op Op) string {
	switch op.(type) {
	case Opp: return ""
	case Oprr: return "rr"
	case Opri: return "ri"
	case Oprrr: return "rrr"
	}
	return ""
}

func flipMap(m map[string]string) map[string]string {
	fm := make(map[string]string)
	for k, v := range m {
		fm[v] = k
	}

	return fm
}

var registerToB = map[string]string {
	"r0": "0000",
	"r1": "0001",
	"r2": "0010",
	"r3": "0011",
	"r4": "0100",
	"r5": "0101",
	"r6": "0110",
	"r7": "0111",
	"r8": "1000",
	"r9": "1001",
	"r10":"1010",
	"r11":"1011",
	"r12":"1100",
	"sp": "1101",
	"lr": "1110",
	"pc": "1111",
}

var bToRegister = flipMap(registerToB)

// built in identifiers .WriteString

// should i have used a addition field
// for knowing instruction type?
// <type> <op>
// type ::= OP | OP i | OP r i | OP r i i | OP r r | OP r r i

// 8 bit long instructions?
// op
// op i
// op ri (register, imediate)
// op rr (register, rr)
var opToB = map[string]string {
	"halt":  "00000000",
	"movrr": "00000001",
	"movri": "00000010",
	"addrrr":"00000011",
	"addrri":"00000100",
	"subrrr":"00000101",
	"subrri":"00000110",
	"cmprr" :"00000111",
	"cmpri" :"00001000",
}

var bToOp = flipMap(opToB)

// i think it was a mistake to encode as a string
// need to encode as a int32
func Encode(op Op, labels map[string]int) string {
	// TODO encode labels 
	bs := ""
	switch v := op.(type) {
	case Opp: return padding(opToB[v.Op])
	case Oprr: return padding(opToB[v.Op + "rr"] + registerToB[v.R1] + registerToB[v.R2])
	case Opri: return padding(opToB[v.Op + "ri"] + registerToB[v.R1] + iToB20(v.I))
	case Oprrr: return padding(opToB[v.Op + "rrr"] + registerToB[v.R1] + registerToB[v.R2] + registerToB[v.R3])
	case Oprri: return padding(opToB[v.Op + "rri"] + registerToB[v.R1] + registerToB[v.R2] + iToB16(v.I))
	}
	return bs 
}

func Decode(s string) (Op, error) {
    // Ensure we have exactly 32 bits
    if len(s) != 32 {
		panic("Did not receive 32bit op")
    }
    
    // Extract opcode (first 8 bits)
    opcode := s[:8]
    
    // Find operation name
    opName, exists := bToOp[opcode]
    if !exists {
    	panic(fmt.Errorf("Dont recognize [%s] instruction", opcode))
    }
    
    // Decode based on operation type
    switch opName {
    case "halt":
        return Opp{Op: "halt"}, nil
        
    case "movrr":
        // Format: 8-bit op + 4-bit r1 + 4-bit r2
        r1 := bToRegister[s[8:12]]
        r2 := bToRegister[s[12:16]]
        return Oprr{Op: opName, R1: r1, R2: r2}, nil
        
    case "movri":
        // Format: 8-bit op + 4-bit r1 + 20-bit immediate
        r1 := bToRegister[s[8:12]]
        i, _ := strconv.ParseInt(s[12:32], 2, 32)
        return Opri{Op: opName, R1: r1, I: int32(i)}, nil
        
    case "addrrr", "subrrr":
        // Format: 8-bit op + 4-bit r1 + 4-bit r2 + 4-bit r3
        // Note: remaining bits might be unused or for future expansion
        r1 := bToRegister[s[8:12]]
        r2 := bToRegister[s[12:16]]
        r3 := bToRegister[s[16:20]]
        return Oprrr{Op: opName, R1: r1, R2: r2, R3: r3}, nil
        
    case "addrri", "subrri":
        // Format: 8-bit op + 4-bit r1 + 4-bit r2 + 16-bit immediate
        r1 := bToRegister[s[8:12]]
        r2 := bToRegister[s[12:16]]
        i, _ := strconv.ParseInt(s[16:32], 2, 32)
        return Oprri{Op: opName, R1: r1, R2: r2, I: int32(i)}, nil
        
    case "cmprr":
        // Format: 8-bit op + 4-bit r1 + 4-bit r2
        r1 := bToRegister[s[8:12]]
        r2 := bToRegister[s[12:16]]
        return Oprr{Op: opName, R1: r1, R2: r2}, nil
        
    case "cmpri":
        // Format: 8-bit op + 4-bit r1 + 20-bit immediate
        r1 := bToRegister[s[8:12]]
        i, _ := strconv.ParseInt(s[12:32], 2, 32)
        return Opri{Op: opName, R1: r1, I: int32(i)}, nil
    }
    
    return nil, fmt.Errorf("What?")
}

func padding(s string) string {
	if len(s) == 32 { return s }
	if len(s) > 32 { panic("length of instruction is too big")}

	for ; len(s) < 32 ; {
		s += "0"
	}

	return s 
}

func iToB20(n int32) string {
	const (
		width = 20
		max   = (1 << (width - 1)) - 1  //  524287
		min   = -(1 << (width - 1))     // -524288
	)

	if n > max || n < min {
		panic("immediate out of 20-bit signed range")
	}

	// mask to 20 bits (two's complement)
	u := uint32(n) & ((1 << width) - 1)

	return fmt.Sprintf("%020b", u)
}

func iToB16(n int32) string {
	const (
		width = 16
		max   = (1 << (width - 1)) - 1  //  32767
		min   = -(1 << (width - 1))     // -32768
	)

	if n > max || n < min {
		panic("immediate out of 16-bit signed range")
	}

	// mask to 16 bits (two's complement)
	u := uint32(n) & ((1 << width) - 1)

	return fmt.Sprintf("%016b", u)
}
