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
func Same(t1, t2 *tree.Tree) bool

func main() {
    t := tree.New(1)
    ch := make(chan int)
    go Walk(t, ch)

    for i := 1; i < 10; i++ {
        fmt.Printf("%d ", <- ch)
    }
}