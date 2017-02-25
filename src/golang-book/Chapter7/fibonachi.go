package main

import "fmt"

func fib(n uint64) uint64 {
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }

    return fib(n-1) + fib(n-2)
}


func main() {
    var i uint64
    for i = 0; i< 20; i++ {
        fmt.Println(fib(i))
    }
}
