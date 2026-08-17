package counter

import(
	"sync"
	"testing"
)

func TestCounterConcurrentAdd(t *testing.T){
	const workers=100
	const addsPerWorker=100
	
	var c Counter
	var wg sync.WaitGroup
	wg.Add(workers)

	for range workers{
		go func(){
			defer wg.Done()
			for range addsPerWorker{
				c.Add()
			}
		}()
	}
	wg.Wait()
	if got, want := c.Value(), workers*addsPerWorker; got != want {
		t.Fatalf("hits = %d, want %d", got, want)
	}
}