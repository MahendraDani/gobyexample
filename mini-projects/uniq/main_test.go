package main

import (
	"testing"
)

func TestParseLines(t *testing.T) {
	tests := []struct {
		name      string
		prevLine  string
		currLine  string
		isFirst   bool
		want      string
		wantPrev  string // What prevLine should be after the call
		wantFirst bool   // What isFirst should be after the call
	}{
		{
			name:      "duplicate lines",
			prevLine:  "line1",
			currLine:  "line1",
			isFirst:   false,
			want:      "",
			wantPrev:  "line1",
			wantFirst: false,
		},
		{
			name:      "different lines",
			prevLine:  "line1",
			currLine:  "line2",
			isFirst:   false,
			want:      "line1",
			wantPrev:  "line2",
			wantFirst: false,
		},
		{
			name:      "first line",
			prevLine:  "",
			currLine:  "line1",
			isFirst:   true,
			want:      "",
			wantPrev:  "line1",
			wantFirst: false,
		},
		{
			name:      "empty line after content",
			prevLine:  "line1",
			currLine:  "",
			isFirst:   false,
			want:      "line1",
			wantPrev:  "",
			wantFirst: false,
		},
		{
			name:      "duplicate empty lines",
			prevLine:  "",
			currLine:  "",
			isFirst:   false,
			want:      "",
			wantPrev:  "",
			wantFirst: false,
		},
		{
			name:      "content after empty line",
			prevLine:  "",
			currLine:  "line1",
			isFirst:   false,
			want:      "",
			wantPrev:  "line1",
			wantFirst: false,
		},
		{
			name:      "same content separated by empty line (prev is empty)",
			prevLine:  "",
			currLine:  "banana",
			isFirst:   false,
			want:      "",
			wantPrev:  "banana",
			wantFirst: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prevLine := tt.prevLine
			currLine := tt.currLine
			isFirst := tt.isFirst

			got := parseLines(&currLine, &prevLine, &isFirst)

			if got != tt.want {
				t.Errorf("parseLines() result = %q, want %q", got, tt.want)
			}

			if prevLine != tt.wantPrev {
				t.Errorf("parseLines() prevLine after call = %q, want %q", prevLine, tt.wantPrev)
			}

			if isFirst != tt.wantFirst {
				t.Errorf("parseLines() isFirst after call = %v, want %v", isFirst, tt.wantFirst)
			}
		})
	}
}
