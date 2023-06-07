package unit_testing

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var a = 4
var b = 2
var calculator = Calculator{a: a, b: b}

func TestMultiply(t *testing.T) {
	result := calculator.Multiply()
	expected := 9
	assert.Equal(t, result, expected, fmt.Sprintf("Hasil perkaliannya adalah %d tapi yang di dapetkan %d", expected, result))
	//if result != expected {
	//	t.Errorf("Expected %d, but got %d", expected, result)
	//}
}

func TestAdd(t *testing.T) {

	result := calculator.Add()
	expected := 6
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSubstract(t *testing.T) {
	result := calculator.Multiply()
	expected := 8
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestDivine(t *testing.T) {
	result := calculator.Divine()
	expected := 2
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
