package main

import "fmt"

func main() {
    print(area, 2, 4)
    print(sum, 2, 4)
}

func print(f func(int, int) int, b, c int) {
    fmt.Println(f(b, c))
}

func area(b, c int) int {
    return b + c
}

func sum(a, b int) int {
    return a + b
}