package array

import (
	"strconv"
	"strings"
	"testing"
)

type (
	TCase struct {
		name  string
		input int
		iter  int
		want  []int
	}
)

func testFun[T any](l *Array[T]) bool {
	return true
}

func TestArray(t *testing.T) {
	cases := [11]TCase{
		{
			name:  "Add:0",
			input: 1,
			iter:  3,
			want:  []int{1, 1, 1},
		},
		{
			name:  "Remove:0",
			input: 0,
			iter:  2,
			want:  []int{1},
		},
		{
			name:  "AddFirst",
			input: 2,
			iter:  3,
			want:  []int{2, 2, 2, 1},
		},
		{
			name:  "RemoveFirst",
			input: 0,
			iter:  1,
			want:  []int{2, 2, 1},
		},
		{
			name:  "AddLast",
			input: 3,
			iter:  5,
			want:  []int{2, 2, 1, 3, 3, 3, 3, 3},
		},
		{
			name:  "RemoveLast",
			input: 0,
			iter:  1,
			want:  []int{2, 2, 1, 3, 3, 3, 3},
		},
		{
			name:  "Remove:2",
			input: 0,
			iter:  2,
			want:  []int{2, 2, 3, 3, 3},
		},
		{
			name:  "Displace:0,2",
			input: 0,
			iter:  1,
			want:  []int{3, 2, 2, 3, 3},
		},
		{
			name:  "Displace:1,3",
			input: 0,
			iter:  3,
			want:  []int{3, 3, 2, 2, 3},
		},
		{
			name:  "AddLast",
			input: 1,
			iter:  1,
			want:  []int{3, 3, 2, 2, 3, 1},
		},
		{
			name:  "Displace:1,5",
			input: 0,
			iter:  3,
			want:  []int{3, 1, 2, 2, 3, 3},
		},
	}

	a, err := New[int](map[string]any{"test": testFun[int]})
	if err != nil {
		t.Errorf("on calling `New` function error happened:\nerr => %s", err.Error())
		return
	}

	for k, c := range cases {
		if c.name == "AddFirst" {
			for i := 0; i < c.iter; i++ {
				temp := c.input
				a.PushFirst(&temp)
			}
		} else if c.name == "RemoveFirst" {
			for i := 0; i < c.iter; i++ {
				a.PopFirst()
			}
		} else if c.name == "AddLast" {
			for i := 0; i < c.iter; i++ {
				temp := c.input
				a.PushLast(&temp)
			}
		} else if c.name == "RemoveLast" {
			for i := 0; i < c.iter; i++ {
				a.PopLast()
			}
		} else if strings.Contains(c.name, "Remove:") {
			num, _ := strconv.Atoi(c.name[7:])
			for i := 0; i < c.iter; i++ {
				a.Pop(num)
			}
		} else if strings.Contains(c.name, "Add:") {
			num, _ := strconv.Atoi(c.name[4:])
			for i := 0; i < c.iter; i++ {
				temp := c.input
				a.Push(&temp, num)
			}
		} else if strings.Contains(c.name, "Displace:") {
			numsSplit := strings.Split(c.name[9:], ",")
			num1, _ := strconv.Atoi(numsSplit[0])
			num2, _ := strconv.Atoi(numsSplit[1])
			for i := 0; i < c.iter; i++ {
				a.Displace(num1, num2)
			}
		}

		slice := a.ReturnAsSlice()

		if len(c.want) != len(slice) {
			t.Errorf("(%s, %d) => want = %v, got = %v", c.name, k, c.want, slice)
			return
		}

		for i := range c.want {
			if c.want[i] != slice[i] {
				t.Errorf("(%s, %d) => want = %v, got = %v", c.name, k, c.want, slice)
				return
			}
		}
	}

	if res, err := a.Run("test"); res == nil || err != nil {
		t.Errorf("on calling `Run` for `test` function error happened:\nerr => %s", err.Error())
		return
	}
}
