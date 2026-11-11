
package vm

import (
	"testing"

	"github.com/MrBocch/arm-emu/cmd/assembler"
)

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func newTestComputer(bin []uint32) *Computer {
	c := initComputer(16, bin)
	return &c
}

func runOne(c *Computer, op assembler.Op) {
	bin := []uint32{assembler.Encode(op)}
	c.LoadProgram(bin)
	c.Step()
}

func assertReg(t *testing.T, c *Computer, idx int, want uint32) {
	t.Helper()
	if got := c.registers[idx]; got != want {
		t.Errorf("R%d = %d; want %d", idx, got, want)
	}
}

// ---------------------------------------------------------------------------
// MOV: registro ← inmediato
// ---------------------------------------------------------------------------

func TestVM_MOV_Immediate(t *testing.T) {
	c := newTestComputer(nil)
	runOne(c, assembler.Opri{Op: "movri", R1: 0, I: 42})
	assertReg(t, c, 0, 42)
}

func TestVM_MOV_Register(t *testing.T) {
	c := newTestComputer(nil)
	c.registers[1] = 7
	runOne(c, assembler.Oprr{Op: "movrr", R1: 0, R2: 1})
	assertReg(t, c, 0, 7)
}

// ---------------------------------------------------------------------------
// ADD
// ---------------------------------------------------------------------------

func TestVM_ADD_ThreeRegisters(t *testing.T) {
	c := newTestComputer(nil)
	c.registers[1] = 3
	c.registers[2] = 4
	runOne(c, assembler.Oprrr{Op: "addrrr", R1: 0, R2: 1, R3: 2})
	assertReg(t, c, 0, 7)
}

func TestVM_ADD_RegisterImmediate(t *testing.T) {
	c := newTestComputer(nil)
	c.registers[1] = 10
	runOne(c, assembler.Oprri{Op: "addrri", R1: 0, R2: 1, I: 5})
	assertReg(t, c, 0, 15)
}

// ---------------------------------------------------------------------------
// SUB
// ---------------------------------------------------------------------------

func TestVM_SUB_ThreeRegisters(t *testing.T) {
	c := newTestComputer(nil)
	c.registers[1] = 10
	c.registers[2] = 3
	runOne(c, assembler.Oprrr{Op: "subrrr", R1: 0, R2: 1, R3: 2})
	assertReg(t, c, 0, 7)
}

func TestVM_SUB_RegisterImmediate(t *testing.T) {
	c := newTestComputer(nil)
	c.registers[1] = 20
	runOne(c, assembler.Oprri{Op: "subrri", R1: 0, R2: 1, I: 8})
	assertReg(t, c, 0, 12)
}

// ---------------------------------------------------------------------------
// CMP y banderas (nflag, zflag)
// ---------------------------------------------------------------------------

func TestVM_CMP_Register_Equal_SetsZ(t *testing.T) {
	c := newTestComputer(nil)
	c.registers[0] = 5
	c.registers[1] = 5
	runOne(c, assembler.Oprr{Op: "cmprr", R1: 0, R2: 1})

	if !c.zflag {
		t.Errorf("zflag debería estar activa tras cmprr de valores iguales")
	}
	if c.nflag {
		t.Errorf("nflag no debería estar activa tras cmprr de valores iguales")
	}
}

func TestVM_CMP_Register_Less_SetsN(t *testing.T) {
	c := newTestComputer(nil)
	c.registers[0] = 3
	c.registers[1] = 7
	runOne(c, assembler.Oprr{Op: "cmprr", R1: 0, R2: 1})

	if !c.nflag {
		t.Errorf("nflag debería estar activa cuando R0 < R1")
	}
	if c.zflag {
		t.Errorf("zflag no debería estar activa cuando R0 < R1")
	}
}

func TestVM_CMP_Immediate_Equal_SetsZ(t *testing.T) {
	c := newTestComputer(nil)
	c.registers[0] = 10
	runOne(c, assembler.Opri{Op: "cmpri", R1: 0, I: 10})

	if !c.zflag {
		t.Errorf("zflag debería estar activa tras cmpri de valores iguales")
	}
}

// ---------------------------------------------------------------------------
// PC: avance tras instrucción normal
// ---------------------------------------------------------------------------

