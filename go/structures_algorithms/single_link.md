```
package single

import (
	"errors"
	"fmt"
)

type Node struct {
	Next *Node
	Data interface{}
}

type SingleLink struct {
	Head *Node
	Tail *Node
	Len uint
}

func NewSingleLink() *SingleLink  {
	return &SingleLink{
		Head:nil,
		Tail:nil,
		Len:0,
	}
}

// 在头部插入元素
func (l *SingleLink) AppendToHead(n *Node) (bool, error) {
	if n == nil {
		return false, errors.New("node is false")
	}

	oldHead := l.Head
	l.Head = n
	l.Head.Next = oldHead

	if l.Len == 0 {
		l.Tail = n
		l.Tail.Next = nil
	}

	l.Len++
	return true, nil
}

func (l *SingleLink) AppendToTail(n *Node) (bool, error) {
	if n == nil {
		return false, errors.New("node is nil")
	}

	if l.Len == 0 {
		res, err := l.AppendToHead(n)
		return res, err
	}

	oldTail := l.Tail
	l.Tail = n
	l.Tail.Next = nil
	oldTail.Next = n

	l.Len++
	return true, nil
}

// 在one节点后面插入新节点
func (l *SingleLink) AppendAfter(pre, new *Node) (bool, error) {
	if new == nil {
		return false, errors.New("is nil")
	}

	oldNext := pre.Next
	pre.Next = new
	new.Next = oldNext
	l.Len++
	return true, nil
}

// 将节点插入固定位置m, Head节点看做位置0
func (l *SingleLink) InsertAt (index uint, n *Node) (bool, error){
	if index > l.Len || n == nil  {
		return false, errors.New("len or node is false")
	}

	if index == 0{
		res, err := l.AppendToHead(n)
		return res, err
	}

	node := l.Head
	for i := 1; i < int(index); i++{
		node = node.Next
	}

	res, err := l.AppendAfter(node, n)
	return res, err
}

// 根据索引获取固定节点
func (l *SingleLink) GetByIndex(index uint)*Node{
	node := l.Head
	for i := 0; i< int(index); i++{
		node = node.Next
	}

	return node
}

// 删除索引上的节点
func (l *SingleLink) DelByIndex(index uint)(bool, error)  {
	if l.Len <= index || l.Len == 0{
		return false, errors.New("index is fail")
	}

	if index == 0 {
		l.Head = l.Head.Next
		l.Len--
		return true, nil
	}

	node := l.Head
	if index == l.Len - 1 {
		for i := 2; i < int(l.Len); i++{
			node = node.Next
		}

		node.Next = nil
		l.Tail = node
		l.Len--
		return true, nil
	}

	// 前一个节点
	preNode := l.Head
	for i := 1; i < int(index); i++{
		preNode = preNode.Next
	}

	// 要删除的节点
	for i := 0; i < int(index); i++{
		node = node.Next
	}

	preNode.Next = node.Next

	l.Len--
	return true, nil
}

// 打印链表
func (l *SingleLink) Print() {
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