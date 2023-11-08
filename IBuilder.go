package main

type IBuilder interface {
	setTitle(title string)
	setDescription(description string)
	getTask() Task
}

func GetBuilder(taskType string) IBuilder {
	if taskType == "domestic" {
		return newDomesticBuilder()
	}

	if taskType == "work" {
		return newWorkBuilder()
	}

	return nil
}