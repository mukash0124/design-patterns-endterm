package main

import "fmt"

type TaskManager struct {
	domesticTasks map[Task]*Memento
	workTasks map[Task]*Memento
	builder IBuilder
}

var taskManagerInstance *TaskManager

func GetInstance() *TaskManager {
	if taskManagerInstance == nil {
            taskManagerInstance = &TaskManager{}
			taskManagerInstance.domesticTasks = make(map[Task]*Memento)
			taskManagerInstance.workTasks = make(map[Task]*Memento)
    } else {
            fmt.Println("Single instance already created.")
	}

    return taskManagerInstance
}

func (tm *TaskManager) setBuilder(b IBuilder) {
	tm.builder = b
}

func (tm *TaskManager) buildTask(title, description string) {
	tm.builder.setTitle(title)
	tm.builder.setDescription(description)
	
	if _, ok := tm.builder.(*DomesticBuilder); ok {
		taskManagerInstance.addDomesticTask(tm.builder.getTask())
	} else {
		taskManagerInstance.addWorkTask(tm.builder.getTask())
	}
}

func (tm *TaskManager) addDomesticTask (t Task) {
	tm.domesticTasks[t] = &Memento{}
}

func (tm *TaskManager) removeDomesticTask (t Task) {
	delete(tm.domesticTasks, t)
}

func (tm *TaskManager) addWorkTask (t Task) {
	tm.workTasks[t] = &Memento{}
}

func (tm *TaskManager) removeWorkTask (t Task) {
	delete(tm.workTasks, t)
}
