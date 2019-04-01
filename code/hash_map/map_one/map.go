package map_one

import "sync"

type RWmap struct {
	ma map[string]interface{}
	rw sync.RWMutex
}

func Init() *RWmap {
	return &RWmap{}
}

func (m *RWmap) Add()  {

}
