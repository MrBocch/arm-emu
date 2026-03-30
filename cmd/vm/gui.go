package vm

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
)

var APP_TITLE = "FIME CORE VM"

var (
	R0Label *widget.Label
	R1Label *widget.Label
	R2Label *widget.Label
	R3Label *widget.Label
	R4Label *widget.Label
	R5Label *widget.Label
	R6Label *widget.Label
	R7Label *widget.Label
	R8Label *widget.Label
	R9Label *widget.Label
	R10Label *widget.Label
	R11Label *widget.Label
	R12Label *widget.Label
	LRLabel *widget.Label
	SPLabel *widget.Label
	PCLabel *widget.Label
)

var vm = initComputer(16, make([]uint32, 0))

func RunGui(mem []uint32) {
	// TODO : check for memory things.
	vm.mem = mem


	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())

	w := a.NewWindow(APP_TITLE)
	w.Resize(fyne.NewSize(700, 700))


	split := container.NewHSplit(makeLeft(), makeRight())

	w.SetContent(split)
	w.ShowAndRun()
}

func makeLeft() fyne.CanvasObject {
	R1Label  = widget.NewLabel("R1:  0x00000000")
	R2Label  = widget.NewLabel("R2:  0x00000000")
	R3Label  = widget.NewLabel("R3:  0x00000000")
	R4Label  = widget.NewLabel("R4:  0x00000000")
	R5Label  = widget.NewLabel("R5:  0x00000000")
	R6Label  = widget.NewLabel("R6:  0x00000000")
	R7Label  = widget.NewLabel("R7:  0x00000000")
	R8Label  = widget.NewLabel("R8:  0x00000000")
	R9Label  = widget.NewLabel("R9:  0x00000000")
	R10Label = widget.NewLabel("R10: 0x00000000")
	R11Label = widget.NewLabel("R11: 0x00000000")
	R12Label = widget.NewLabel("R12: 0x00000000")
	LRLabel  = widget.NewLabel("LR:  0x00000000")
	SPLabel  = widget.NewLabel("SP:  0x00000000")
	PCLabel  = widget.NewLabel("PC:  0x00000000")

	return container.NewBorder(
		widget.NewLabel("REGISTERS"),
		nil, nil, nil,
		container.NewVBox(
			R1Label, R2Label, R3Label, R4Label,
			R5Label, R6Label, R7Label, R8Label,
			R9Label, R10Label, R11Label, R12Label,
			LRLabel, SPLabel, PCLabel,
		),
	)
}

func updateRegisters() {
	R1Label.SetText(fmt.Sprintf("R0:  0x%08X", vm.registers[0]))
	R2Label.SetText(fmt.Sprintf("R1:  0x%08X", vm.registers[1]))
	R3Label.SetText(fmt.Sprintf("R2:  0x%08X", vm.registers[2]))
	R4Label.SetText(fmt.Sprintf("R3:  0x%08X", vm.registers[3]))
	R5Label.SetText(fmt.Sprintf("R4:  0x%08X", vm.registers[4]))
	R6Label.SetText(fmt.Sprintf("R5:  0x%08X", vm.registers[5]))
	R7Label.SetText(fmt.Sprintf("R6:  0x%08X", vm.registers[6]))
	R8Label.SetText(fmt.Sprintf("R7:  0x%08X", vm.registers[7]))
	R9Label.SetText(fmt.Sprintf("R8:  0x%08X", vm.registers[8]))
	R10Label.SetText(fmt.Sprintf("R9: 0x%08X", vm.registers[9]))
	R11Label.SetText(fmt.Sprintf("R10: 0x%08X", vm.registers[10]))
	R12Label.SetText(fmt.Sprintf("R11: 0x%08X", vm.registers[11]))
	R12Label.SetText(fmt.Sprintf("R12: 0x%08X", vm.registers[12]))
	LRLabel.SetText(fmt.Sprintf("LR:  0x%08X", vm.registers[13]))
	SPLabel.SetText(fmt.Sprintf("SP:  0x%08X", vm.registers[14]))
	PCLabel.SetText(fmt.Sprintf("PC:  0x%08X", vm.registers[15]))
}

func makeRight() fyne.CanvasObject {
    return container.NewVBox(
        widget.NewButtonWithIcon("Step", theme.MediaPlayIcon(), func() {
            vm.Step()
            updateRegisters()
        }),
    )
}
