package datatree

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func NewStr() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

type DataTree struct {
	Root   *Node
	Length int
}

func NewTree() DataTree {
	return DataTree{
		Root:   nil,
		Length: 0,
	}
}

func (t *DataTree) Initialize(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("No previous data!")
		fmt.Print("Enter a subject to begin: ")
		t.Root = &Node{
			ID:         1,
			Data:       NewStr(),
			Left:       nil,
			Right:      nil,
			IsQuestion: false,
		}
		t.Length = 1
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&t.Root)
	t.Length = count(t.Root)
	fmt.Println(t.Length)
	if err != nil {
		panic("invalid json file")
	}

}

func traverse(c *Node, length *int) {
	if c.IsQuestion {
		fmt.Print(c.Data)
		var choice bool
		fmt.Scanln(&choice)
		if choice {
			traverse(c.GoRight(), length)
		} else {
			traverse(c.GoLeft(), length)
		}
	} else {
		fmt.Printf("Is the subject on your mind '%v'? ", c.Data)
		var choice bool
		fmt.Scanln(&choice)
		if choice {
			fmt.Println("great!")
		} else {
			c.Learn(length)
		}
	}
}

func (t *DataTree) Play() {
	traverse(t.Root, &t.Length)
}

func (t *DataTree) WriteFile() {
	file, err := os.Create("data.json")
	if err != nil {
		panic("NOOOO")
	}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(t.Root); err != nil {
		panic("failed to encode tree")
	}

}

func count(c *Node) int {
	if c != nil {
		return 1 + count(c.Left) + count(c.Right)
	} else {
		return 0
	}
}
