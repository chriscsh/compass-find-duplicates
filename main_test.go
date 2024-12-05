package main

import (
	"testing"
)

func TestShouldMatchHighAcc(t *testing.T) {
	matches := findDuplicates("assets/ShouldMatchHighAcc.csv")
	if len(matches) == 0 || matches[0].Accuracy != "High" {
		t.Fatalf("Accuracy should be High")
	}
}

func TestShouldMatchLowAcc(t *testing.T) {
	matches := findDuplicates("assets/ShouldMatchLowAcc.csv")
	if len(matches) == 0 || matches[0].Accuracy != "Low" {
		t.Fatalf("Accuracy should be Low")
	}
}

func TestShouldNotMatch(t *testing.T) {
	if matches := findDuplicates("assets/ShouldNotMatch.csv"); len(matches) != 0 {
		t.Fatalf("Should not match")
	}
}
