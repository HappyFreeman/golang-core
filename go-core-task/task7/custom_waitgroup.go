package main

type CustomWaitGroup struct {
	semaphore chan struct{}
}

func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{
		semaphore: make(chan struct{}, 1),
	}
}

func (wg *CustomWaitGroup) Add(delta int) {

	for i := 0; i < delta; i++ {
		wg.semaphore <- struct{}{} // Положить в канал delta горутин
	}
}

func (wg *CustomWaitGroup) Done() {
	<-wg.semaphore // Принять из канала одну горутину
}

// Wait блокирует выполнение до завершения всех горутин
func (wg *CustomWaitGroup) Wait() { // help
	for i := 0; i < cap(wg.semaphore); i++ {
		<-wg.semaphore
	}
}