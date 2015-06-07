// Copyright (c) 2015 Joshua Scoggins
//
// This software is provided 'as-is', without any express or implied
// warranty. In no event will the authors be held liable for any damages
// arising from the use of this software.
//
// Permission is granted to anyone to use this software for any purpose,
// including commercial applications, and to alter it and redistribute it
// freely, subject to the following restrictions:
//
// 1. The origin of this software must not be misrepresented; you must not
//    claim that you wrote the original software. If you use this software
//    in a product, an acknowledgement in the product documentation would be
//    appreciated but is not required.
// 2. Altered source versions must be plainly marked as such, and must not be
//    misrepresented as being the original software.
// 3. This notice may not be removed or altered from any source distribution.

// convert a phrase to the corresponding nato alphabet to make life much easier
package main

import (
	"flag"
	"fmt"
	"strings"
)

var onePerLine = flag.Bool("one-per-line", false, "The translation of each character is on its own line")
var input = flag.String("input", "", "The input to translate")

func translateLetter(a rune) string {
	switch a {
	case 'a', 'A':
		return "Alfa"
	case 'b', 'B':
		return "Bravo"
	case 'c', 'C':
		return "Charlie"
	case 'd', 'D':
		return "Delta"
	case 'e', 'E':
		return "Echo"
	case 'f', 'F':
		return "Foxtrot"
	case 'g', 'G':
		return "Golf"
	case 'h', 'H':
		return "Hotel"
	case 'i', 'I':
		return "India"
	case 'j', 'J':
		return "Juliett"
	case 'k', 'K':
		return "Kilo"
	case 'l', 'L':
		return "Lima"
	case 'm', 'M':
		return "Mike"
	case 'n', 'N':
		return "November"
	case 'o', 'O':
		return "Oscar"
	case 'p', 'P':
		return "Papa"
	case 'q', 'Q':
		return "Quebec"
	case 'r', 'R':
		return "Romeo"
	case 's', 'S':
		return "Sierra"
	case 't', 'T':
		return "Tango"
	case 'u', 'U':
		return "Uniform"
	case 'v', 'V':
		return "Victor"
	case 'w', 'W':
		return "Whiskey"
	case 'x', 'X':
		return "Xray"
	case 'y', 'Y':
		return "Yankee"
	case 'z', 'Z':
		return "Zulu"
	case '1':
		return "One"
	case '2':
		return "Two"
	case '3':
		return "Three"
	case '4':
		return "Four"
	case '5':
		return "Five"
	case '6':
		return "Six"
	case '7':
		return "Seven"
	case '8':
		return "Eight"
	case '9':
		return "Nine"
	case '0':
		return "Zero"
	case '.':
		return "Decimal"
	default:
		return fmt.Sprintf("%c", a)
	}
}
func translateWord(str string) []string {
	output := make([]string, len(str))
	for index, value := range str {
		output[index] = translateLetter(value)
	}
	return output
}

func main() {
	flag.Parse()
	if *input == "" {
		fmt.Println("No input provided")
		flag.Usage()
		return
	}
	fmt.Println("Input was", *input)
	result := translateWord(*input)
	if *onePerLine {
		for _, value := range result {
			if *onePerLine {
				fmt.Println(value)
			}
		}
	} else {
		fmt.Println(strings.Join(result, " "))
	}
}
