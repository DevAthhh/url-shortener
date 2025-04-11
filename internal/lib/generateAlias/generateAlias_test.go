package generateAlias_test

import (
	"testing"

	utils "github.com/DevAthhh/url-shortener/internal/lib/generateAlias"
)

func TestGenerateAlias(t *testing.T) {
	var testcases = []struct {
		Name  string
		Input int
		Want  int
	}{
		{"Test 1 - 'Zero len'", 0, 0},
		{"Test 2 - 'size = 5'", 5, 5},
		{"Test 3 - 'size = 1000'", 1000, 1000},
		{"Test 4 - 'size = 1'", 1, 1},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			result := utils.GenerateStr(tc.Input)
			if len(result) != tc.Want {
				t.Errorf("got %v, want %v", len(result), tc.Want)
			}
		})
	}
}
