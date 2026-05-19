package vm

import (
	// "os"
	"fmt"
	"strings"
	"github.com/MrBocch/arm-emu/cmd/assembler"
)

type Computer struct {
	registers []uint32
	mem       []uint32
	nflag     bool
	zflag     bool
	// armlite appears to not even use these for anything.
	// the manual says (pg 30) "we will not need them for now"
	// and never brings it up again
	// the reference manual does, BCS, BVS, BMI
	// will comeback to this.
	// cflag     bool
	// vflag     bool
	//
	// ugh
	stdout strings.Builder
}

var MEM_LIMIT = 1_000_000

func initComputer(registerCount int, memory []uint32) Computer {
    c := Computer{
        registers: make([]uint32, registerCount),
        mem:       make([]uint32, MEM_LIMIT),
		nflag: false,
		zflag: false,
    }
    c.LoadProgram(memory)
    return c
}

func (c *Computer) LoadProgram(bin []uint32) {
    if len(bin) > MEM_LIMIT {
        panic("program exceeds memory limit")
    }
    clear(c.mem)       // zero out first (Go 1.21+)
    copy(c.mem, bin)
}

var LR = 13
var SP = 14
var PC = 15

func (c *Computer) Step() {
	// fetch
	// what if pannics? here?
	addr := c.registers[PC]
	instr := c.mem[addr]
	c.registers[PC] += 1


	op, err := assembler.Decode(instr)
	// fmt.Println(op)
	if err != nil {
		panic("error at runtime")
	}

	 decode_execute(c, op)
}



func decode_execute(c *Computer, op assembler.Op) {
	switch v := op.(type) {
	case assembler.Opp:
		executeOp(c, v.Op)

	case assembler.Opl:
		executeOpl(c, v.Op, v.I)

	case assembler.Opri:
		executeOpri(c, v.Op, v.R1, v.I)

	case assembler.Oprr:
		executeOprr(c, v.Op, v.R1, v.R2)

	case assembler.Oprri:
		executeOprri(c, v.Op, v.R1, v.R2, v.I)

	case assembler.Oprrr:
		executeOprrr(c, v.Op, v.R1, v.R2, v.R3)

	default:
		panic("runtime error, unknown instruction")
	}
}

func executeOp(c *Computer, op string) {
	switch op {
	case "halt":
		fmt.Println("halted")
		// need to add support for opening a seperate asm file to run through a pop up
		// for ; ; {} wtf was I thinking?
	default:
		fmt.Printf("havent implemented (%v) yet? \n", op)
		panic("")
	}
}

func executeOpl(c *Computer, op string, i int32) {
	switch op {
	// dont confuse with Branch with Link
	case "bl":
		c.registers[PC] = uint32(i)
	case "bltl":
		if c.nflag && !c.zflag { c.registers[PC] = uint32(i) }
	case "beql":
		if !c.nflag && c.zflag { c.registers[PC] = uint32(i) }
	case "bgtl":
		if !c.nflag && !c.zflag { c.registers[PC] = uint32(i) }
	case "sys":
		runSys(c, i)
	default:
		fmt.Printf("havent implemented (%v) yet? \n", op)
		panic("")
	}
}

func executeOpri(c *Computer, op string, r1 uint8, i int32) {
	switch op {
	case "movri":
		c.registers[r1] = uint32(i)
	case "cmpri":
		a := int32(c.registers[r1])
		if a == i {
			c.nflag = false
			c.zflag = true
		}
		if a > i {
			c.nflag = false
		 	c.zflag = false
		}
		if a < i {
			c.nflag = true
			c.zflag = false
		}
	// should i check for out of bounds checking?
	case "ldrDirect":
		c.registers[r1] = c.mem[i]
	case "strDirect":
		c.mem[i] = c.registers[r1]

	default:
		fmt.Printf("havent implemented (%v) yet? \n", op)
		panic("")
	}
}

func executeOprr(c *Computer, op string, r1 uint8, r2 uint8) {
	switch op {
	case "movrr":
		c.registers[r1] = c.registers[r2]
	case "cmprr":
		a := int32(c.registers[r1])
		b := int32(c.registers[r2])
		if a == b {
			c.nflag = false
			c.zflag = true
		}
		if a > b {
			c.nflag = false
		 	c.zflag = false
		}
		if a < b {
			c.nflag = true
			c.zflag = false
		}
	// bounds checking?
	case "ldrIndirect":
			c.registers[r1] = c.mem[c.registers[r2]]
	case "strIndirect":
			c.mem[c.registers[r2]] = c.registers[r1]
	default:
		fmt.Printf("havent implemented (%v) yet? \n", op)
		panic("")
	}
}

func executeOprri(c *Computer, op string, r1 uint8, r2 uint8, i int32) {
	switch op {
	case "addrri":
		c.registers[r1] = c.registers[r2] + uint32(i)
	case "subrri":
		c.registers[r1] = c.registers[r2] - uint32(i)
	case "lslrri":
		c.registers[r1] = c.registers[r2] << uint32(i)
	case "lsrrri":
		c.registers[r1] = c.registers[r2] >> uint32(i)
	default:
		fmt.Printf("havent implemented (%v) yet? \n", op)
		panic("")
	}
}

func executeOprrr(c *Computer, op string, r1 uint8, r2 uint8, r3 uint8) {
	switch op {
	case "addrrr":
		c.registers[r1] = c.registers[r2] + c.registers[r3]
	case "subrrr":
		c.registers[r1] = c.registers[r2] - c.registers[r3]
	case "lslrrr":
		c.registers[r1] = c.registers[r2] << c.registers[r3]
	case "lsrrrr":
		c.registers[r1] = c.registers[r2] >> c.registers[r3]
	default:
		fmt.Printf("havent implemented (%v) yet? \n", op)
		panic("")
	}
}

func runSys(c *Computer, i int32) {
	// this code base deserves to be destroyed!
	switch i {
	// print r0 as signed int
	case 0:
		c.stdout.WriteString(
			fmt.Sprintf("%d", int32(c.registers[0])),
		)

	// print r0 as unsigned int
	case 1:
		c.stdout.WriteString(
			fmt.Sprintf("%d", c.registers[0]),
		)

	// print r0 as hex
	case 2:
		c.stdout.WriteString(
			fmt.Sprintf("0x%08X", c.registers[0]),
		)

	// print r0 as char
	case 3:
		c.stdout.WriteByte(byte(c.registers[0]))

	// print newline
	case 4:
		c.stdout.WriteByte('\n')

	// clear output buffer
	case 5:
		c.stdout.Reset()

	// halt and catch fire
	// https://en.wikipedia.org/wiki/Halt_and_Catch_Fire_(computing)
	case 67:
		panic("HALT AND CATCHING FIRE")

	default:
		fmt.Printf("Dont recognize this syscalls, panicing")
		panic("")
	}
}
