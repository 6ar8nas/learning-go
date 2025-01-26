package jobs

import (
	"math/rand"
	"time"

	"github.com/6ar8nas/learning-go/worker/utils"
)

func MineNumbers(divisor, count int) []int {
	res := make([]int, 0, count)
	c := make(chan int)
	fn := func() int { return mine(divisor) }
	for i := 0; i < count; i++ {
		go func() { c <- utils.First(fn, fn, fn, fn, fn) }()
	}
	timeout := time.After(5 * time.Minute)
	for i := 0; i < count; i++ {
		select {
		case num := <-c:
			res = append(res, num)
		case <-timeout:
			return res
		}
	}
	return res
}

func mine(divisor int) int {
	for {
		num := rand.Int()
		if num%divisor == 0 {
			return num
		}
	}
}
