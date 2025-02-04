package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
)

// PrintType 2. Определяет тип каждой переменной и выводит его на экран.
func PrintType(arg interface{}) {
	fmt.Printf("Тип переменной: %v\n", reflect.TypeOf(arg))
}

// ConvertToString 3. Преобразует все переменные в строковый тип и объединяет их в одну строку.
func ConvertToString(arg interface{}) string {
	switch arg := arg.(type) {
	case int:
		return strconv.Itoa(arg)
	case float64:
		return strconv.FormatFloat(arg, 'f', -1, 64)
	case string:
		return arg
	case bool:
		return strconv.FormatBool(arg)
	case complex64:
		return fmt.Sprint(arg)
	default:
		return fmt.Sprintf("%v", arg)
	}
}

func ConvertAllStrings(args ...interface{}) []string {
	strings := make([]string, len(args))
	for _, v := range args {
		strings = append(strings, ConvertToString(v))
	}
	return strings
}

func SliceInStr(strings []string) string {
	str := ""
	for _, v := range strings {
		str += v
	}
	return str
}

// StrInSliceRune 4. Преобразовать эту строку в срез рун.
func StrInSliceRune(str string, salt string) []rune {
	runeSlice := []rune(str)
	mid := len(runeSlice) / 2
	sliceWithSalt := append(runeSlice[:mid], append([]rune(salt), runeSlice[mid:]...)...)
	return sliceWithSalt
}

// HashRunes 5. Захэшировать этот срез рун SHA256, добавив в середину соль "go-2024" и вывести результат.
func HashRunes(sliceWithSalt []rune) string {
	hash := sha256.Sum256([]byte(string(sliceWithSalt)))
	return hex.EncodeToString(hash[:])
}

func main() {

	// 1. Создает несколько переменных различных типов данных:
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	// 2. Определяет тип каждой переменной и выводит его на экран.
	PrintType(numDecimal)
	PrintType(numOctal)
	PrintType(numHexadecimal)
	PrintType(pi)
	PrintType(name)
	PrintType(isActive)
	PrintType(complexNum)

	// 3. Преобразует все переменные в строковый тип и объединяет их в одну строку.
	strings := ConvertAllStrings(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	bigStr := SliceInStr(strings)

	// 4. Преобразовать эту строку в срез рун.
	sliceWithSalt := StrInSliceRune(bigStr, "go-2024")

	// 5. Захэшировать этот срез рун SHA256, добавив в середину соль "go-2024" и вывести результат.
	hash := HashRunes(sliceWithSalt)

	fmt.Println(hash)
}
