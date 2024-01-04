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

    result := widget.NewLabel("")

    action := func(input string) {
        if isInputString(input) {
            if isStringValid(input) {
                numeral := strings.ToUpper(input)
                romanInt := fmt.Sprint(RomanToInt(numeral))
                if romanInt == "0" {
                    result.SetText("Input text in incorrect format.")
                } else {
                    result.SetText(romanInt)
                }
            } else {
                result.SetText("Roman Numeral format doesn't allow four letters in a row.")
            }
        } else {
            num, _ := strconv.Atoi(input)
            if num < 1 {
                result.SetText("Roman Numerals doesn't exist for 0 or negative values.")
            } else if num > 3999 {
                result.SetText("Input number too big. Largest Roman Numeral is 3999.")
            } else {
                result.SetText(IntToRoman(num))
            }
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
