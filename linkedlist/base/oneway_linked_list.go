package base

import "fmt"

//base struct
type OneWayNode struct {
	Val  int
	Next *OneWayNode
}

//Traversal from start from head
func OneWayTraversal(node *OneWayNode)  {
	for node!=nil{
		fmt.Printf("%+v\n",*node)
		node=node.Next
	}
}
//todo ReverseTraversal  start from end
func OneWayReverseTraversal(node *OneWayNode)  {
	for node!=nil{
		fmt.Printf("%+v\n",*node)
		node=node.Next
	}
}

//insert from head
func OneWayAdd() {
	var head=new(OneWayNode)
	head.Val=0
	var tail *OneWayNode
	tail=head
	for i:=1;i<10;i++{
		var node  =OneWayNode{Val:i}
		node.Next=tail
		tail=&node
	}
	OneWayTraversal(tail)
}
//insert from tail
func OneWayAppend() {
	var head=new(OneWayNode)
	head.Val=0
	var tail *OneWayNode
	tail=head
	for i:=1;i<10;i++{
		var node  =OneWayNode{Val:i}
		(*tail).Next=&node
		tail=&node
	}
	OneWayTraversal(head)
}
//todo remove one node
func OneWayRemove()  {
	
}