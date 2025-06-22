package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	const (
		Create = "create"
		Read   = "read"
		Update = "update"
		Delete = "delete"
	)

	tasks := make([]string, 0)

	fmt.Println("Welcome to the TO DO List CLI app!")

	for {
		fmt.Println()
		fmt.Println("Enter your command (create, read, update, delete):")
		var command string
		fmt.Scan(&command)

		if command == "exit" {
			fmt.Println("Game Over")
			break
		}

		switch command {
		case Create:
			fmt.Println("Enter task name:")
			newTask, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			newTask = newTask[:len(newTask)-1]

			tasks = append(tasks, newTask)

		case Read:
			// empty list
			if len(tasks) < 1 {
				fmt.Println("List is empty! Please create a task.")
			}
			// show all tasks
			for index, task := range tasks {
				fmt.Printf("%d: %s\n", index+1, task)
			}

		case Update:
			fmt.Println("Type a task would you like to change?:")
			input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			input = input[:len(input)-1]

			// save index of old task in variable
			oldTaskIndex := -1
			for i, task := range tasks {
				if task == input {
					oldTaskIndex = i
					break
				}
			}

			if oldTaskIndex == -1 {
				fmt.Println("Invalid name. Please try again.")
				continue
			}
			fmt.Println("Type a new name:")
			newTaskName, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			newTaskName = newTaskName[:len(newTaskName)-1]

			// check the new task more 3 letters
			if len(newTaskName) < 3 {
				fmt.Println("The new task name is too short! Please, try again.")
				continue
			}

			tasks[oldTaskIndex] = newTaskName
			fmt.Printf("Updated task with name \"%s\" successfully!\n", newTaskName)

		case Delete:
			fmt.Println("What is task would you like to remove?:")
			input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			input = input[:len(input)-1]

			indexToRemove := -1
			for i, task := range tasks {
				if task == input {
					indexToRemove = i
					break
				}
			}

			if indexToRemove == -1 {
				fmt.Println("You have no task to remove!")
				continue
			}

			// the best way use different structures, map for example, but the test project. The problem is slow to remove first or index in the middle from slices.
			oldTaskName := tasks[indexToRemove]
			tasks = append(tasks[:indexToRemove], tasks[indexToRemove+1:]...) // remove elements before and after our task element
			fmt.Printf("Removed task #%d with name \"%s\" successfully!\n", indexToRemove, oldTaskName)
		default:
			fmt.Println("Invalid command! Please, try again")

		}

	}
}
