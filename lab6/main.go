package main

import (
	"fmt"
	"sync"
)

type TurnstileCounter struct {
	mutex   sync.Mutex
	counter int
}

func counterWithSynchronizationPrimitive(getIt int) int {
	turnsCount := TurnstileCounter{}
	var goroutines sync.WaitGroup
	goroutines.Add(getIt)
	for num := 0; num < getIt; num++ {
		go func() {
			defer goroutines.Done()
			turnsCount.mutex.Lock()
			turnsCount.counter++
			turnsCount.mutex.Unlock()
		}()
	}
	goroutines.Wait()
	return turnsCount.counter
}

func counterWithoutSynchronizationPrimitive(getIt int) int {
	turnsCount := TurnstileCounter{}
	var goroutines sync.WaitGroup
	goroutines.Add(getIt)
	for num := 0; num < getIt; num++ {
		go func() {
			defer goroutines.Done()
			turnsCount.counter++
		}()
	}
	goroutines.Wait()
	return turnsCount.counter
}

func main() {
	fmt.Println(counterWithSynchronizationPrimitive(500), counterWithoutSynchronizationPrimitive(500))

}
