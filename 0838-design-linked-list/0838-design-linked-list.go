package main

import "fmt"

type Node struct {
    Val  int
    Next *Node
}

type MyLinkedList struct {
    Head *Node
    Tail *Node
    Size int
}

func Constructor() MyLinkedList {
    return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
    if index < 0 || index >= this.Size {
        return -1
    }
    cur := this.Head
    for i := 0; i < index; i++ {
        cur = cur.Next
    }
    return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
    newNode := &Node{Val: val, Next: this.Head}
    this.Head = newNode
    if this.Tail == nil {
        this.Tail = newNode
    }
    this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
    newNode := &Node{Val: val}
    if this.Tail == nil {
        this.Head = newNode
        this.Tail = newNode
    } else {
        this.Tail.Next = newNode
        this.Tail = newNode
    }
    this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
    if index < 0 || index > this.Size {
        return
    }
    if index == 0 {
        this.AddAtHead(val)
        return
    }
    if index == this.Size {
        this.AddAtTail(val)
        return
    }
    newNode := &Node{Val: val}
    cur := this.Head
    for i := 0; i < index-1; i++ {
        cur = cur.Next
    }
    newNode.Next = cur.Next
    cur.Next = newNode
    this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
    if index < 0 || index >= this.Size {
        return
    }
    if index == 0 {
        this.Head = this.Head.Next
        if this.Head == nil {
            this.Tail = nil
        }
        this.Size--
        return
    }
    cur := this.Head
    for i := 0; i < index-1; i++ {
        cur = cur.Next
    }
    cur.Next = cur.Next.Next
    if cur.Next == nil {
        this.Tail = cur
    }
    this.Size--
}
