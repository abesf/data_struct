package base

import "fmt"

//base struct
type ListNode struct {
	Val  int
	Next *ListNode
}

//foreach
func ShowNode(node *ListNode)  {
	for node!=nil{
		fmt.Printf("%+v\n",*node)
		node=node.Next
	}
}

//insert from head
func Add() {
	var head=new(ListNode)
	head.Val=0
	var tail *ListNode
	tail=head
	for i:=1;i<10;i++{
		var node  =ListNode{Val:i}
		node.Next=tail
		tail=&node
	}
	ShowNode(tail)
}
//insert from tail
func Append() {
	var head=new(ListNode)
	head.Val=0
	var tail *ListNode
	tail=head
	for i:=1;i<10;i++{
		var node  =ListNode{Val:i}
		(*tail).Next=&node
		tail=&node
	}
	ShowNode(head)
}
