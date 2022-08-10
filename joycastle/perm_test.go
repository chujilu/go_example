package main

import "testing"

func TestPerm(t *testing.T) {
	test := map[string]int{
		"ABC":  6,
		"AB":   2,
		"ABCD": 24,
	}
	for s, n := range test {
		r := make([]string, 0, n)
		Perm([]rune(s), func(runes []rune) {
			r = append(r, string(runes))
		})
		if len(r) != n {
			t.Failed()
		}
	}
	t.Log("Perm pass")
}

func FuzzPerm(f *testing.F) {
	test := map[string]int{
		"ABC":  6,
		"AB":   2,
		"ABCD": 24,
	}
	for s, _ := range test {
		f.Add(s)
	}
	var getNum func(i int) int
	getNum = func(i int) int {
		if i == 0 {
			return 1
		}
		return i * getNum(i-1)
	}
	f.Fuzz(func(t *testing.T, s string) {
		t.Log(s)
		r := make([]string, 0, getNum(len(s)))
		Perm([]rune(s), func(runes []rune) {
			r = append(r, string(runes))
		})
		if len(r) != getNum(len(s)) {
			t.Failed()
		}
	})
}
