package main

type DomesticBuilder struct {
	title string
	description string
}

func newDomesticBuilder() *DomesticBuilder {
	return &DomesticBuilder{}
}

func (b *DomesticBuilder) setTitle (title string) {
	b.title = title
}

func (b *DomesticBuilder) setDescription (description string) {
	b.description = description
}

func (b *DomesticBuilder) getTask() Task {
	return Task{
		title: b.title,
		description: b.description,
		currentState: &NotStartedState{},
	}
}