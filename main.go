package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	Text string
	Done bool
	ID   int
}

func main() {
	// Read existing todo list from a file (optional)
	todos := readTodosFromFile("todos.txt")

	// User interaction loop
	for {
		// Display current todo list
		displayTodos(todos)

		// Get user input
		action, text := getUserInput()

		switch action {
		case "add":
			todos = append(todos, Todo{Text: text, Done: false, ID: len(todos) + 1})
		case "done":
			id, err := getTodoID(text, todos)
			if err != nil {
				fmt.Println(err)
				continue
			}
			todos[id-1].Done = true
		case "delete":
			id, err := getTodoID(text, todos)
			if err != nil {
				fmt.Println(err)
				continue
			}
			todos = removeTodo(id, todos)
		case "quit":
			writeTodosToFile("todos.txt", todos)
			return
		default:
			fmt.Println("Invalid command. Please use 'add', 'done', 'delete', or 'quit'.")
		}

		// Save updated todo list to file (optional)
		writeTodosToFile("todos.txt", todos)
	}
}

// Read todos item from text file
func readTodosFromFile(filename string) []Todo {
	todos := []Todo{}
	data, err := os.ReadFile(filename)
	if err != nil {
		return todos // Empty list will be returned
	}
	for _, line := range strings.Split(string(data), "\n") {
		parts := strings.Split(line, "|")
		if len(parts) != 3 {
			continue // Ekip invalid line
		}
		id, _ := strconv.Atoi(parts[0])
		done, _ := strconv.ParseBool(parts[1])
		todos = append(todos, Todo{Text: parts[2], Done: done, ID: id})
	}
	return todos
}

// Writes todo item to text file
func writeTodosToFile(filename string, todos []Todo) {
	data := ""
	for _, todo := range todos {
		doneStr := "false"
		if todo.Done {
			doneStr = "true"
		}
		data += fmt.Sprintf("%d|%s|%s\n", todo.ID, doneStr, todo.Text)
	}
	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

// Display todo item in a formatted way
func displayTodos(todos []Todo) {
	fmt.Println("********** TODO LIST **********")
	for _, todo := range todos {
		doneStr := "[ ]"
		if todo.Done {
			doneStr = "[x]"
		}
		fmt.Printf("%d. %s %s\n", todo.ID, doneStr, todo.Text)
	}
}

// Get user input for action
func getUserInput() (string, string) {
	var action, text string
	fmt.Println("Enter action (add, done, delete, quit): ")
	fmt.Scanln(&action)
	if action == "add" || action == "delete" {
		fmt.Println("Enter todo text: ")
		fmt.Scanln(&text)

	}
	return action, text
}

// Find ID of a todo item
func getTodoID(text string, todos []Todo) (int, error) {
	for i, todo := range todos {
		if todo.Text == text {
			return i + 1, nil
		}
	}
	return 0, fmt.Errorf("Todo item '%s' not found", text)
}

// Remove a todo item
func removeTodo(id int, todos []Todo) []Todo {
	return append(todos[:id-1], todos[id:]...)
}
