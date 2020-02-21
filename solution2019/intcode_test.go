package solution2019

import "testing"

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestIntcodeDay2Part1(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{1,9,10,3,2,3,11,0,99,30,40,50}, []int{3500,9,10,70,2,3,11,0,99,30,40,50}},
		{[]int{1,0,0,0,99}, []int{2,0,0,0,99}},
		{[]int{2,3,0,3,99}, []int{2,3,0,6,99}},
		{[]int{2,4,4,5,99,0}, []int{2,4,4,5,99,9801}},
		{[]int{1,1,1,4,99,5,6,0,99}, []int{30,1,1,4,2,5,6,0,99}},
	}

	for _, c := range cases {
		_, state := Intcode(c.in)
		if !equal(state, c.want) {
			t.Errorf("\nGot %v\nWant %v", state, c.want)
		}
	}
}