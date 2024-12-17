package main

import (
	"testing"
)

func TestHorizontalInRow(t *testing.T) {
	bots := []point{
		{1, 1},
		{2, 1},
		{3, 1},
	}
	if !horizontalInRow(bots, 3) {
		t.Fatal("failed to detect 3 in a row")
	}
	bots = []point{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 1},
	}
	if !horizontalInRow(bots, 3) {
		t.Fatal("failed to detect 3 in a row")
	}
	bots = []point{
		{1, 1},
		{2, 1},
		{3, 1},
		{0, 0},
	}
	if !horizontalInRow(bots, 3) {
		t.Fatal("failed to detect 3 in a row")
	}
	bots = []point{
		{1, 1},
		{1, 1},
		{2, 1},
		{3, 1},
	}
	if !horizontalInRow(bots, 3) {
		t.Fatal("failed to detect 3 in a row")
	}
}
