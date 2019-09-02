package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go worker(i, rand.Intn(10), &wg)
	}
	wg.Wait()
	fmt.Println("END")
}

func worker(id int, val int, wg *sync.WaitGroup) {
	fmt.Println(id, "-", val)
	for i := 0; i <= val; i++ {
		fmt.Println(id, "===========", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("ID:", id, ":OK")
	wg.Done()
}
