package main

import "fmt"

func main() {
	tm := GetInstance()

	db := GetBuilder("domestic")

	tm.setBuilder(db)

	tm.buildTask("1", "dssss")

	for task, _ := range tm.domesticTasks {
		fmt.Printf("Title: %s, Description: %s", task.title, task.description)
	}
}