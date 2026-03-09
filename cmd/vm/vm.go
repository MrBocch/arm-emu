package vm

import (
	"fmt"
	"os"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
	// "github.com/charmbracelet/bubbles"
	"strings"
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
	padMemory := padStringRight(memory, 1000)
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

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
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
 
func step(c *Computer) {
	fmt.Println(c.registers)
	fmt.Println(c.mem)
}
