package main

import "fmt"

func swap(x *int64, y *int64) {
    temp := *x
    *x = *y
    *y = temp
}

func main() {
    x := int64(1)
    y := int64(2)
    swap(&x, &y)

    fmt.Println(x)
    fmt.Println(y)
}
