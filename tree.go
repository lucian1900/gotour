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
    vals := t1.Value == t2.Value

    if vals == false {
        return false
    }

    lefts := Same(t1.Left, t2.Left)
    rights := Same(t1.Right, t2.Right)

    return t1 == t2 && lefts && rights
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