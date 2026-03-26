package vm

import (
	"fmt"
	"os"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
	// "github.com/charmbracelet/bubbles"
	"github.com/MrBocch/arm-emu/cmd/assembler"
)

type Computer struct {
	registers []uint32
	mem       []uint32
}

func initComputer(registerCount int, memory []uint32) Computer {
	return Computer {
		registers: make([]uint32, registerCount),
		mem      : memory,
	}
}

type model struct {
	// terminal screen
	width   int
	height  int

	computer Computer
}

func initialModel(memory []uint32) model {
	if len(memory) > 4_000 { panic("hit memory?") }

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

func renderRegisters(reg []uint32) string {
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

func Run(mem []uint32) {
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
