package vm

import (
	"fmt"
	"os"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
	// "github.com/charmbracelet/bubbles"
	"strings"
	"github.com/MrBocch/arm-emu/cmd/assembler"
)

type Computer struct {
	registers []int32
	mem       string 
}

func padStringRight(s string, size int) string {
	if len(s) >= size {
		panic("Not enough memory to write code")
	}

	return s + strings.Repeat("0", size-len(s))
}

func initComputer(register int, memory string) Computer {
	padMemory := padStringRight(memory, 32_000)
	return Computer {
		registers: make([]int32, register),
		mem      : padMemory,
	}
}

type model struct {
	// terminal screen
	width   int 
	height  int 

	computer Computer 
}

func initialModel(memory string) model {
	cpu := initComputer(16, memory)
	return model {
		computer: cpu,
	}
}

func (m model) Init() tea.Cmd {
    // Just return `nil`, which means "no I/O right now, please."
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    // Is it a key press?
    case tea.KeyPressMsg:
        // Cool, what was the actual key pressed?
        switch msg.String() {
        // These keys should exit the program.
        case "q":
            return m, tea.Quit
        case "n":
        	step(&m.computer)
        }

    case tea.WindowSizeMsg:
    	m.width = msg.Width
    	m.height = msg.Height
    }

    //fmt.Println(m.computer.registers)
    return m, nil
}

func (m model) View() tea.View {
	table := renderRegisters(m.computer.registers)

	help := lipgloss.NewStyle().
		Bold(true).
		Render("press (n) for next   (q) to quit")

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		help,
		"",
		table,
	)

	centered := lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)

	return tea.NewView(centered)
}

func renderRegisters(reg []int32) string {
	const (
		labelWidth = 12
		valueWidth = 10
	)

	headerStyle := lipgloss.NewStyle().
		Bold(true).
		BorderBottom(true)

	labelStyle := lipgloss.NewStyle().
		Width(labelWidth).
		Padding(0, 1)

	valueStyle := lipgloss.NewStyle().
		Width(valueWidth).
		Padding(0, 1)

	var lines []string

	// ----- Header -----
	header := lipgloss.JoinHorizontal(
		lipgloss.Top,
		labelStyle.Render("Register"),
		valueStyle.Render("Value"),
	)

	lines = append(lines, headerStyle.Render(header))

	// ----- Rows (16 registers) -----
	for i := 0; i < 16 && i < len(reg); i++ {
		row := lipgloss.JoinHorizontal(
			lipgloss.Top,
			labelStyle.Render(fmt.Sprintf("R%d", i)),
			valueStyle.Render(fmt.Sprintf("%d", reg[i])),
		)

		lines = append(lines, row)
	}

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

func Run(mem string) {
    p := tea.NewProgram(initialModel(mem))
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}

var regToI = map[string]int {
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
 
func step(c *Computer) {
	pc := c.registers[15]
	// fetch
	ins := c.mem[pc:pc+32]
	c.registers[15] += 32

	op, err := assembler.Decode(ins)
	// fmt.Println(op)
	if err != nil {
		panic("error")
	}

	 decode(c, op)
}

func decode(c *Computer, op assembler.Op) {
	// switch v := op.(type) {
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
		//fmt.Printf("%T\n", v)
	}
}

func executeOp(c *Computer, op string) {
	switch op {
	case "halt":
		os.Exit(0)
	}
}

func executeOpr(c *Computer, op string, ) {
	switch op {
	case "halt":
		os.Exit(0)
	}
}

func executeOprr(c *Computer, op string, r1 string, r2 string) {
	rs1 := regToI[r1]
	rs2 := regToI[r2]
	switch op {
	case "movrr":
		c.registers[rs1] = c.registers[rs2]
	}
}

func executeOpri(c *Computer, op string, r1 string, i int32) {
	rd := regToI[r1]
	switch op {
	case "movri":
		c.registers[rd] = i
	}
}

func executeOprri(c *Computer, op string, r1 string, r2 string, i int32) {
	rd := regToI[r1]
	rs1 := regToI[r2]

	switch op {
	case "subrri":
		c.registers[rd] = c.registers[rs1] - i
	case "addrri":
		c.registers[rd] = c.registers[rs1] + i
	}
}

func executeOprrr(c *Computer, op string, r1 string, r2 string, r3 string) {
	rd := regToI[r1]
	rs1 := regToI[r2]
	rs2 := regToI[r3]

	switch op {
	case "subrrr":
		c.registers[rd] = c.registers[rs1] - c.registers[rs2]
	case "addrrr":
		c.registers[rd] = c.registers[rs1] + c.registers[rs2]
	}
}
