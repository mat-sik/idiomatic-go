package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

var buildMapFunc func() map[int]float64 = func() map[int]float64 {
    return buildMap(100_000, 1)
}

var getCachedMap func() map[int]float64 = sync.OnceValue(buildMapFunc);

func main() {
    start := time.Now()
    m := getCachedMap()
    m = getCachedMap()
    m = getCachedMap()
    m = getCachedMap()
    m = getCachedMap()
    m = getCachedMap()
    duration := time.Since(start)

	fmt.Printf("Execution time: %v\n", duration)

	fmt.Printf("map: %v", m)

    for i := 0; i < 100_000; i += 1000 {
        fmt.Println(i, m[i])
    }
}

func ex1() {
	expectedDataAmount := 20
	dataCh := make(chan int, expectedDataAmount)

	var producerWG sync.WaitGroup
	producerWG.Add(2)
	go ex1Producer(&producerWG, dataCh, 0)
	go ex1Producer(&producerWG, dataCh, 1)

	var consumerWg sync.WaitGroup
	consumerWg.Add(1)
	go ex1Consumer(&consumerWg, dataCh, 2)

	producerWG.Wait()
	close(dataCh)
	consumerWg.Wait()
}

func ex1Producer(wg *sync.WaitGroup, dataCh chan<- int, id int) {
	defer func() {
		wg.Done()
	}()

	for i := 0; i < 10; i++ {
		dataCh <- id
	}
}

func ex1Consumer(wg *sync.WaitGroup, dataCh <-chan int, id int) {
	defer func() {
		wg.Done()
	}()

	for i := range dataCh {
		fmt.Printf("Consumer: %d, has consumed: %d\n", id, i)
	}
}

func ex2() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	work := func(ch chan<- int) {
		defer func() {
			close(ch)
		}()

		for i := 0; i < 10; i++ {
			ch <- i
		}
	}

	go work(ch1)
	go work(ch2)

	count := 2
	for count != 0 {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				count--
				break
			}
			fmt.Println(v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				count--
				break
			}
			fmt.Println(v)
		}
	}
}

func buildMap(end int, workers int) map[int]float64 {
	s := make([]float64, end)
	var wg sync.WaitGroup
	wg.Add(workers)

	work := func(start, end int) {
        defer wg.Done()

		for i := start; i < end; i++ {
			s[i] = math.Sqrt(float64(i))
		}
	}

	workPerWorker := end / workers
	for i := 0; i < workers; i++ {
		startIdx := i * workPerWorker
		endIdx := startIdx + workPerWorker
		if i == workers-1 {
			endIdx = end
		}
		go work(startIdx, endIdx)
	}

	wg.Wait()

    m := make(map[int]float64, len(s))
    for i, v := range s {
        m[i] = v
    }

	return m
}
