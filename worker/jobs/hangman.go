package jobs

import (
	"math/rand"
	"time"

	"github.com/6ar8nas/learning-go/worker/utils"
)

type Tuple[T any, Y any] struct {
	First  T
	Second Y
}

func GuessWord(target string, names ...string) string {
	terminate := make(chan bool)
	defer close(terminate)

	generators := make([]<-chan *Tuple[string, string], len(names))
	for i, name := range names {
		generators[i] = utils.Generator(func() *Tuple[string, string] {
			return &Tuple[string, string]{First: generateRandomString(len(target)), Second: name}
		}, terminate)
	}

	guessCh := utils.Multiplex(generators...)
	timeout := time.After(5 * time.Minute)
	for i := 0; ; i++ {
		select {
		case guess := <-guessCh:
			if guess.First == target {
				return guess.Second
			}
		case <-timeout:
			return ""
		}
	}
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
