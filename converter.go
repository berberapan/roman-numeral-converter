package main

import (
	"fmt"
)

type RomanNumeral string

const (
    I RomanNumeral = "I"
    V RomanNumeral = "V"
    X RomanNumeral = "X"
    L RomanNumeral = "L"
    C RomanNumeral = "C"
    D RomanNumeral = "D"
    M RomanNumeral = "M"
)

var RomanIntMap = map[RomanNumeral]int {
    I: 1,
    V: 5,
    X: 10,
    L: 50,
    C: 100,
    D: 500,
    M: 1000,
}

func main() {
    test := romanToInt("XXX")
    test2 := intToRoman(1988)
    fmt.Println(test, test2)
}

func romanToInt(num string) int {
    conversion := 0
    for i := 0; i < len(num)-1; i++ {
        if RomanIntMap[RomanNumeral(string(num[i]))] < RomanIntMap[RomanNumeral(string(num[i+1]))] {
            conversion -= RomanIntMap[RomanNumeral(string(num[i]))]
        } else {
            conversion += RomanIntMap[RomanNumeral(string(num[i]))] 
        }
    }
    conversion += RomanIntMap[RomanNumeral(string(num[len(num)-1]))]
    return conversion
}

func intToRoman (num int) string {
    placeThousand := (num / 1000) % 10
    placeHundred := (num / 100) % 10
    placeTen := (num / 10) % 10
    placeOne := num % 10

    numeral := getStringValue(placeThousand, 3) + 
               getStringValue(placeHundred, 2) +
               getStringValue(placeTen, 1) + 
               getStringValue(placeOne, 0)

    return numeral 
}

func getStringValue(value int, place int) string {
    onePlaces := []string {"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
    tenPlaces := []string {"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
    hundredPlaces := []string {"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
    thousandPlaces := []string {"", "M", "MM", "MMM"}

    switch(place) {
        case 0:
            return onePlaces[value]
        case 1:
            return tenPlaces[value]
        case 2:
            return hundredPlaces[value]
        case 3:
            return thousandPlaces[value]
        default:
            return ""
    }
}
