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
        result := romanToInt(test.stringValue)

        if result != test.intValue {
            t.Errorf("romanToInt(%s) FAILED. Expected %d got %d\n", test.stringValue, test.intValue, result)
        } else {
            t.Logf("romanToInt(%s) PASSED.", test.stringValue)
        }
    }
}

func TestIntToRoman(t *testing.T) {
    for _, test := range testCases {
        result := intToRoman(test.intValue)

        if result != test.stringValue {
            t.Errorf("intToRoman(%d) FAILED. Expected %s got %s\n", test.intValue, test.stringValue, result) 
        } else {
            t.Logf("intToRoman(%d) PASSED.", test.intValue)
        }
    }
}
