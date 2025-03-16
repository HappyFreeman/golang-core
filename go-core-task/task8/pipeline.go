package main

// createPipeline создает конвейер: читает из inputChan, возводит в куб и пишет в outputChan
func createPipeline(inputChan <-chan uint8, outputChan chan<- float64) {
	go func() { // запускаем в отдельной горутине иначе будет deadlock
		for v := range inputChan {
			result := float64(v) * float64(v) * float64(v)
			outputChan <- result
		}
		close(outputChan)	
	}()
}