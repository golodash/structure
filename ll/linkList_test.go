package ll

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

func testFun() bool {
	return true
}

func TestLinkList(t *testing.T) {
	cases := [7]TCase{
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
	}

	ll := New[int](map[string]interface{}{"test": testFun})

	for k, c := range cases {
		if c.name == "AddFirst" {
			for i := 0; i < c.iter; i++ {
				temp := c.input
				ll.AddFirst(&temp)
			}
		} else if c.name == "RemoveFirst" {
			for i := 0; i < c.iter; i++ {
				ll.RemoveFirst()
			}
		} else if c.name == "AddLast" {
			for i := 0; i < c.iter; i++ {
				temp := c.input
				ll.AddLast(&temp)
			}
		} else if c.name == "RemoveLast" {
			for i := 0; i < c.iter; i++ {
				ll.RemoveLast()
			}
		} else if strings.Contains(c.name, "Remove:") {
			num, _ := strconv.Atoi(c.name[7:])
			for i := 0; i < c.iter; i++ {
				ll.Remove(num)
			}
		} else if strings.Contains(c.name, "Add:") {
			num, _ := strconv.Atoi(c.name[4:])
			for i := 0; i < c.iter; i++ {
				temp := c.input
				ll.Add(num, &temp)
			}
		}

		slice := ll.ReturnAsSlice()

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

	if res, err := ll.Run("test"); res == nil || err != nil {
		t.Errorf("on calling test something happened:\nerr => %s", err)
	}
}
