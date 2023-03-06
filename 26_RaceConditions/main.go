package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello Race Condition")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(4)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Adding 1")

		mut.Lock()
		score = append(score, 1)
		mut.Unlock()

		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Adding 2")

		mut.Lock()
		score = append(score, 2)
		mut.Unlock()

		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Adding 3")

		mut.Lock()
		score = append(score, 3)
		mut.Unlock()

		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {

		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()

		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
