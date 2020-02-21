package solution2019

import "testing"

func TestDay2(t *testing.T) {
	want := make(map[string]int)
	want["Part 1"] = 5434663
	want["Part 2"] = 4559

	got := Day2()
	for k := range want {
		if got[k] != want[k] {
			t.Errorf("Day2(%s) == %d, want %d", k, got[k], want[k])
		}
	}
}