package main

import "fmt"

func average(xs []float64) float64 {
    total := 0.0
    for _, v := range xs {
        total += v
    }
    return total / float64(len(xs))
}

func max(args ...float64) (max float64) {
    max = args[0]
    for _,x := range args {
        if x > max {
            max = x
        }
    }
    return
}

func half(n int) (int,bool) {
    halfn := int(n) / 2

    if halfn % 2 == 0 {
       return halfn, true
    } else {
       return halfn, false
    }
}

func main() {
    xs := []float64{98,93,77,82,83}

    fmt.Println(average(xs))

    x,ok := half(1)
    fmt.Println("Half from 1", x, ok)
    x,ok = half(2)
    fmt.Println("Half from 2", x, ok)
    fmt.Println("Max test: ", max(1,2,3,23.0,5,6,7,8.0,3,11))
}
