package main

type WorkBuilder struct {
	id int
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

func (b *WorkBuilder) setID (id int) {
	b.id = id
}

func (b *WorkBuilder) getTask() *Task {
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