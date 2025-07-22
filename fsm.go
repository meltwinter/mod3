package fsm

type State string
type States map[State]int
type FinalStates States
type Inputs map[rune]struct{}
type Transition struct {
	Start State
	Input rune
}
type Transitions map[Transition]State

func CreateMod3FSM() (States, Inputs, State, FinalStates, Transitions) {

	Q := States{
		State("S0"): 0,
		State("S1"): 1,
		State("S2"): 2,
	}
	Sigma := Inputs{
		'0': {},
		'1': {},
	}
	q0 := State("S0")
	F := FinalStates(Q)
	delta := Transitions{
		Transition{Start: State("S0"), Input: '0'}: State("S0"),
		Transition{Start: State("S0"), Input: '1'}: State("S1"),
		Transition{Start: State("S1"), Input: '0'}: State("S2"),
		Transition{Start: State("S1"), Input: '1'}: State("S0"),
		Transition{Start: State("S2"), Input: '0'}: State("S1"),
		Transition{Start: State("S2"), Input: '1'}: State("S2"),
	}
	return Q, Sigma, q0, F, delta

}

//functions for generating other FSM below
