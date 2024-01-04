package main

import "testing"

type test struct {
    text string
    result bool
}

func TestIsInputString(t *testing.T) {
    testData := []test {
        {"This is a string", true},
        {"1010", false},
        {"12 12", true},
        {"99test99", true},
        {"8_", true},
    }

    for _, example := range testData {
        result := isInputString(example.text)

        if result != example.result {
            t.Errorf("isInputString(%s) FAILED. Expected %t got %t", example.text, example.result, result)
        } else {
            t.Logf("isInputString(%s) PASSED.", example.text)
        }
    }
}

func TestIsStringValid(t *testing.T) {
    testData := []test {
        {"XXX", true},
        {"XXXX", false},
        {"CXXXIXX", true},
        {"CXXXXI", false},
        {"XXXx", false},
    }

    for _, example := range testData {
        result := isStringValid(example.text)

         if result != example.result {
            t.Errorf("isStringValid(%s) FAILED. Expected %t got %t", example.text, example.result, result)
        } else {
            t.Logf("isStringValid(%s) PASSED.", example.text)
        }

    }
}
