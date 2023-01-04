package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {

	var wg sync.WaitGroup //baiscally a semaphore in golang

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words))

	for i, word := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, word), &wg)
	}

	wg.Wait()

	wg.Add(1)
	printSomething("This is the second thing to be printed!", &wg)
}
