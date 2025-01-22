package util

import "sync"

type AutoIncrement struct {
	sync.Mutex
	id int
}

func (ai *AutoIncrement) Id() (id int) {
	ai.Lock()
	defer ai.Unlock()

	ai.id++
	id = ai.id
	return
}
