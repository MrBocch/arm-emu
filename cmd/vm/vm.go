package vm

import (
	"os"
	"fmt"
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

	 decode(c, op)
}



func decode(c *Computer, op assembler.Op) {
	switch v := op.(type) {
	case assembler.Opp:
		executeOp(c, v.Op)

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
		os.Exit(0)
	default:
		panic("havent implemented (this instruction) yet?")
	}
}

func executeOpri(c *Computer, op string, r1 uint8, i int32) {
	switch op {
	case "movri":
		c.registers[r1] = uint32(i)
	default:
		panic("havent implemented (this instruction) yet?")
	}
}

func executeOprr(c *Computer, op string, r1 uint8, r2 uint8) {
	switch op {
	case "movrr":
		c.registers[r1] = c.registers[r2]
	default:
		panic("havent implemented (this instruction) yet?")
	}
}

func executeOprri(c *Computer, op string, r1 uint8, r2 uint8, i int32) {
	switch op {
	case "addrri":
		c.registers[r1] = c.registers[r2] + uint32(i)
	case "subrri":
		c.registers[r1] = c.registers[r2] - uint32(i)
	default:
		panic("haven't implemented (this instruction) yet?")
	}
}

func executeOprrr(c *Computer, op string, r1 uint8, r2 uint8, r3 uint8) {
	switch op {
	case "addrrr":
		c.registers[r1] = c.registers[r2] + c.registers[r3]
	case "subrrr":
		c.registers[r1] = c.registers[r2] - c.registers[r3]
	default:
		panic("havent implemented (this instruction) yet?")
	}
}
