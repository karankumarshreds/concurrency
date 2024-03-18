package deadlock

import (
	"sync"
	"time"
  "fmt"
)

type Value struct {
    mu sync.Mutex
    value int
}

func Main() {
    var wg sync.WaitGroup
    printSum := func(a, b *Value) {
        defer wg.Done()
        a.mu.Lock()
        defer a.mu.Unlock()

        time.Sleep(2*time.Second)

        b.mu.Lock()
        defer b.mu.Unlock()

        fmt.Printf("sum=%v\n", a.value + b.value)
    }

    var a, b Value
    wg.Add(2)
    go printSum(&a, &b)
    go printSum(&b, &a)
    wg.Wait()
}
