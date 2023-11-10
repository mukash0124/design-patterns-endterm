package main

import "fmt"

func main() {
	taskManager := GetInstance()

	domesticBuilder := GetBuilder("domestic")
	workBuilder := GetBuilder("work")

	fmt.Println("Hello, dear user! We are glad to see you in our task managing application!")

	for {
		fmt.Println("What would you like to do?")
		fmt.Printf("1. Show tasks.\n2. Add task.\n3. Delete task\n")
		var operation int
		fmt.Scan(&operation)
		switch operation {
		case 1:
			fmt.Println("What kind of tasks you want to see?")
			fmt.Printf("1. Domestic tasks.\n2. Work tasks.\n")
			var kind int
			fmt.Scan(&kind)
			switch kind {
			case 1:
				if taskManager.checkDomesticRemovability() {
					PrintTasks(taskManager.domesticTasks)
					fmt.Println("What do you want to do next?")
					fmt.Printf("1. Edit task.\n2. Undo last edit.\n3. Start task.\n4. Finish task.\n5. Back.\n")
					fmt.Scan(&operation)
					switch operation {
					case 1:
						PrintTasks(taskManager.domesticTasks)
						var id int
						var title, description string
						fmt.Println("Enter id of task you want to edit:")
						fmt.Scan(&id)
						fmt.Println("Enter new title of the task:")
						fmt.Scan(&title)
						fmt.Println("Enter new description of the task:")
						fmt.Scan(&description)
						id--
						taskManager.domesticTasks[id].backup = taskManager.domesticTasks[id].task.createMemento()
						err := taskManager.domesticTasks[id].task.currentState.edit(title, description)
						if err == nil {
							fmt.Println("Task has successfully been edited!")
						} else {
							fmt.Println(err)
						}
					case 2:
						PrintTasks(taskManager.domesticTasks)
						var id int
						fmt.Println("Enter id of task you want to restore:")
						fmt.Scan(&id)
						id--
						if taskManager.domesticTasks[id].backup != (Memento{}) {
							taskManager.domesticTasks[id].task.restoreMemento(taskManager.domesticTasks[id].backup)
							fmt.Println(taskManager.domesticTasks[id].task.stringifyTask())
						} else {
							fmt.Println("You can't restore task that hasn't been edited!")
						}
					case 3:
						PrintTasks(taskManager.domesticTasks)
						var id int
						fmt.Println("Enter id of task you want to start:")
						fmt.Scan(&id)
						id--
						err := taskManager.domesticTasks[id].task.currentState.start()
						if err == nil {
							fmt.Println("Task has been successfully started!")
						} else {
							fmt.Println("Current state does not allow to use start command!")
						}
					case 4:
						PrintTasks(taskManager.domesticTasks)
						var id int
						fmt.Println("Enter id of task you want to finish:")
						fmt.Scan(&id)
						id--
						err := taskManager.domesticTasks[id].task.currentState.finish()
						if err == nil {
							fmt.Println("Task has been successfully finished!")
						} else {
							fmt.Println("Current state does not allow to use finish command!")
						}
					case 5:
						continue
					default:
						fmt.Println("Choose one of operations provided above!")

					}
				} else {
					fmt.Println("There are no domestic tasks")
				}

			case 2:
				if taskManager.checkWorkRemovability() {
					PrintTasks(taskManager.workTasks)
					fmt.Println("What do you want to do next?")
					fmt.Printf("1. Edit task.\n2. Undo last edit.\n3. Start task.\n4. Finish task.\n5. Back.\n")
					fmt.Scan(&operation)
					switch operation {
					case 1:
						PrintTasks(taskManager.workTasks)
						var id int
						var title, description string
						fmt.Println("Enter id of task you want to edit:")
						fmt.Scan(&id)
						fmt.Println("Enter new title of the task:")
						fmt.Scan(&title)
						fmt.Println("Enter new description of the task:")
						fmt.Scan(&description)
						id--
						taskManager.workTasks[id].backup = taskManager.workTasks[id].task.createMemento()
						err := taskManager.workTasks[id].task.currentState.edit(title, description)
						if err == nil {
							fmt.Println("Task has successfully been edited!")
						} else {
							fmt.Println(err)
						}
					case 2:
						PrintTasks(taskManager.workTasks)
						var id int
						fmt.Println("Enter id of task you want to restore:")
						fmt.Scan(&id)
						id--
						if taskManager.domesticTasks[id].backup != (Memento{}) {
							taskManager.domesticTasks[id].task.restoreMemento(taskManager.domesticTasks[id].backup)
							fmt.Println(taskManager.domesticTasks[id].task.stringifyTask())
						} else {
							fmt.Println("You can't restore task that hasn't been edited!")
						}
					case 3:
						PrintTasks(taskManager.workTasks)
						var id int
						fmt.Println("Enter id of task you want to start:")
						fmt.Scan(&id)
						id--
						err := taskManager.workTasks[id].task.currentState.start()
						if err == nil {
							fmt.Println("Task has been successfully started!")
						} else {
							fmt.Println("Current state does not allow to use start command!")
						}
					case 4:
						PrintTasks(taskManager.workTasks)
						var id int
						fmt.Println("Enter id of task you want to finish:")
						fmt.Scan(&id)
						id--
						err := taskManager.workTasks[id].task.currentState.finish()
						if err == nil {
							fmt.Println("Task has been successfully finished!")
						} else {
							fmt.Println("Current state does not allow to use finish command!")
						}
					case 5:
						continue
					default:
						fmt.Println("Choose one of operations provided above!")
					}
				} else {
					fmt.Println("There are no work tasks")
				}

			default:
				fmt.Println("Choose one of kinds provided above!")
			}
		case 2:
			fmt.Println("What kind of tasks you want to add?")
			fmt.Printf("1. Domestic tasks.\n2. Work tasks.\n")
			var kind int
			fmt.Scan(&kind)
			switch kind {
			case 1:
				var title, description string
				fmt.Println("Enter title of new task:")
				fmt.Scan(&title)
				fmt.Println("Enter description of new task:")
				fmt.Scan(&description)
				taskManager.setBuilder(domesticBuilder)
				taskManager.buildTask(title, description)
				fmt.Println("New task was created successfully!")
			case 2:
				var title, description string
				fmt.Println("Enter title of new task:")
				fmt.Scan(&title)
				fmt.Println("Enter description of new task:")
				fmt.Scan(&description)
				taskManager.setBuilder(workBuilder)
				taskManager.buildTask(title, description)
				fmt.Println("New task was created successfully!")
			default:
				fmt.Println("Choose one of kinds provided above!")
			}
		case 3:
			fmt.Println("What is the kind of task you want to delete?")
			fmt.Printf("1. Domestic tasks.\n2. Work tasks.\n")
			var kind int
			fmt.Scan(&kind)
			switch kind {
			case 1:
				PrintTasks(taskManager.domesticTasks)
				var id int
				fmt.Println("Enter id of task you want to delete:")
				fmt.Scan(&id)
				id--
				taskManager.removeDomesticTask(id)
			case 2:
				PrintTasks(taskManager.workTasks)
				var id int
				fmt.Println("Enter id of task you want to delete:")
				fmt.Scan(&id)
				id--
				taskManager.removeWorkTask(id)
			default:
				fmt.Println("Choose one of kinds provided above!")
			}
		default:
			fmt.Println("Choose one of operations provided above!")

		}

	}

}
