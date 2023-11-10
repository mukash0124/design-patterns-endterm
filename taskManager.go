package main

import "fmt"

type Pair struct {
	task *Task
	backup Memento
}

type taskManager struct {
	domesticTasks []Pair
	workTasks []Pair
	builder IBuilder
}

var taskManagerInstance *taskManager

func GetInstance() *taskManager {
	if taskManagerInstance == nil {
            taskManagerInstance = &taskManager{}
			taskManagerInstance.domesticTasks = make([]Pair, 0)
			taskManagerInstance.workTasks = make([]Pair, 0)
    } else {
            fmt.Println("Single instance already created.")
	}

    return taskManagerInstance
}

func (tm *taskManager) setBuilder(b IBuilder) {
	tm.builder = b
}

func (tm *taskManager) buildTask(title, description string) {
	tm.builder.setTitle(title)
	tm.builder.setDescription(description)
	if _, ok := tm.builder.(*DomesticBuilder); ok {
		tm.builder.setID(len(tm.domesticTasks))
		taskManagerInstance.addDomesticTask(tm.builder.getTask())
	} else {
		tm.builder.setID(len(tm.workTasks))
		taskManagerInstance.addWorkTask(tm.builder.getTask())
	}
}

func (tm *taskManager) addDomesticTask (t *Task) {
	tm.domesticTasks = append(tm.domesticTasks, Pair{t, Memento{}})
}

func (tm *taskManager) removeDomesticTask (id int) {
	tm.domesticTasks = append(tm.domesticTasks[:id], tm.domesticTasks[id + 1:]...)
}

func (tm *taskManager) addWorkTask (t *Task) {
	tm.workTasks = append(tm.workTasks, Pair{t, Memento{}})
}

func (tm *taskManager) removeWorkTask (id int) {
	tm.workTasks = append(tm.workTasks[:id], tm.workTasks[id + 1:]...)
}

func (tm *taskManager) checkDomesticRemovability () bool {
	return len(tm.domesticTasks) != 0
}

func (tm *taskManager) checkWorkRemovability () bool {
	return len(tm.workTasks) != 0
}

func PrintTasks (tasks []Pair) {
	for _, task := range tasks {
		fmt.Printf("%s\n", task.task.stringifyTask())
	}
}
