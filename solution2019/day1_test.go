package solution2019

import "testing"

func TestCalcFuel(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, c := range cases {
		got := calcFuel(c.in)
		if got != c.want {
			t.Errorf("calcFuel(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestCalcFuelRec(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, c := range cases {
		got := calcFuelRec(c.in)
		if got != c.want {
			t.Errorf("calcFuelRec(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestDay1(t *testing.T) {
	want := make(map[string]int)
	want["Part 1"] = 3297896
	want["Part 2"] = 4943969

	got := Day1()
	for k := range want {
		if got[k] != want[k] {
			t.Errorf("Day1(%s) == %d, want %d", k, got[k], want[k])
		}
	}
}