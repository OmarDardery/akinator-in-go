package main

import (
	"fmt"

	datatree "github.com/OmarDardery/akinator-in-go/data-tree"
)

func main() {
	tree := datatree.NewTree()
	tree.Initialize("data.json")
	defer tree.WriteFile()
	fmt.Println(`
	------------------------------------
	Hello and welcome to my Go Akinator!
	------------------------------------
	How it works:
	============
	- It uses a tree to traverse between subjects, using questions
	answered by the user.
	- It then writes the tree into json format to a file named data.json
	- It then reinitializes the tree using the same preorder format used
	to write the data to the file.
	------------------------------------
	Start? :`)

	var play bool

	fmt.Scanln(&play)

	for play {
		tree.Play()
		fmt.Print("Do you want to play again?: ")
		fmt.Scanln(&play)
	}
}
