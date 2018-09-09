package leetcode

import (
	"fmt"
	"testing"
)

func applyTest(num1, num2, expectation string, t *testing.T) {
	if result := multiply(num1, num2); result != expectation {
		t.Errorf("expecting %v, got %v", expectation, result)
	}
}
func TestMultiply(t *testing.T) {
	var num1 string
	var num2 string
	var expectation string

	num1 = "123"
	num2 = "456"
	expectation = "56088"
	applyTest(num1, num2, expectation, t)

	num1 = "1239"
	num2 = "45699"
	expectation = fmt.Sprint(1239 * 45699)
	applyTest(num1, num2, expectation, t)
}

func TestConvertToIngeter(t *testing.T) {
	t.Fatalf("first element is %v", int(rune("123"[2])-rune('0')))
}

func TestMultiplyOneDigit(t *testing.T) {
	if result := multiplyOneDigit("1234", rune("8"[0])); result != fmt.Sprint(1234*8) {
		t.Errorf("expecting %v, got %v", 1234*8, result)
	}
	if result := multiplyOneDigit("45699", rune("9"[0])); result != fmt.Sprint(45699*9) {
		t.Errorf("expecting %v, got %v", 45699*9, result)
	}
}

func TestAddStrings(t *testing.T) {
	if result := addStrings("1234", "9999"); result != fmt.Sprint(1234+9999) {
		t.Errorf("expecting %v, got %v", 1234+9999, result)
	}
	if result := addStrings("1234666", "9999"); result != fmt.Sprint(1234666+9999) {
		t.Errorf("expecting %v, got %v", 1236664+9999, result)
	}
}
