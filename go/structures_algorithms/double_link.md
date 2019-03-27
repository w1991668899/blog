```
package double

import (
	"errors"
	"fmt"
)

type Node struct {
	Pre *Node
	Next *Node
	Data interface{}
}

type Link struct {
	Head *Node
	Tail *Node
	Len uint
}

func NewLink() *Link {
	return &Link{
		Head:nil,
		Tail:nil,
		Len:0,
	}
}

func (l *Link) AppendToHead(n *Node)(bool, error) {
	if n == nil{
		return false, errors.New("node is nil")
	}

	if l.Len == 0 {
		n.Next = nil
		n.Pre = nil
		l.Head = n
		l.Tail = n
		l.Len++
		return true, nil
	}

	oldHead := l.Head
	l.Head = n
	l.Head.Pre = nil
	l.Head.Next = oldHead
	l.Len++

	return true, nil
}

func (l *Link) AppendToTail(n *Node)(bool, error)  {
	if n == nil {
		return false, errors.New("node is nil")
	}

	if l.Len == 0 {
		n.Next = nil
		n.Pre = nil
		l.Head = n
		l.Tail = n
		l.Len++
		return true, nil
	}

	oldTail := l.Tail
	l.Tail = n
	oldTail.Next = l.Tail
	l.Tail.Pre = oldTail
	l.Tail.Next = nil
	l.Len++

	return true, nil
}

func (l *Link) AppendAfter(pre, new *Node)(bool, error)  {
	if pre == nil || new == nil {
		return false, errors.New("is nil")
	}

	oldNext := pre.Next
	pre.Next = new
	new.Pre = pre
	new.Next = oldNext
	l.Len++

	return true, nil
}

func (l *Link) AppendAt(index uint, n *Node)(bool, error)  {
	if index > l.Len || n == nil {
		return false, errors.New("node is nil or index is false")
	}

	if index == 0 {
		return l.AppendToHead(n)
	}

	pre := l.Head
	for i := 1; i < int(index); i++{
		pre = pre.Next
	}

	return l.AppendAfter(pre, n)
}

func (l *Link) DelByIndex(index uint)(bool, error) {
	if index >= l.Len {
		return false, errors.New("index is false")
	}

	if index == 0 {
		l.Head = l.Head.Next
		l.Head.Pre = nil
		if l.Len == 1 {
			l.Tail = l.Head
			l.Tail.Next  = nil
			l.Head.Next = nil
		}
		l.Len--
		return true, nil
	}

	if index == l.Len - 1 {
		l.Tail = l.Tail.Pre
		l.Tail.Next = nil
		return true, nil
	}

	n := l.Head
	for i := 1; i < int(index); i++{
		n = n.Next
	}

	n.Next = n.Next.Next
	n.Next.Pre = n
	l.Len--
	return true, nil
}

func (l *Link) GetNodeByIndex(index uint) *Node {
	if index >= l.Len {
		return nil
	}

	n := l.Head
	for i := 0; i < int(index); i++{
		n = n.Next
	}
	return nil
}

func (l *Link) Print()  {
	cur := l.Head
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.Data)
		cur = cur.Next
		if cur != nil {
			format += "->"
		}
	}

	fmt.Println(format)
	fmt.Printf("length: %d\n", l.Len)
}







```