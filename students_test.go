package coverage

import (
	"os"
	"testing"
	"time"
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

func Test_Less(t *testing.T) {
	tData := map[string]struct {
		p People
		expected bool
	}{
		"equal": 
			{
				People{
						{"Aaron", "Aaronson", time.Time{}},
						{"Aaron", "Aaronson", time.Time{}},
					},
					false,
			},
		"only different last name, where 0 < 1":
			{
				People{
					{"Aaron", "Aaronson", time.Time{}},
					{"Aaron", "Baronson", time.Time{}},
				},
				true,
			},
		"only different last name, where 0 > 1":
			{
				People{
					{"Aaron", "Caronson", time.Time{}},
					{"Aaron", "Aaronson", time.Time{}},
				},
				false,
			},	
		"only different first name, where 0 < 1":
			{
				People{
					{"Aaron", "Aaronson", time.Time{}},
					{"Baron", "Aaronson", time.Time{}},
				},
				true,
			},	
		"only different first name, where 0 > 1":
			{
				People{
					{"Caron", "Aaronson", time.Time{}},
					{"Aaron", "Aaronson", time.Time{}},
				},
				false,
			},
		"only different birthday, where 0 < 1":
			{
				People{
					{"Aaron", "Aaronson", time.Time{}},
					{"Aaron", "Aaronson", time.Time{}.Add(time.Hour * 10)},
				},
				false,
			},
		"only different birthday, where 0 > 1":
			{
				People{
					{"Aaron", "Aaronson", time.Time{}.Add(time.Hour * 10)},
					{"Aaron", "Aaronson", time.Time{}},
				},
				true,
			},
	}
	for name, tcase := range tData {
		t.Run(name, func(t *testing.T) {
			got := tcase.p.Less(0, 1)
			if got != tcase.expected {
				t.Errorf("[%s] expected: %t, got: %t", name, tcase.expected, got)
			}
		})
	}
}