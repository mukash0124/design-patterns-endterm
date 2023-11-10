package main

import "fmt"

type Task struct {
	id int
	title        string
	description  string
	currentState State

	notStarted State
	inProgress State
	finished State
}

func (t *Task) setState(s State) {
	t.currentState = s
}

func (t *Task) createMemento() Memento {
	return Memento{title: t.title, description: t.description}
}

func (t *Task) restoreMemento(m Memento) {
	t.title = m.getTitle()
	t.description = m.getDescription()
}

func (t *Task) stringifyTask() string {
	return fmt.Sprintf("%d. %s - %s", t.id + 1, t.title, t.description)
}