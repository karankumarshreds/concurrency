package mutex

import (
	"fmt"
	"sync"
)

func ExampleA() {
    var memoryAccess sync.Mutex
    var value int
    go func() {
        memoryAccess.Lock()
        value++
        memoryAccess.Unlock()
    }()

    memoryAccess.Lock()
    if value == 0 {
        fmt.Println("Value unchanged", value)
    } else {
        fmt.Println("Value changed to:", value)
    }
}
