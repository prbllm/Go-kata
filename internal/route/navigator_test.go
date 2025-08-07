package route

import (
	"reflect"
	"testing"
)

func TestRouteGetResult(t *testing.T) {
	cases := []struct {
		name   string
		rows   int
		cols   int
		lines  []string
		expect []string
	}{
		{
			name: "Set #1",
			rows: 3,
			cols: 3,
			lines: []string{
				"B..",
				".#.",
				"..A",
			},
			expect: []string{
				"B..",
				".#.",
				"..A",
			},
		},
		{
			name: "Set #2",
			rows: 5,
			cols: 5,
			lines: []string{
				".....",
				".#A#.",
				"...B.",
				".#.#.",
				".....",
			},
			expect: []string{
				"aaa..",
				".#A#.",
				"...Bb",
				".#.#b",
				"....b",
			},
		},
		{
			name: "Set #3",
			rows: 7,
			cols: 9,
			lines: []string{
				".........",
				".#.#.#.#.",
				"..AB.....",
				".#.#.#.#.",
				".........",
				".#.#.#.#.",
				".........",
			},
			expect: []string{
				"aaa......",
				".#a#.#.#.",
				"..ABbbbbb",
				".#.#.#.#b",
				"........b",
				".#.#.#.#b",
				"........b",
			},
		},
		{
			name: "Set #4",
			rows: 5, cols: 5,
			lines: []string{
				"A....",
				".#.#.",
				"..#..",
				".#.#.",
				"....B",
			},
			expect: []string{
				"A....",
				".#.#.",
				"..#..",
				".#.#.",
				"....B",
			},
		},
		{
			name: "Set #5",
			rows: 5, cols: 7,
			lines: []string{
				".......",
				".#.#.#.",
				".A..B..",
				".#.#.#.",
				".......",
			},
			expect: []string{
				"a......",
				"a#.#.#.",
				"aA..B..",
				".#.#b#.",
				"....bbb",
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			navigator := NewNavigator(tc.rows, tc.cols)
			for i, ln := range tc.lines {
				if err := navigator.ParseLine(ln, i); err != nil {
					t.Fatalf("unexpected parse error: %v", err)
				}
			}
			navigator.FindRoute()
			got := navigator.GetResult()
			if !reflect.DeepEqual(got, tc.expect) {
				t.Errorf("GetResult() = %v, want %v", got, tc.expect)
			}
		})
	}
}
