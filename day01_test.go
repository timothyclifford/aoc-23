package main

import (
	"strings"
	"testing"
)

func TestRunDay01Part1(t *testing.T) {

	input := `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`

	lines := strings.Split(input, "\n")

	got := runDay01(lines)
	want := 142

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestRunDay01Part2(t *testing.T) {

	input := `two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen`

	lines := strings.Split(input, "\n")

	got := runDay01(lines)
	want := 281

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestFilterLineForWords(t *testing.T) {
	input := "ninefive5threetwo73pxjfive"

	got := filterLineForWords(input)
	want := map[int]int{0: 9, 4: 5, 9: 3, 14: 2, 22: 5}

	for k, v := range want {
		if v != got[k] {
			t.Errorf("got %d, wanted %d", got[k], v)
		}
	}
}

func TestFilterLineForDigits(t *testing.T) {
	input := "1abc2"

	got := filterLineForDigits(input)
	want := map[int]int{0: 1, 4: 2}

	for k, v := range want {
		if v != got[k] {
			t.Errorf("got %d, wanted %d", v, want[k])
		}
	}
}

func TestGetFirstLastValues(t *testing.T) {
	input := map[int]int{0: 4, 2: 5, 5: 8}

	first, last := getFirstLastValues(input)
	wantFirst := 4
	wantLast := 8

	if first != wantFirst {
		t.Errorf("got %d, wanted %d", first, wantFirst)
	}

	if last != wantLast {
		t.Errorf("got %d, wanted %d", last, wantLast)
	}
}
