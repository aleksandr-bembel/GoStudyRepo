package main

import "fmt"

func sum(slice []float64) (ret float64) {
    ret = 0.0
    for _,x := range slice {
        ret +=x
    }
    return
}

func main() {
    x :=[]float64{1,2,3,4,5,6,7,8,9,10,11,}

    fmt.Println(sum(x))
}
