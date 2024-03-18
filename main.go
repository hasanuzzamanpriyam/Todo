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

var (
	todos []Todo
)

func main() {
	// Read existing todo list from a file (optional)
	readTodosFromFile("todos.txt")

	// User interaction loop
	for {
		// Display current todo list
		displayTodos()

		// Get user input
		action, text := getUserInput()

		switch action {
		case "add":
			addTodo(text)
		case "done":
			completeTodoByText(text)
		case "delete":
			deleteTodoByText(text)
		case "quit":
			writeTodosToFile("todos.txt")
			return
		default:
			fmt.Println("Invalid command. Please use 'add', 'done', 'delete', or 'quit'.")
		}

		// Save updated todo list to file (optional)
		writeTodosToFile("todos.txt")
	}
}

// Read todos item from text file
func readTodosFromFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	for _, line := range strings.Split(string(data), "\n") {
		parts := strings.Split(line, "|")
		if len(parts) != 3 {
			continue // Skip invalid lines
		}
		id, _ := strconv.Atoi(parts[0])
		done, _ := strconv.ParseBool(parts[1])
		todos = append(todos, Todo{Text: parts[2], Done: done, ID: id})
	}
}

// Writes todo item to text file
func writeTodosToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for _, todo := range todos {
		doneStr := "false"
		if todo.Done {
			doneStr = "true"
		}
		_, err := fmt.Fprintf(file, "%d|%s|%s\n", todo.ID, doneStr, todo.Text)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

// Display todo item in a formatted way
func displayTodos() {
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
	if action != "quit" {
		fmt.Println("Enter todo text: ")
		fmt.Scanln(&text)
	}
	return action, text
}

// Add a new todo item
func addTodo(text string) {
	todos = append(todos, Todo{Text: text, Done: false, ID: generateID()})
}

// Complete a todo item by text
func completeTodoByText(text string) {
	for i, todo := range todos {
		if todo.Text == text {
			todos[i].Done = true
			return
		}
	}
	fmt.Printf("Todo item '%s' not found\n", text)
}

// Delete a todo item by text
func deleteTodoByText(text string) {
	for i, todo := range todos {
		if todo.Text == text {
			todos = append(todos[:i], todos[i+1:]...)
			return
		}
	}
	fmt.Printf("Todo item '%s' not found\n", text)
}

// Generate a unique ID for a new todo item
func generateID() int {
	maxID := 0
	for _, todo := range todos {
		if todo.ID > maxID {
			maxID = todo.ID
		}
	}
	return maxID + 1
}
