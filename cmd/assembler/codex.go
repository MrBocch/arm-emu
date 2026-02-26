package assembler

import "fmt"

// go is shit for not having sum types
type Op interface {
	isOp()
}

// halt 
type Opp struct {
	op string
}
func (Opp) isOp() {}
func (o Opp) String() string {
	return fmt.Sprintf("%s", o.op)
}

// OP, R, R
type Oprr struct { 
	op string 
	r1 string 
	r2 string 
}
func (Oprr) isOp() {}
func (o Oprr) String() string {
	return fmt.Sprintf("%s %s, %s", o.op, o.r1, o.r2)
}


// OP R, I
type Opri struct {
	op string
	r1 string
	i  uint32
}
func (Opri) isOp() {}
func (o Opri) String() string {
	return fmt.Sprintf("%s %s, #%d", o.op, o.r1, o.i)
}


// OP R, R, I
type Oprri struct {
	op string
	r1 string
	r2 string
	i  uint32
}
func (Oprri) isOp() {}
func (o Oprri) String() string {
	return fmt.Sprintf("%s %s, %s, #%d", o.op, o.r1, o.r2, o.i)
}

// OP R, R, R
type Oprrr struct {
	op string
	r1 string
	r2 string
	r3 string
}
func (Oprrr) isOp() {}
func (o Oprrr) String() string {
	return fmt.Sprintf("%s %s, %s, %s", o.op, o.r1, o.r2, o.r3)
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


// 8 bit long instructions?
// op
// op i
// op ri (register, imediate)
// op rr (register, rr)
var opToB = map[string]string {
	"halt": "00000000",
}

var bToOp = flipMap(opToB)

func Encode(t Token, labels map[string]int) string {
	bs := ""
	return bs 
}

func Decode() {}

func padding(s string) string {
	if len(s) == 32 { return s }
	if len(s) > 32 { panic("length of instruction is too big")}

	for ; len(s) < 32 ; {
		s += "0"
	}

	return s 
}
