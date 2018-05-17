package main

import (
    "fmt")

// Use closure to implement a tree
// Set and get tree levels 


type Node string 

func NodeFunc(newScope Node, thisScope Node) (func () Node) {
    return func () Node {
        return thisScope + newScope
    }
}

func main() {
    root := NodeFunc("root", "")
    x := root()
    fmt.Printf("%v", x)
    leftChild := NodeFunc(root, "left")
    fmt.Printf("leftChild %v", leftChild)
}