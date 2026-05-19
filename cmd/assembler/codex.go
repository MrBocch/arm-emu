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

// OP, l (b|beq|blt|bgt label)
type Opl struct {
	Op string
	I int32
}
func (Opl) isOp() {}
func (o Opl) String() string {
	return fmt.Sprintf("%s %d", o.Op, o.I)
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
func flipMap[K comparable, V comparable](m map[K]V) map[V]K {
    flipped := make(map[V]K)
    for k, v := range m {
        flipped[v] = k
    }
    return flipped
}

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
	"bl"    :9, // not to be consfused with Branch with Link, its Branch to Label
	"bltl"  :10,// thinking about it, doesn't make sense to have (btype) (label) should have just, btype
	"beql"  :11,
	"bgtl"  :12,
	"lslrri":13,
	"lslrrr":14,
	"lsrrri":15,
	"lsrrrr":16, //this is getting ridiculous

	// ldr r0, 5
	"ldrDirect": 17,
	"strDirect": 18,

	// ldr r0, [r1] // loard r1 with value at address(r1)
	"ldrIndirect": 19,
	"strIndirect": 20,

	// I am a cheap hack
	// also this will encoded as a Opl 
	"sys" : 21,
}

var bToOp = flipMap(opToB)

// func Encode(op Op, labels map[string]int32) uint32 {
// encoder shouldn't get the map labels passed in right?
// that should be resolved by Op maker?
func Encode(op Op) uint32 {
	// what about immediates?
	// fuck golang for making me check in each case
	// this is likely a skill issue
	switch v := op.(type) {
	case Opp:
		if _, ok := opToB[v.Op]; !ok { panic("invalid opcode: " + v.Op + " likely error in analyzer?") }
		return packOp(opToB[v.Op])
	case Opl:
		if _, ok := opToB[v.Op]; !ok { panic("invalid opcode: " + v.Op + " likely error in analyzer?") }
		return packOpl(opToB[v.Op], v.I)
	case Oprr:
		if _, ok := opToB[v.Op]; !ok { panic("invalid opcode: " + v.Op + " likely error in analyzer?") }
		return packOprr(opToB[v.Op], v.R1, v.R2)
	case Opri:
		if _, ok := opToB[v.Op]; !ok { panic("invalid opcode: " + v.Op + " likely error in analyzer?") }
		return packOpri(opToB[v.Op], v.R1, uint32(v.I))
	case Oprrr:
		if _, ok := opToB[v.Op]; !ok { panic("invalid opcode: " + v.Op + " likely error in analyzer?") }
		return packOprrr(opToB[v.Op], v.R1, v.R2, v.R3)
	case Oprri:
		if _, ok := opToB[v.Op]; !ok { panic("invalid opcode: " + v.Op + " likely error in analyzer?") }
		return packOprri(opToB[v.Op], v.R1, v.R2, uint32(v.I))
	}
	panic("unknown op type")
}

