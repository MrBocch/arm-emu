package vm

import (
	"os"
	"github.com/MrBocch/arm-emu/cmd/assembler"
)

type Computer struct {
	registers []uint32
	mem       []uint32
}

func initComputer(registerCount int, memory []uint32) Computer {
	// should decide how much
	// TODO
	if len(memory) > 1_000_000 { panic("hit memory limit") }
	return Computer {
		registers: make([]uint32, registerCount),
		mem      : memory,
	}
}

var PC = 15

func step(c *Computer) {
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

	// case assembler.Oprri:
	case assembler.Oprrr:
		executeOprrr(c, v.Op, v.R1, v.R2, v.R3)

	default:
		panic("runtime error, unknown instruction")
	}
}

func executeOp(c *Computer, op string) {
	switch op {
	case "halt":
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

func executeOprrr(c *Computer, op string, r1 uint8, r2 uint8, r3 uint8) {
	switch op {
	default:
		panic("havent implemented (this instruction) yet?")
	}
}
