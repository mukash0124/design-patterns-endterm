package main

type Task struct {
	title string
	description string
	currentState State
}

func (t *Task) setState(s State) {
    t.currentState = s
}

func (t *Task) createMemento() *Memento {
	return &Memento{title: t.title, description: t.description}
}

func (t *Task) restoreMemento(m *Memento) {
	t.title = m.title
	t.description = m.description
}