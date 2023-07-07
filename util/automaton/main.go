package main

import (
	"fmt"
	"github.com/geange/lucene-go/core/util/automaton"
	"log"
)

func main() {

	automaton := automaton.NewAutomaton()

	state1 := automaton.CreateState()
	state2 := automaton.CreateState()
	state3 := automaton.CreateState()
	state4 := automaton.CreateState()

	//automaton.SetAccept(state1, true)
	//automaton.SetAccept(state2, true)
	//automaton.SetAccept(state3, true)
	//automaton.SetAccept(state4, true)

	automaton.AddTransitionLabel(state1, state2, int(1))
	automaton.AddTransitionLabel(state1, state3, int(2))
	automaton.AddTransitionLabel(state2, state4, int(3))
	automaton.AddTransitionLabel(state3, state4, int(4))

	state := automaton.GetNumTransitionsWithState(state1)

	step := automaton.Step(state1, int(1))
	if step != state2 {
		log.Panicf("step = %d", step)
	}
	fmt.Println(state)
}