func TestVM_PC_IncrementsAfterInstruction(t *testing.T) {
	c := newTestComputer(nil)
	runOne(c, assembler.Opri{Op: "movri", R1: 0, I: 1})

	// PC empieza en 0, Step lo incrementa a 1
	if c.registers[PC] != 1 {
		t.Errorf("PC = %d; want 1", c.registers[PC])
	}
}

// ---------------------------------------------------------------------------
// Branch incondicional (bl)
// ---------------------------------------------------------------------------

func TestVM_BL_JumpsToAddress(t *testing.T) {
	c := newTestComputer(nil)
	runOne(c, assembler.Opl{Op: "bl", I: 8})

	if c.registers[PC] != 8 {
		t.Errorf("PC = %d tras bl; want 8", c.registers[PC])
	}
}

// ---------------------------------------------------------------------------
// Branch condicional NO tomado (bltl / beql / bgtl)
// ---------------------------------------------------------------------------

func TestVM_BLTL_NotTaken_WhenZFlagSet(t *testing.T) {
	c := newTestComputer(nil)
	c.zflag = true // bltl requiere !nflag && zflag, así que NO debe saltar
	runOne(c, assembler.Opl{Op: "bltl", I: 100})

	// PC debe avanzar normalmente a 1, no saltar a 100
	if c.registers[PC] != 1 {
		t.Errorf("PC = %d; want 1 (bltl no tomado)", c.registers[PC])
	}
}

func TestVM_BLTL_Taken_WhenConditionsMet(t *testing.T) {
	c := newTestComputer(nil)
	c.zflag = false
	c.nflag = false // bltl se toma cuando !nflag && zflag? 
	// Revisa la condición exacta en vm.go: case "bltl": if c.nflag && !c.zflag
	// Según el código: if c.nflag && !c.zflag { PC = i }
	c.nflag = true
	c.zflag = false
	runOne(c, assembler.Opl{Op: "bltl", I: 100})

	if c.registers[PC] != 100 {
		t.Errorf("PC = %d; want 100 (bltl tomado)", c.registers[PC])
	}
}

// ---------------------------------------------------------------------------
// Memoria: STR/LDR directo e indirecto
// ---------------------------------------------------------------------------

func TestVM_STR_Then_LDR_Direct(t *testing.T) {
	c := newTestComputer(nil)
	c.registers[0] = 123
	runOne(c, assembler.Opri{Op: "strDirect", R1: 0, I: 10}) // mem[10] = R0
	runOne(c, assembler.Opri{Op: "movri", R1: 0, I: 0})      // limpiar R0
	runOne(c, assembler.Opri{Op: "ldrDirect", R1: 0, I: 10}) // R0 = mem[10]

	assertReg(t, c, 0, 123)
}

func TestVM_STR_Then_LDR_Indirect(t *testing.T) {
	c := newTestComputer(nil)
	c.registers[0] = 456
	c.registers[1] = 20 // dirección base
	runOne(c, assembler.Oprr{Op: "strIndirect", R1: 0, R2: 1}) // mem[R1] = R0
	runOne(c, assembler.Opri{Op: "movri", R1: 0, I: 0})
	runOne(c, assembler.Oprr{Op: "ldrIndirect", R1: 0, R2: 1}) // R0 = mem[R1]

	assertReg(t, c, 0, 456)
}

// ---------------------------------------------------------------------------
// HALT
// ---------------------------------------------------------------------------

func TestVM_Halt_StopsExecution(t *testing.T) {
	c := newTestComputer(nil)
	runOne(c, assembler.Opp{Op: "halt"})

	if !c.halted {
		t.Errorf("la VM debería estar detenida tras halt")
	}
}

// ---------------------------------------------------------------------------
// Estado inicial limpio
// ---------------------------------------------------------------------------

func TestVM_InitialStateIsZero(t *testing.T) {
	c := newTestComputer(nil)

	for i := 0; i < 16; i++ {
		if c.registers[i] != 0 {
			t.Errorf("R%d inicial = %d; want 0", i, c.registers[i])
		}
	}
	if c.nflag || c.zflag {
		t.Errorf("banderas iniciales deben estar en falso")
	}
	if c.halted {
		t.Errorf("halted inicial debe ser false")
	}
}
