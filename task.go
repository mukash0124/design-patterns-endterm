package main

type Task struct {
	title string
	description string
	currentState State
}

func (t *Task) setState(s State) {
    t.currentState = s
}