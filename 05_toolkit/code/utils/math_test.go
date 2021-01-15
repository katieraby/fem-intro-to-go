package utils

import (
	"fem-intro-to-go/05_toolkit/code/exercise_5a_solution/utils"
	"testing"
)

func TestAdd(t *testing.T) {
input := utils.Add(4, 2, 6, 8)
expected := 20

	if input != expected {
		t.Errorf("Input to add function gives us actual: %d, when looking for %d", input, expected)
	}
}

