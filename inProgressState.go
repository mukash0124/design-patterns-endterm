package main

import "fmt"

type InProgressState struct {
	task *Task
}

func (s *InProgressState) edit(title, description string) error {
	return fmt.Errorf("You can't edit the task in progress!")
}

func (s *InProgressState) start() error {
	return fmt.Errorf("Task is already in progress!")
}

func (s *InProgressState) finish() error {
	s.task.setState(&FinishedState {
		task : s.task,
	})
	return nil
}