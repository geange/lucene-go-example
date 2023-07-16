package main

import (
	"fmt"
	"log"

	"github.com/geange/lucene-go/core/util/automaton"
)

func main() {
	auto := automaton.NewAutomaton()

	state1 := auto.CreateState()
	state2 := auto.CreateState()
	state3 := auto.CreateState()
	state4 := auto.CreateState()

	auto.AddTransitionLabel(state1, state2, int(1))
	auto.AddTransitionLabel(state1, state3, int(2))
	auto.AddTransitionLabel(state2, state4, int(3))
	auto.AddTransitionLabel(state3, state4, int(4))

	state := auto.GetNumTransitionsWithState(state1)

	step := auto.Step(state1, int(1))
	if step != state2 {
		log.Panicf("step = %d", step)
	}
	fmt.Println(state)
}
