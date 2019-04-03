package single_link

import "sync"

type SingleLink struct {
	Head *Node
	Tail *Node
	Count uint
	sync.RWMutex
}

type Node struct {
	Next *Node
	Data interface{}
}

func NewSingleLink() *SingleLink {
	return &SingleLink{
		Head:nil,
		Tail:nil,
		Count:0,
	}
}