func Decode(bin uint32) (Op, error) {
    op := uint8((bin >> 24) & 0xFF)

    // Find operation name
    opName, exists := bToOp[op]
    if !exists {
    	panic(fmt.Errorf("Dont recognize [%d] instruction", op))
    }

    // Decode based on operation type
    // looking at it now, should have done a
    // uint32 -> Op, uint32 -> Oprrri ....
    switch opName {
    case "halt":
    	return Opp{ Op : "halt", }, nil
    case "movrr":
        r1 := uint8((bin >> 20) & 0xF)
        r2 := uint8((bin >> 16) & 0xF)
        return Oprr{Op: "movrr", R1: r1, R2: r2}, nil
    case "movri":
        r1 := uint8((bin >> 20) & 0xF)
        imm := int32(bin & 0x000FFFFF)
        return Opri{Op: "movri", R1: r1, I: imm}, nil
    case "addrrr", "subrrr":
        // Format: 8-bit op + 4-bit r1 + 4-bit r2 + 4-bit r3
        // Note: remaining bits might be unused or for future expansion
        r1 := uint8((bin >> 20) & 0xF)
        r2 := uint8((bin >> 16) & 0xF)
		r3 := uint8((bin >> 12) & 0xF)
        return Oprrr{Op: opName, R1: r1, R2: r2, R3: r3}, nil

    case "addrri", "subrri":
        // Format: 8-bit op + 4-bit r1 + 4-bit r2 + 16-bit immediate
		r1 := uint8((bin >> 20) & 0xF)
		r2 := uint8((bin >> 16) & 0xF)
		i  := int32(bin & 0xFFFF) // signed 16-bit immediate
        return Oprri{Op: opName, R1: r1, R2: r2, I: i}, nil
    case "cmpri":
        r1 := uint8((bin>>20) & 0xF)
        imm := int32(bin & 0x000FFFFF)
        return Opri{Op: opName, R1: r1, I: imm}, nil
    case "cmprr":
        r1 := uint8((bin>>20) & 0xF)
        r2 := uint8((bin >> 16) & 0xF)
        return Oprr{Op: opName, R1: r1, R2: r2}, nil

    case "bl", "bltl", "beql", "bgtl" :
   	    imm := int32(bin & 0x00FFFFFF)
    	return Opl{Op: opName, I: imm}, nil

    case "sys":
   	    imm := int32(bin & 0x00FFFFFF)
    	return Opl{Op: opName, I: imm}, nil

    case "lslrrr", "lsrrrr":
	    r1 := uint8((bin>>20) & 0xF)
		r2 := uint8((bin >> 16) & 0xF)
		r3 := uint8((bin >> 12) & 0xF)
		return Oprrr{Op: opName, R1: r1, R2: r2, R3: r3 }, nil

    case "lslrri", "lsrrri":
	    r1 := uint8((bin>>20) & 0xF)
		r2 := uint8((bin >> 16) & 0xF)
		i  := int32(bin & 0xFFFF) // signed 16-bit immediate
		return Oprri{Op: opName, R1: r1, R2: r2, I: i}, nil

	case "ldrDirect", "strDirect":
		r1 := uint8((bin>>20) & 0xF)
		imm := int32(bin & 0x000FFFFF)
		return Opri{Op: opName, R1: r1, I:imm}, nil

	case "ldrIndirect", "strIndirect":
		r1 := uint8((bin >> 20) & 0xF)
    	r2 := uint8((bin >> 16) & 0xF)
     	return Oprr{Op: opName, R1: r1, R2:r2}, nil
    }


    return nil, fmt.Errorf("error: operation %q not implemented yet", opName)
}

// i wrote these functions why am i not using them?
// for encoding/decoding
func packOp(op uint8) uint32 {
	// op is 8 bit wide (32-8)
    return uint32(op) << 24
}

func packOpl(op uint8, label int32) uint32 {
	res := uint32(0)
    res |= uint32(op) << 24
    // label in bits 0-23
    res |= uint32(label) & 0x00FFFFFF

    return res
}
// so much gunk in code base man wtf am i doing?
func packOpi(op uint8, imm uint32) uint32 {
    res := uint32(0)
    res |= uint32(op) << 24
    res |= imm & 0x00FFFFFF // 24-bit immediate, bits 0-23
    return res
}

func packOpr(op uint8, r1 int8) uint32 {
	// op 8 , reg 4
	res := uint32(0)
	res |= uint32(op) << 24
	res |= uint32(r1) << 20
	return res
}

func packOprr(op uint8, r1 uint8, r2 uint8) uint32 {
	res := uint32(0)
	res |= uint32(op) << 24
	res |= uint32(r1) << 20
	res |= uint32(r2) << 16
	return res
}

func packOpri(op uint8, r1 uint8, imm uint32) uint32 {
    res := uint32(0)
    res |= uint32(op) << 24
    res |= uint32(r1) << 20
    res |= imm & 0x000FFFFF // 20-bit immediate, bits 0-19
    return res
}

func packOprri(op uint8, r1 uint8, r2 uint8, imm uint32) uint32 {
    res := uint32(0)
    res |= uint32(op) << 24
    res |= uint32(r1) << 20
    res |= uint32(r2) << 16
    res |= imm & 0x0000FFFF // 16-bit immediate, bits 0-15
    return res
}

func packOprrr(op uint8, r1 uint8, r2 uint8, r3 uint8) uint32 {
    res := uint32(0)
    res |= uint32(op) << 24
    res |= uint32(r1) << 20
    res |= uint32(r2) << 16
    res |= uint32(r3) << 12

    return res
}
