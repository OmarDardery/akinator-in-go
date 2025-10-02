package datatree

import "fmt"

type Node struct {
	ID         int    `json:"id"`
	Data       string `json:"data"`
	Left       *Node  `json:"right,omitempty"`
	Right      *Node  `json:"left,omitempty"`
	IsQuestion bool   `json:"is_question"`
}

func NewNode(data string, id int) Node {
	return Node{
		ID:         id,
		Data:       data,
		Left:       nil,
		Right:      nil,
		IsQuestion: false,
	}
}

func (n *Node) Learn(length *int) {
	fmt.Print("whats the subject on your mind?")
	rightSubject := Node{
		ID:         (*length) + 1,
		Data:       NewStr(),
		Left:       nil,
		Right:      nil,
		IsQuestion: false,
	}
	leftSubject := Node{
		ID:         (*length) + 2,
		Data:       n.Data,
		Left:       nil,
		Right:      nil,
		IsQuestion: false,
	}
	*length = *length + 2
	n.Left = &leftSubject
	n.Right = &rightSubject
	fmt.Print("whats a yes or no question that diffeentiates between your subject and my wrong guess? format it so that it's answer is yes for your subject and no for mine.")
	n.Data = NewStr()
	n.IsQuestion = true
}

func (n Node) GoLeft() *Node {
	return n.Left
}

func (n Node) GoRight() *Node {
	return n.Right
}
