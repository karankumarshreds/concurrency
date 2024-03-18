package main

import (
	"fmt"
	"time"

	_ "github.com/karankumarshreds/concurrency/deadlock"
	"github.com/karankumarshreds/concurrency/mutex"
)

func main() {
    mutex.Main()
    // deadlock.Main() this code will not run (on purpose)
    for range time.Tick(1 *time.Second) {
        fmt.Println("tick")
    }
}
