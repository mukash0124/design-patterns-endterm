package main

import "fmt"

type FinishedState struct {
	task *Task
}

func (s *FinishedState) edit(title, description string) error {
	return fmt.Errorf("You can't edit finished task!")
}

func (s *FinishedState) start() error {
	return fmt.Errorf("You can't start finished task!")
}

func (s *FinishedState) finish() error {
	return fmt.Errorf("Task is already finished!")
}