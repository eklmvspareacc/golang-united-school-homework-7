package coverage

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

func TestPeople_Len(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		p := People{}
		got := p.Len()
		assert.Equal(t, 0, got)
	})
	t.Run("populated", func(t *testing.T) {
		p := People{{},{}}
		got := p.Len()
		assert.Equal(t, 2, got)
	})
}

func TestPeople_Less(t *testing.T) {
	t.Run("equal", func(t *testing.T) {
		p := People{{"Aaron", "Aaronson", time.Time{}}, 
					{"Aaron", "Aaronson", time.Time{}},}
		got := p.Less(0, 1)
		assert.False(t, got)
	})
	t.Run("only different last name, where 0 < 1", func(t *testing.T) {
		p := People{{"Aaron", "Aaronson", time.Time{}},
					{"Aaron", "Baronson", time.Time{}},}
		got := p.Less(0, 1)
		assert.True(t, got)
	})
	t.Run("only different last name, where 0 > 1", func(t *testing.T) {
		p := People{{"Aaron", "Caronson", time.Time{}},
					{"Aaron", "Aaronson", time.Time{}},}
		got := p.Less(0, 1)
		assert.False(t, got)
	})
	t.Run("only different first name, where 0 < 1", func(t *testing.T) {
		p := People{{"Aaron", "Aaronson", time.Time{}},
					{"Baron", "Aaronson", time.Time{}},}
		got := p.Less(0, 1)
		assert.True(t, got)
	})
	t.Run("only different first name, where 0 > 1", func(t *testing.T) {
		p := People{{"Caron", "Aaronson", time.Time{}},
					{"Aaron", "Aaronson", time.Time{}},}
		got := p.Less(0, 1)
		assert.False(t, got)
	})
	t.Run("only different birthday, where 0 younger than 1", func(t *testing.T) {
		p := People{{"Aaron", "Aaronson", time.Time{}.Add(time.Hour * 10)},
					{"Aaron", "Aaronson", time.Time{}},}
		got := p.Less(0, 1)
		assert.True(t, got)
	})
	t.Run("only different birthday, where 0 older than 1", func(t *testing.T) {
		p := People{{"Aaron", "Aaronson", time.Time{}},
					{"Aaron", "Aaronson", time.Time{}.Add(time.Hour * 10)},}
		got := p.Less(0, 1)
		assert.False(t, got)
	})
}

func TestPeople_Swap(t *testing.T) {
	t.Run("swap two indexes", func(t *testing.T) {
		p 	:= 	People{{firstName: "Aaron"},
					   {firstName: "Baron"},}
		exp := 	People{{firstName: "Baron"},
					   {firstName: "Aaron"},}
		p.Swap(0, 1)
		assert.Equal(t, exp, p)
	})
}

func TestMatrix_New(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		str	:= "1 2 3\n4 5 6\n7 8 9"
		exp := &Matrix{3, 3, []int{1,2,3,4,5,6,7,8,9}}
		got, err := New(str)
		if assert.NoError(t, err) {
			assert.Equal(t, exp, got)
		}
	})
	t.Run("different row lenght", func(t *testing.T) {
		str := "1 2 3\n4 5 6 7\n7 8 9 10 11"
		got, err := New(str)
		if assert.EqualError(t, err, "Rows need to be the same length") {
			assert.Nil(t, got)
		}
	})
	t.Run("invalid characters", func(t *testing.T) {
		str := "1 2 3:\n4a 5 6\n7 8sd 9"
		got, err := New(str)
		if assert.ErrorIs(t, err, strconv.ErrSyntax ) {
			assert.Nil(t, got)
		}
	})
}

func TestMatrix_Row(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		m := Matrix{0, 0, []int{}}
		assert.Empty(t, m.Rows())
	})
	t.Run("one row, one column", func(t *testing.T) {
		m := Matrix{1, 1, []int{1}}
		exp	:= [][]int{{1}}
		assert.Equal(t, exp, m.Rows())
	})
	t.Run("one row, many columns", func(t *testing.T) {
		m := Matrix{1, 3, []int{1, 2, 3}}
		exp := [][]int{{1, 2, 3}}
		assert.Equal(t, exp, m.Rows())
	})
	t.Run("many rows, one column", func(t *testing.T) {
		m := Matrix{3, 1, []int{1, 2, 3}}
		exp := [][]int{{1},{2},{3}}
		assert.Equal(t, exp, m.Rows())
	})
	t.Run("many rows, many columns", func(t *testing.T) {
		m := Matrix{3, 3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}
		exp := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		assert.Equal(t, exp, m.Rows())
	})
}

func TestMatrix_Cols(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		m := Matrix{0, 0, []int{}}
		assert.Empty(t, m.Cols())
	})
	t.Run("one row, one column", func(t *testing.T) {
		m := Matrix{1, 1, []int{1}}
		exp := [][]int{{1}}
		assert.Equal(t, exp, m.Cols())
	})
	t.Run("one row, many columns", func(t *testing.T) {
		m := Matrix{1, 3, []int{1, 2, 3}}
		exp := [][]int{{1}, {2}, {3}}
		assert.Equal(t, exp, m.Cols())
	})
	t.Run("many rows, one column", func(t *testing.T) {
		m := Matrix{3, 1, []int{1, 2, 3}}
		exp := [][]int{{1, 2, 3}}
		assert.Equal(t, exp, m.Cols())
	})
	t.Run("many rows, many columns", func(t *testing.T) {
		m := Matrix{3, 3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}
		exp := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
		assert.Equal(t, exp, m.Cols())
	})
}

func TestMatrix_Set(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		m := &Matrix{0, 0, []int{}}
		exp := &Matrix{0, 0, []int{}}
		assert.False(t, m.Set(0, 0, 0))
		assert.Equal(t, exp, m)
	})
	t.Run("negative row", func(t *testing.T) {
		m := &Matrix{0, 0, []int{}}
		exp := &Matrix{0, 0, []int{}}
		assert.False(t, m.Set(-1, 0, 0))
		assert.Equal(t, exp, m)
	})
	t.Run("negative column", func(t *testing.T) {
		m := &Matrix{0, 0, []int{}}
		exp := &Matrix{0, 0, []int{}}
		assert.False(t, m.Set(0, -1, 0))
		assert.Equal(t, exp, m)
	})
	t.Run("row too big", func(t *testing.T) {
		m := &Matrix{0, 0, []int{}}
		exp := &Matrix{0, 0, []int{}}
		assert.False(t, m.Set(3, 0, 0))
		assert.Equal(t, exp, m)
	})
	t.Run("column too big", func(t *testing.T) {
		m := &Matrix{0, 0, []int{}}
		exp := &Matrix{0, 0, []int{}}
		assert.False(t, m.Set(0, 3, 0))
		assert.Equal(t, exp, m)
	})
	t.Run("correct", func(t *testing.T) {
		m := &Matrix{3, 3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}
		exp := &Matrix{3, 3, []int{1, 2, 3, 4, 5, 10, 7, 8, 9}}
		assert.True(t, m.Set(1, 2, 10))
		assert.Equal(t, exp, m)
	})
}