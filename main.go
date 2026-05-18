package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func roundUp(numres float64) string {
	numres = math.Round(numres*1e9) / 1e9
	return strconv.FormatFloat(numres, 'f', -1, 64)
}

func solveSqrt(str []string) []string {
	for i := 0; i < len(str)-1; i++ {
		val := str[i]
		if val == "√" {
			nbr1, err := strconv.ParseFloat(str[i+1], 64)
			if err != nil {
				return nil
			}
			numres := math.Sqrt(nbr1)

			result := roundUp(numres)

			str = append(str[:i], append([]string{result}, str[i+2:]...)...)
			i = -1
		}
	}
	return str
}

func solvePower(str []string) []string {
	for i := 1; i < len(str)-1; i++ {
		val := str[i]
		if val == "^" {
			nbr1, err := strconv.ParseFloat(str[i-1], 64)
			if err != nil {
				return nil
			}
			nbr2, err := strconv.ParseFloat(str[i+1], 64)
			if err != nil {
				return nil
			}

			numres := math.Pow(nbr1, nbr2)

			result := roundUp(numres)

			str = append(str[:i-1], append([]string{result}, str[i+2:]...)...)
			i = -1
		}
	}
	return str
}

func solvePercent(str []string) []string {
	for i := 0; i < len(str); i++ {
		val := str[i]
		if val == "%" && i-1 >= 0 {
			nbr1, err := strconv.ParseFloat(str[i-1], 64)
			if err != nil {
				return nil
			}
			numres := nbr1 / 100
			result := roundUp(numres)
			str = append(str[:i-1], append([]string{result}, str[i+1:]...)...)
			i = -1
		}
	}
	return str
}

func solveMulDiv(str []string) []string {
	for i := 1; i < len(str)-1; i++ {
		val := str[i]
		if val == "x" || val == "÷" {
			nbr1, err := strconv.ParseFloat(str[i-1], 64)
			if err != nil {
				return nil
			}
			nbr2, err := strconv.ParseFloat(str[i+1], 64)
			if err != nil {
				return nil
			}
			var numres float64

			switch val {
			case "÷":
				numres = nbr1 / nbr2
			case "x":
				numres = nbr1 * nbr2
			}

			result := roundUp(numres)

			str = append(str[:i-1], append([]string{result}, str[i+2:]...)...)
			i = -1
		}
	}
	return str
}
func solveAddSub(str []string) []string {
	for i := 1; i < len(str)-1; i++ {
		val := str[i]
		if val == "+" || val == "-" {
			nbr1, err := strconv.ParseFloat(str[i-1], 64)
			if err != nil {
				return nil
			}
			nbr2, err := strconv.ParseFloat(str[i+1], 64)
			if err != nil {
				return nil
			}

			var numres float64

			switch val {
			case "+":
				numres = nbr1 + nbr2
			case "-":
				numres = nbr1 - nbr2
			}

			result := roundUp(numres)

			str = append(str[:i-1], append([]string{result}, str[i+2:]...)...)
			i = -1
		}
	}
	return str
}

func Calculate(str []string) []string {
	if str == nil {
		return []string{}
	}

	newstring := ""
	strArr := []string{}

	for _, val := range str {
		if val >= "0" && val <= "9" || val == "." {
			newstring += string(val)
		}
		if val == "÷" || val == "x" || val == "+" || val == "-" || val == "√" || val == "^" || val == "%" {
			if newstring != "" {
				strArr = append(strArr, newstring)
				newstring = ""
			}
			strArr = append(strArr, val)
			newstring = ""
		}
	}
	if newstring != "" {
		strArr = append(strArr, newstring)
		newstring = ""
	}

	fmt.Printf("%#v\n\n", strArr)
	strArr = solveSqrt(strArr)
	fmt.Printf("%#v\n", strArr)
	strArr = solvePower(strArr)
	fmt.Printf("%#v\n", strArr)
	strArr = solvePercent(strArr)
	fmt.Printf("%#v\n", strArr)
	strArr = solveMulDiv(strArr)
	fmt.Printf("%#v\n", strArr)
	strArr = solveAddSub(strArr)
	fmt.Printf("%#v\n", strArr)

	return strArr
}

func displayCalc(w fyne.Window, a fyne.App) {

	dashboardLabel := widget.NewLabelWithStyle("Ruach Calculator\n", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	display := widget.NewLabelWithStyle("\n\n\n\n", fyne.TextAlignTrailing, fyne.TextStyle{Bold: false})

	answer := widget.NewLabelWithStyle("", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true})

	buttonvalues := []string{
		"C", "√", "^", "DEL",
		"7", "8", "9", "÷",
		"4", "5", "6", "x",
		"1", "2", "3", "-",
		"0", ".", "%", "+",
	}

	var buttons []fyne.CanvasObject
	var button *widget.Button
	currentInput := []string{}
	currentInputS := ""

	for _, values := range buttonvalues {

		val := values
		if val != "C" && val != "DEL" {
			button = widget.NewButton(val, func() {
				currentInput = append(currentInput, val)
				currentInputS = strings.Join(currentInput, "")
				display.SetText(currentInputS)
				answer.SetText("")
			})
		}
		if val == "C" {
			button = widget.NewButton(val, func() {
				currentInput = nil
				currentInputS = strings.Join(currentInput, "")
				display.SetText(currentInputS)
				answer.SetText("")
			})
		}
		if val == "DEL" {
			button = widget.NewButton(val, func() {
				if len(currentInput) >= 1 {
					currentInput = currentInput[:len(currentInput)-1]
					answer.SetText("")
				} else {
				}
				currentInputS = strings.Join(currentInput, "")
				display.SetText(currentInputS)
				answer.SetText("")

			})
		}
		buttons = append(buttons, button)
	}

	grid := container.NewGridWithColumns(4, buttons...)

	equalsbutton := widget.NewButton(
		"=",
		func() {
			result := Calculate(currentInput)
			if len(result) == 1 {
				output := strings.Join(result, "")
				answer.SetText(output)

			}
			if len(result) != 1 {
				answer.SetText("Error! Invalid Expression (" + display.Text + ")")
			}
		})

	exitbutton := widget.NewButton(
		"Exit",
		func() {
			a.Quit()
		},
	)

	content := container.NewVBox(
		dashboardLabel,
		display,
		answer,
		grid,
		equalsbutton,
		exitbutton,
	)

	w.SetContent(content)

}

func main() {
	fmt.Println("Starting app...")

	a := app.New()

	fmt.Println("App created")

	w := a.NewWindow("Ruach Calculator")

	fmt.Println("Window created")

	displayCalc(w, a)

	w.Resize(fyne.NewSize(700, 500))

	fmt.Println("Window resized")

	w.Show()

	fmt.Println("Window shown")

	a.Run()

	fmt.Println("App ended")
}
