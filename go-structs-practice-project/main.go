package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/go-structs-practice-project/note"
)

func main() {
	title, content := getNoteData()

	n, err := note.New(title,content)

	if err != nil {
		fmt.Println(err)
		return
	}

	n.Display()
	err = n.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}
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
