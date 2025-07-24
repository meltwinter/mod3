package fsm

type State struct {
	Name   string
	Output int
}

type States map[string]*State
type FinalStates States
type Inputs map[rune]struct{}
type TransitionFunc func(*State, rune) *State

type FSM struct {
	Q     States
	Sigma Inputs
	Q0    *State
	F     FinalStates
	Delta TransitionFunc
}

func CreateMod3FSM() *FSM {

	Q := States{
		"S0": &State{"S0", 0},
		"S1": &State{"S1", 1},
		"S2": &State{"S2", 2},
	}
	Sigma := Inputs{
		'0': {},
		'1': {},
	}
	q0 := Q["S0"]
	F := FinalStates(Q)
	delta := TransitionFunc(func(s *State, r rune) *State {
		switch {
		case s == Q["S0"] && r == '0':
			return Q["S0"]
		case s == Q["S0"] && r == '1':
			return Q["S1"]
		case s == Q["S1"] && r == '0':
			return Q["S2"]
		case s == Q["S1"] && r == '1':
			return Q["S0"]
		case s == Q["S2"] && r == '0':
			return Q["S1"]
		case s == Q["S2"] && r == '1':
			return Q["S2"]
		default:
			return nil
		}
	})
	return &FSM{Q, Sigma, q0, F, delta}

}

//functions for generating other FSM below
