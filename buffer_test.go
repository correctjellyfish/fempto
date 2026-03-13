package main

import (
	"slices"
	"testing"
)

func TestRowSplit(t *testing.T) {
	testRow := newRow("Hello", 0)
	testSplitRow := testRow.split(4)
	wantRemain := newRow("Hell", 0)
	wantSplit := newRow("o", 0)
	if !slices.Equal(testRow.content, wantRemain.content) || !slices.Equal(testSplitRow.content, wantSplit.content) {
		t.Errorf(`SplitRow("Hello",4) = %q, %q, want match for %q, %q`, testRow.content, testSplitRow.content, wantRemain.content, wantSplit.content)
	}
}
