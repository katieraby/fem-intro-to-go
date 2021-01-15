package utils

import (
	"fem-intro-to-go/05_toolkit/code/exercise_5a_solution/utils"
	"testing"
)

func TestMakeExcited(t *testing.T) {
	expected := "OMG SO EXCITING!"
	actual := utils.MakeExcited("omg so exciting")
	if actual != expected {
		t.Errorf("Average was incorrect! Expected: %s, Actual: %s", expected, actual)
	}
}
