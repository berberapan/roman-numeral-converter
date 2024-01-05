package main

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func UserWindow() {
    app := app.New()
    window := app.NewWindow("Converter")
    window.Resize(fyne.NewSize(500,250))

    result := canvas.NewText("", color.White)
    result.TextSize = 18

    action := func(input string) {
        if !isInputString(input) {
            num, _ := strconv.Atoi(input)
            if !IsValueInRange(num) {
                result.Text = "Roman Numerals smallest value is 1 and biggest 3999."
            } else {
                result.Text = IntToRoman(num)
            }
            return
        }

        if !isStringValid(input) {
            result.Text = "Roman Numeral format doesn't allow four letters in a row."
            return
        }

        numeral := strings.ToUpper(input)
        romanInt := fmt.Sprint(RomanToInt(numeral))
        correctFormat := IntToRoman(RomanToInt(numeral))

        if romanInt == "0" {
            result.Text = "Input text in incorrect format."
            return
        }

        if numeral == correctFormat {
            result.Text = romanInt
        } else if !IsValueInRange(RomanToInt(numeral)) {
            result.Text = fmt.Sprintf("Value too big. That order would be %s, if allowed.", romanInt)
        } else {
            result.Text = fmt.Sprintf("%s. Correct format is %s", romanInt, correctFormat)
        }
    }

    userInput := widget.NewEntry()
    userInput.SetPlaceHolder("Add number or numeral you want to convert")
    userInput.OnSubmitted = func(content string) { action(content) }

    convertButton := widget.NewButton("Convert", func() { action(userInput.Text) })

    emptySpace := widget.NewLabel("")

    window.SetContent(container.NewBorder(
        container.NewBorder(
            container.NewCenter(canvas.NewText("Roman Numeral Converter", color.White)),
            emptySpace,
            emptySpace,
            convertButton,
            userInput),
        emptySpace,
        nil,
        nil,
        container.NewCenter(result),
        ))

    window.ShowAndRun()
}

func isInputString(input string) bool {
    if _, err := strconv.Atoi(input); err == nil {
        return false
    }
    return true
}

func isStringValid(input string) bool {
    formattedInput := strings.ToUpper(input)
    var repeatChar rune
    repeatCount := 1
    for _, char := range formattedInput {
        if char == repeatChar {
            repeatCount++
            if repeatCount > 3 {
                return false
            }
        } else {
            repeatChar = char
            repeatCount = 1
        }
    }
    return true
}

