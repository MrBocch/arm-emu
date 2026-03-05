package vm

import (
	"fmt"
	"os"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
	// "github.com/charmbracelet/bubbles"
)

type Computer struct {
	registers []int32
	mem       []int32 
}

func initComputer(register int, memory int) Computer {
	return Computer {
		registers: make([]int32, register),
		mem      : make([]int32, int(memory)),
	}
}

type model struct {
	// terminal screen
	width   int 
	height  int 

	computer Computer 
}

func initialModel() model {
	cpu := initComputer(16, 1000)
	return model{ cpu }
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
        case "up", "k":
			m.computer.registers[0]++
        case "down", "j":
			m.computer.registers[0]--
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

	centered := lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		table,
	)

	return tea.NewView(centered)
}

func renderRegisters(reg []int) string {
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

func Run() {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}
