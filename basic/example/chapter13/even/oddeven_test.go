package even

import "testing"

func TestEven(t *testing.T) {
	if Even(10) {
		t.Log("Everything OK: 10 is even, just a test to see failed output!")
		t.Fail()
	}
}

func TestOdd(t *testing.T) {
	if !Odd(11) {
		t.Log("11 must be odd!")
		t.Fail()
	}

	if !Odd(10) {
		t.Log("10 is not odd!")
		t.Fail()
	}
}
