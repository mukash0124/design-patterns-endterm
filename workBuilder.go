package main

type WorkBuilder struct {
	title string
	description string
}

func newWorkBuilder() *WorkBuilder {
	return &WorkBuilder{}
}

func (b *WorkBuilder) setTitle (title string) {
	b.title = title
}

func (b *WorkBuilder) setDescription (description string) {
	b.description = description
}

func (b *WorkBuilder) getTask() Task {
	return Task{
		title: b.title,
		description: b.description,
		currentState: &NotStartedState{},
	}
}