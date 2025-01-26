package main

import (
	"log"

	"github.com/6ar8nas/learning-go/worker/jobs"
)

func main() {
	name := jobs.GuessWord("node", "John", "Mary", "Peter", "Chris")
	log.Println(name)

	numbers := jobs.MineNumbers(100000000, 10)
	log.Println(numbers)
}
