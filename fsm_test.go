package fsm

import (
	"testing"
)

func TestMod3FSM(t *testing.T) {
	Q, inputs, q0, F, sigma := CreateMod3FSM()
	prevState := q0
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"110 get 0", "110", 0},
		{"1010 get 1", "1010", 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var finalState State
			for _, r := range test.input {
				if _, ok := inputs[r]; !ok {
					t.Errorf("invalid input :%v", r)
				}
				transition := Transition{
					Start: prevState,
					Input: r,
				}
				Output, ok := sigma[transition]
				if !ok {
					t.Errorf("impossible transaction: %v", transition)
				}
				//check the output is valid
				if _, ok = Q[Output]; !ok {
					t.Errorf("invalid state: %v", Output)
				}

				prevState = Output
				finalState = Output
			}
			result, ok := F[finalState]
			if !ok {
				t.Errorf("impossible finalState: %v", finalState)
			}
			if result != test.expected {
				t.Errorf("wrong result. expected: %d, actual: %d", 0, result)
			}
		})
	}
}
