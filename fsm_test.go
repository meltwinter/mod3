package fsm

import (
	"testing"
)

func TestMod3FSM(t *testing.T) {
	mod3 := CreateMod3FSM()
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"110 get 0", "110", 0},
		{"1010 get 1", "1010", 1},
		{"1101 get 1", "1101", 1},
		{"1110 get 2", "1110", 2},
		{"1111 get 0", "1111", 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			prevState := mod3.Q0
			var finalState *State
			for _, r := range test.input {
				if _, ok := mod3.Sigma[r]; !ok {
					t.Errorf("test: %s, invalid input :%v", test.name, r)
				}
				Output := mod3.Delta(prevState, r)
				if Output == nil {
					t.Errorf("test: %s, input %v is not allowed for State %v", test.name, r, prevState)
				}

				prevState = Output
				finalState = Output
			}
			result, ok := mod3.F[finalState.Name]
			if !ok {
				t.Errorf("impossible finalState: %v", finalState)
			}
			if result.Output != test.expected {
				t.Errorf("test: %s, wrong result. expected: %d, actual: %d", test.name, test.expected, result.Output)
			}
		})
	}
}
