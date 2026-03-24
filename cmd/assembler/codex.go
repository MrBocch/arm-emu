package assembler

import (
	"fmt"
)

// go is shit for not having sum types
type Op interface {
	isOp()
}

// halt
// should i really be storing Op as a string? i wont know what 0000-0001 or 1000-1000 would mean
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
	R1 uint8
	R2 uint8
}
func (Oprr) isOp() {}
func (o Oprr) String() string {
	return fmt.Sprintf("%s %d, %d", o.Op, o.R1, o.R2)
}


// OP R, I
type Opri struct {
	Op string
	R1 uint8
	I  int32
}
func (Opri) isOp() {}
func (o Opri) String() string {
	return fmt.Sprintf("%s %d, #%d", o.Op, o.R1, o.I)
}


// OP R, R, I
type Oprri struct {
	Op string
	R1 uint8
	R2 uint8
	I  int32
}
func (Oprri) isOp() {}
func (o Oprri) String() string {
	return fmt.Sprintf("%s %d, %d, #%d", o.Op, o.R1, o.R2, o.I)
}

// OP R, R, R
type Oprrr struct {
	Op string
	R1 uint8
	R2 uint8
	R3 uint8
}
func (Oprrr) isOp() {}
func (o Oprrr) String() string {
	return fmt.Sprintf("%s %d, %d, %d", o.Op, o.R1, o.R2, o.R3)
}


func opToS(op Op) string {
	switch op.(type) {
	case Opp: return ""
	case Oprr: return "rr"
	case Opri: return "ri"
	case Oprrr: return "rrr"
	}
	panic("something is wrong here")
	//return ""
}

/*func flipMap(m map[string]string) map[string]uint {
	fm := make(map[string]string)
	for k, v := range m {
		fm[v] = k
	}

	return fm
	}*/

var RegisterToI8 = map[string]uint8 {
	"r0": 0,
	"r1": 1,
	"r2": 2,
	"r3": 3,
	"r4": 4,
	"r5": 5,
	"r6": 6,
	"r7": 7,
	"r8": 8,
	"r9": 9,
	"r10":10,
	"r11":11,
	"r12":12,
	"sp": 13,
	"lr": 14,
	"pc": 15,
}
var IToRegister = [16]string {
	"r0", "r1", "r2", "r3", "r4", "r5", "r6",
	"r7", "r8", "r9", "r10", "r11", "r12",
	"sp", "lr", "pc",
}

//var bToRegister = flipMap(registerToB)

// built in identifiers .WriteString

// should i have used a addition field
// for knowing instruction type?
// <type> <op>
// type ::= OP | OP i | OP r i | OP r i i | OP r r | OP r r i
// [ OP 8bits ]
// [ Reg 4bits ]
// [ Imm whatever is left of size ]

// 8 bit long instructions?
// op     (op)
// op i   (op imm)
// op ri  (op reg, imm)
// op rr  (op reg, reg)
// op rri (op reg, reg, imm)
var opToB = map[string]uint8 {
	"halt":  0,
	"movrr": 1,
	"movri": 2,
	"addrrr":3,
	"addrri":4,
	"subrrr":5,
	"subrri":6,
	"cmprr" :7,
	"cmpri" :8,
}

//var bToOp = flipMap(opToB)

func Encode(op Op, labels map[string]uint32) uint32 {
	// what will i do about labels?
	switch v := op.(type) {
	case Opp:
		// return opToB[v.Op]
	case Oprr:
		// return opToB[v.Op + "rr"] + registerToB[v.R1] + registerToB[v.R2]
	case Opri:
		// return opToB[v.Op + "ri"] + registerToB[v.R1] + iToB20(v.I)
	case Oprrr:
		// return opToB[v.Op + "rrr"] + registerToB[v.R1] + registerToB[v.R2] + registerToB[v.R3]
	case Oprri:
		// return opToB[v.Op + "rri"] + registerToB[v.R1] + registerToB[v.R2] + iToB16(v.I)
	}
	panic("error?")
}



/* first worry about encoding
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
*/

// for encoding/decoding
const (
    opShift = 28
    r1Shift = 24
    r2Shift = 20
    r3Shift = 16

    opMask  = 0xF
    regMask = 0xF
    imm24Mask = 0x00FFFFFF
    imm20Mask = 0x000FFFFF
    imm16Mask = 0x0000FFFF
)

func pack(op uint8, r1 uint8, r2 uint8, r3 uint8, imm uint32) uint32 {
    return uint32(op)  << opShift |
           uint32(r1) << r1Shift |
           uint32(r2) << r2Shift |
           uint32(r3) << r3Shift |
           imm
}
