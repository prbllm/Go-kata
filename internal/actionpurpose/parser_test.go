package actionpurpose

import (
	"reflect"
	"testing"
)

func TestParserGetResult(t *testing.T) {
	cases := []struct {
		name   string
		lines  []string
		expect []string
	}{
		{
			name: "Set #1",
			lines: []string{
				"Andrew: Boris is meowing!",
				"Boris: I am not meowing!",
				"Kate: Andrew is meowing!",
				"Kate: Boris is not meowing!",
				"Kate: I am meowing!",
			},
			expect: []string{"Kate"},
		},
		{
			name: "Set #2",
			lines: []string{
				"Sedan: I am hungry!",
				"Ivan: I am hungry!",
			},
			expect: []string{"Ivan", "Sedan"},
		},
		{
			name: "Set #3",
			lines: []string{
				"I: I am serious!",
				"H: I is serious!",
				"H: I am serious!",
			},
			expect: []string{"I"},
		},
		{
			name: "Set #4",
			lines: []string{
				"A: B is not x!",
				"B: C is not x!",
				"C: B is not x!",
			},
			expect: []string{"A"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := NewActionPurposeParser()
			for _, ln := range tc.lines {
				if err := p.ParseLine(ln); err != nil {
					t.Fatalf("unexpected parse error: %v", err)
				}
			}
			got := p.GetResult()
			if !reflect.DeepEqual(got, tc.expect) {
				t.Errorf("GetResult() = %v, want %v", got, tc.expect)
			}
		})
	}
}

func TestParserGetResult_BadInput(t *testing.T) {
	p := NewActionPurposeParser()

	bad := []string{
		"Boris is meowing!",
		"Kate: Andrew is jumping",
		"Jack: ???",
		"",
	}

	for _, ln := range bad {
		if err := p.ParseLine(ln); err == nil {
			t.Errorf("expected error for %q, got nil", ln)
		}
	}
}
