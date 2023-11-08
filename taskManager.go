package main

import "fmt"

type TaskManager struct {
	tasks map[Task]bool
}

var taskManagerInstance *TaskManager

func getInstance() *TaskManager {
	if taskManagerInstance == nil {
            fmt.Println("Creating single instance now.")
            taskManagerInstance = &TaskManager{}
    } else {
            fmt.Println("Single instance already created.")
	}

    return taskManagerInstance
}

func (tm *TaskManager) addTask (t Task) {
	tm.tasks[t] = true
}

func (tm *TaskManager) removeTask (t Task) {
	delete(tm.tasks, t)
}
