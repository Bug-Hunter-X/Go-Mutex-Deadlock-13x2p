# Go Mutex Deadlock

This repository demonstrates a potential deadlock scenario in Go using mutexes. The `bug.go` file contains code that can deadlock under certain conditions. The `bugSolution.go` file provides a corrected version of the code. 

## Description
The original code creates multiple goroutines that try to acquire a mutex simultaneously. This can lead to a deadlock if the mutex is already held by another goroutine that is waiting for the mutex to be released, and vice versa. 

## Solution
The solution involves using a more appropriate synchronization mechanism, such as channels, or carefully considering the order in which the mutex is acquired and released. 

## How to Run
1. Clone this repository.
2. Navigate to the repository directory.
3. Run `go run bug.go` to see the potential deadlock.
4. Run `go run bugSolution.go` to see the corrected version.
