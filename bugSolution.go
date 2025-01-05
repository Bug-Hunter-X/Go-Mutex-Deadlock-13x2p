```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        ch := make(chan int, 10) 
        var wg sync.WaitGroup
        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        ch <- i
                        fmt.Printf("Goroutine %d is running with i = %d\n", i, <-ch)
                }(i)
        }
        wg.Wait()
}
```