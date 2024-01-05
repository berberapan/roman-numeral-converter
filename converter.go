package main

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

func RomanToInt(num string) int {
    conversion := 0
    if len(num) == 0 {
        return conversion
    }
    for i := 0; i < len(num)-1; i++ {
        currentValue := RomanIntMap[RomanNumeral(string(num[i]))]
        nextValue := RomanIntMap[RomanNumeral(string(num[i+1]))]

        if currentValue == 0 || (i == len(num)-2 && nextValue == 0) {
            return 0
        } else if currentValue < nextValue {
            conversion -= currentValue
        } else {
            conversion += currentValue
        }
    }
    conversion += RomanIntMap[RomanNumeral(string(num[len(num)-1]))]
    return conversion
}

func IntToRoman (num int) string {
    if IsValueInRange(num) {
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
    return ""
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

func IsValueInRange(num int) bool {
    if num < 1 || num > 3999 {
        return false
    }
    return true
}
