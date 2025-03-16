package main

import "sync"

// MergeChannels принимает N каналов и возвращает единый канал со всеми значениями.
func MergeChannels(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
    mergedChannel := make(chan int)

	wg.Add(len(channels))
	for _, c := range channels {
        go func(c <-chan int) {
			for n := range c {
				mergedChannel <- n
			}
			wg.Done()
		}(c)
    }

	go func() {
        wg.Wait()
        close(mergedChannel)
    }()
	
    return mergedChannel
}