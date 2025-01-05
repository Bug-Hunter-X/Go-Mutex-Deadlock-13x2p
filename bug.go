```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var m sync.Mutex
        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        m.Lock()
                        fmt.Printf("Goroutine %d is running with i = %d\n", i, i)
                        m.Unlock()
                }(i)
        }
        wg.Wait()
}
```