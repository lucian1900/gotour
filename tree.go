package main

import (
    "code.google.com/p/go-tour/tree"
    "fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    ch <- t.Value

    if t.Left != nil {
        go Walk(t.Left, ch)
    }

    if t.Right != nil {
        go Walk(t.Right, ch)
    }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go Walk(t1, ch1)
    go Walk(t2, ch2)

    for {
        i1, ok1 := <- ch1
        i2, ok2 := <- ch2

        if ok1 != ok2 || (ok1 == true && ok2 == true && i1 != i2) {
            return false
        }
    }

    return true
}

func main() {
    ch := make(chan int)
    go Walk(tree.New(1), ch)

    for i := 1; i < 10; i++ {
        fmt.Printf("%d ", <- ch)
    }
    fmt.Println("")

    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(1), tree.New(2)))
}