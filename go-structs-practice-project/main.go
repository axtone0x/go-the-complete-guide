package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"example.com/go-structs-practice-project/note"
	"example.com/go-structs-practice-project/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	Save() error
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	n, err := note.New(title,content)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = saveData(todo)

	if err != nil {
		fmt.Println("Saving the todo failed.")
		return
	}

	fmt.Println("Saving the todo succeeded!")

	n.Display()
	err = saveData(n)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Saving the note succeeded!")
}

func outputData(data outputtable) {
	data.Display()
	err := saveData(data)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Save succeeded!")
}

func saveData (data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving failed.")
		return err
	}

	fmt.Println("Saving succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ",prompt)
	
	reader := bufio.NewReader(os.Stdin) // --> Creating a reader that listens to the command line

	text, err := reader.ReadString('\n')

	text = strings.TrimSuffix(text,"\n")
	text = strings.TrimSuffix(text,"\r") // --> For Windows since it uses \n\r

	if err != nil {
		return ""
	}

	return text
}
