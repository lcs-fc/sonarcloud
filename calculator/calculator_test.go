package calculator

import "testing"

func TestSum(t *testing.T) {
	result := Sum(2, 3)

	if result != 5 {
		t.Error("O resultado deveria ser 5")
	}
}
