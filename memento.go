package main

type IMemento interface {
	getTitle() string
	getDescription() string
}

type Memento struct {
	title string
	description string
}

func (m *Memento) getTitle () string {
	return m.title
}

func (m *Memento) getDescription () string {
	return m.description
}