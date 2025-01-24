package utils

import "sync"

type AutoIncrement struct {
	sync.Mutex
	id int
}

func (ai *AutoIncrement) Next() (id int) {
	ai.Lock()
	defer ai.Unlock()

	ai.id++
	id = ai.id
	return
}
