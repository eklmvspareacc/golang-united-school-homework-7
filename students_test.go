package coverage

import (
	"os"
	"reflect"
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

func Test_People_Len(t *testing.T) {
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

func Test_People_Less(t *testing.T) {
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

func Test_People_Swap(t *testing.T) {
	tData := map[string]struct {
		p People
		i, j int
		expected People
	}{
		"swap two indexes":
			{
				People{
					{firstName: "Aaron"},
					{firstName: "Baron"},
				},
				0, 1,
				People{
					{firstName: "Baron"},
					{firstName: "Aaron"},
				},
			},
	}
	for name, tcase := range tData {
		t.Run(name, func(t *testing.T) {
			tcase.p.Swap(tcase.i, tcase.j)
			if len(tcase.p) != len(tcase.expected) {
				t.Errorf("[%s] expected slice len: %d, got: %d", name, len(tcase.expected), len(tcase.p))
			}
			for i, v := range tcase.p {
				if v != tcase.expected[i] {
					t.Errorf("[%s] expected: %+v, got: %+v", name, tcase.expected, tcase.p)
				} 
			}
		})
	}
}

func Test_Matrix_New(t *testing.T) {
	tData := map[string]struct {
		str string
		expected *Matrix
		expectErr bool
	}{
		"correct":
			{
				"1 2 3\n4 5 6\n7 8 9",
				&Matrix{3, 3, []int{1,2,3,4,5,6,7,8,9}},
				false,
			},
		"different row lenght":
			{
				"1 2 3\n4 5 6 7\n7 8 9 10 11",
				&Matrix{},
				true,
			},
		"invalid characters":
			{
				"1 2 3:\n4a 5 6\n7 8sd 9",
				&Matrix{},
				true,
			},
	}
	for name, tcase := range tData {
		t.Run(name, func(t *testing.T) {
			got, err := New(tcase.str) 
			if tcase.expectErr && err == nil {
				t.Errorf("[%s] no expected error", name)
				return
			}
			if !tcase.expectErr && err != nil {
				t.Errorf("[%s] unexpected error: %v", name, err)
				return
			}
			if tcase.expectErr && err != nil {
				return
			}
			if !reflect.DeepEqual(got, tcase.expected) {
				t.Errorf("[%s] expected: %+v, got: %+v", name, tcase.expected, got)
			}
		})
	}
}

func Test_Matrix_Row(t *testing.T) {
	tData := map[string]struct {
		m Matrix
		expected [][]int
	}{
		"empty": 
			{
				Matrix {0, 0, []int{}},
				[][]int{},
			},
		"one row, one column":
			{
				Matrix {1, 1, []int{1}},
				[][]int{{1}},
			},
		"one row, many columns":
			{
				Matrix {1, 3, []int{1, 2, 3}},
				[][]int{{1, 2, 3}},
			},
		"many rows, one column":
			{
				Matrix {3, 1, []int{1, 2, 3}},
				[][]int{{1},{2},{3}},
			},
		"many rows, many columns":
			{	
				Matrix {3, 3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
				[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
	}
	for name, tcase := range tData {
		t.Run(name, func(t *testing.T) {
			got := tcase.m.Rows()
			if !reflect.DeepEqual(got, tcase.expected) {
				t.Errorf("[%s] expected: %v, got: %v", name, tcase.expected, got)
			}
		})
	}
}