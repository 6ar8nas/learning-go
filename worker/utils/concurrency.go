package utils

import (
	"sync"
)

func First[T any](fns ...func() T) T {
	c := make(chan T)
	runFunc := func(i int) { c <- fns[i]() }
	for i := range fns {
		go runFunc(i)
	}
	return <-c
}

func Generator[T any](fn func() T, terminate <-chan bool) <-chan T {
	c := make(chan T)
	go func() {
		defer close(c)
		for {
			select {
			case <-terminate:
				return
			case c <- fn():
				continue
			}

		}
	}()
	return c
}

func Multiplex[T any](generators ...<-chan T) <-chan T {
	ch := make(chan T)
	var wg sync.WaitGroup
	wg.Add(len(generators))

	for _, g := range generators {
		go func(gen <-chan T) {
			defer wg.Done()
			for {
				ch <- <-gen
			}
		}(g)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}
