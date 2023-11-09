package main

type DomesticBuilder struct {
	id int
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

func (b *DomesticBuilder) setID (id int) {
	b.id = id
}

func (b *DomesticBuilder) getTask() *Task {
	t := &Task{
		id: b.id,
		title: b.title,
		description: b.description,
	}

	notStartedState := &NotStartedState{
		task: t,
	}
	inProgressState := &InProgressState{
		task: t,
	}
	finishedState := &FinishedState{
		task: t,
	}

	t.setState(notStartedState)
	t.notStarted = notStartedState
	t.inProgress = inProgressState
	t.finished = finishedState

	return t
}