```
package stack

import (
	"errors"
	"fmt"
)

type Node struct {
	Next *Node
	Data interface{}
}

type Stack struct {
	Head *Node
	Len uint
}

func NewStack()*Stack  {
	return &Stack{
		Head:nil,
		Len:0,
	}
}

func (s *Stack) Push(n *Node) (bool, error) {
	if n == nil {
		return false, errors.New("node is nil")
	}

	n.Next = s.Head
	s.Head = n
	s.Len++

	return true, nil
}

func (s *Stack) Pop() *Node {
	if s.Len == 0 {
		return nil
	}
	n := s.Head
	s.Head = s.Head.Next
	s.Len--
	return n
}

func (s *Stack) Print()  {
	cur := s.Head
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.Data)
		cur = cur.Next
		if cur != nil {
			format += "->"
		}
	}

	fmt.Println(format)
	fmt.Printf("length: %d\n", s.Len)
}

```