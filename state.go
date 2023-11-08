package main

type State interface {
	edit(title, description string) error
	start() error
	finish() error
}

