package main

import "fmt"

func main() {
    fmt.Print("Enter temperature in Farenheit: ")
    var tempInF float64
    fmt.Scanf("%f", &tempInF)
    tempInC := (tempInF - 32.0) * 5.0 / 9.0
    fmt.Println(tempInC)
}
