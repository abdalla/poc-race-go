package main

import (
	"fmt"
	"runtime"
	"sync"
)

const num = 1000

var wg sync.WaitGroup

func code1() {
	c := 0
	//runtime.GOMAXPROCS(5)

	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			x := c
			runtime.Gosched()

			x++
			c = x
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("total:", c)
}

func code2() {
	c := 0

	wg.Add(num)

	var mu sync.Mutex

	for i := 0; i < num; i++ {
		go func() {
			mu.Lock()

			x := c
			runtime.Gosched()

			x++
			c = x

			mu.Unlock()

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("total:", c)

}

func main() {

	code1()

	code2()

}
