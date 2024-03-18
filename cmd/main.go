package main

import (
    "github.com/karankumarshreds/concurrency/mutex"
    "github.com/karankumarshreds/concurrency/deadlock"
)

func main() {
    mutex.Main()
    deadlock.Main()
}
