package calculator

import "testing"

func TestSum(t *testing.T) {
	result := Sum(2, 3)

	if result != 5 {
		t.Error("O resultado deveria ser 5")
	}
}


func TestSub(t *testing.T) {
	result := Sub(3, 3)

	if result != 0 {
		t.Error("O resultado deveria ser 0")
	}
}

func TestTimes(t *testing.T) {
	result := Times(2, 3)

	if result != 6 {
		t.Error("O resultado deveria ser 6")
	}
}

func TestSumX(t *testing.T) {
	result := SumX(2, 3)

	if result != 7 {
		t.Error("O resultado deveria ser 7")
	}
}