package main

import (
  "testing"
)

func TestParseLines(t *testing.T) {
	tests := []struct {
		prevLine string
		currLine string
		result   string
	}{
		{
			prevLine: "line1",
			currLine: "line1",
			result:   "",
		},
		{
			prevLine: "line1",
			currLine: "line2",
			result:   "line1",
		},
		{
			prevLine: "",
			currLine: "line1",
			result:   "",
		},
	}

	for _, tt := range tests {
		if tt.result != parseLines(&tt.currLine, &tt.prevLine) {
			t.Errorf("Err: currLine: %s, prevLine: %s, result: %s", tt.currLine, tt.prevLine, tt.result)
		}
	}
}
