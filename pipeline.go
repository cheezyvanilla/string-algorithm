package main

import (
	"fmt"
	"sync"
	"time"
)

//close the shared channel(in this case is 'done' channel) is effectively broadcasting to all pipelines
//to stop working

func main()  {
	
	done := make(chan struct{}, 2)
	defer close(done)
	
	timeStart := time.Now()
	gen := numGen(done, 4,5)

	out1 := square(done, gen)
	out2 := square(done, gen)
	
	aggr := merge(done, out1, out2)
	// for n := range aggr {
	// 	fmt.Println(n)
	// }
	fmt.Println(<-aggr)

	done <- struct{}{}
	done <- struct{}{}
	
	execTime := time.Now().Sub(timeStart)
	fmt.Println(execTime)
}

func numGen(done <-chan struct{}, nums ...int) <-chan int{
	out := make(chan int)
	go func(){ 
		defer close(out) 
		for _, num := range nums{
			select{
			case out <- num:
			case <-done:
				return
			}
		}
	}()

	return out
} 

func square(done <-chan struct{}, num <-chan int) <-chan int{
	out := make(chan int)

	go func(){
		for v := range num{
			select{
			case out <- v * v:
			case <-done:
				return
			}	
		}
		close(out)
	}()

	return out
}

func merge(done <-chan struct{}, chans ...<-chan int) <-chan int{
	var wg sync.WaitGroup
	out := make(chan int, 1)

	output := func(n <-chan int) {
		defer wg.Done()
		for num := range n{
			select {
			case out <- num:
			case <-done:
				return
			}
		}
	}

	wg.Add(len(chans))
	
	for _, c := range chans{
		go output(c)
	}

	go func(){
		wg.Wait()
		close(out)
	}()
	return out
}