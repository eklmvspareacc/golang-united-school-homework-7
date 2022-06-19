package coverage

import (
	"os"
	"testing"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func Test_Len(t *testing.T) {
	tData := map[string]struct {
		p People
		expected int
	}{
		"empty": {People{}, 0},
		"full": {People{{}, {}}, 2},
	}
	for name, tcase := range tData {
		t.Run(name, func(t *testing.T) {
			got := tcase.p.Len()
			if got != tcase.expected {
				t.Errorf("[%s] expected: %d, got: %d", name, tcase.expected, got)
			}
		})
	}
}