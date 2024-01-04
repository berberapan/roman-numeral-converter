package main

import "testing"

type romanNumeral struct {
    stringValue string
    intValue int
}

var testCases = []romanNumeral {
        {"MDCLXVI", 1666}, 
        {"MCMXCIX", 1999},
        {"DCCC", 800},
        {"XLVIII", 48},
}

func TestRomanToInt(t *testing.T) {
    for _, test := range testCases {
        result := RomanToInt(test.stringValue)

        if result != test.intValue {
            t.Errorf("romanToInt(%s) FAILED. Expected %d got %d\n", test.stringValue, test.intValue, result)
        } else {
            t.Logf("romanToInt(%s) PASSED.", test.stringValue)
        }
    }
}

func TestIntToRoman(t *testing.T) {
    for _, test := range testCases {
        result := IntToRoman(test.intValue)

        if result != test.stringValue {
            t.Errorf("intToRoman(%d) FAILED. Expected %s got %s\n", test.intValue, test.stringValue, result) 
        } else {
            t.Logf("intToRoman(%d) PASSED.", test.intValue)
        }
    }
}

func TestGetStringValue(t *testing.T) {
    type inputData struct {
        value int
        place int
        numeral string
    }

    tests := []inputData {
        {2, 4, ""},
        {0, 2, ""},
        {5, 1, "L"},
        {9, 0, "IX"},
        {1, 3, "M"},
    }

    for _, test := range tests {
        result := getStringValue(test.value, test.place)

        if result != test.numeral {
            t.Errorf("getStringValue(%d, %d) FAILED. Expected %s got %s", test.value, test.place, test.numeral, result)
        } else {
            t.Logf("getStringValue(%d, %d) PASSED.", test.value, test.place)
        }
    }
    
}
