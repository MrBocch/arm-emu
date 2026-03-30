package vm

import (
	"fmt"
	"os"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
	// "github.com/charmbracelet/bubbles"

)

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
        	m.computer.Step()
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

func RunTui(mem []uint32) {
    p := tea.NewProgram(initialModel(mem))
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}
