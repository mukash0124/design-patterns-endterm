package main

import "fmt"

type NotStartedState struct {
	task *Task
}

func (s *NotStartedState) edit(title, description string) error {
	s.task.title = title
	s.task.description = description
	return nil
}

func (s *NotStartedState) start() error {
	s.task.setState(&InProgressState {
		task : s.task,
	})
	return nil
}

func (s *NotStartedState) finish() error {
	return fmt.Errorf("start the task before finishing")
}
