package main

import (
	"testing"
)

func expectationError(t *testing.T, expected, got string, position int) {
	t.Errorf("Expected '%s' in position %d but got '%s' instead", expected, position, got)
}
func BasicConversion(t *testing.T, input string, expected []string) {
	result := translateWord(input)
	if len(result) != len(expected) {
		t.Errorf("Expected %d values back but got %d instead", len(expected), len(result))
	} else {
		for index, value := range result {
			if value != expected[index] {
				t.Errorf("Expected '%s' in position %d but got '%s' instead", expected[index], index, value)
			}
		}
	}
}

var wtfTranslation = []string{"Whiskey", "Tango", "Foxtrot"}

func Test_Conversion1(t *testing.T) {
	BasicConversion(t, "WTF", wtfTranslation)
}
func Test_Conversion2(t *testing.T) {
	BasicConversion(t, "wtf", wtfTranslation)
}
