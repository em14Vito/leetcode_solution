
package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	 Next *ListNode
	
}

func main() {
	fmt.Printf("hello, world\n")

	// var l5 ListNode
	// l5.Val=0
	// l5.Next = nil

	// var l6 ListNode
	// l6.Val=2
	// l6.Next = nil

	// var l3 ListNode
	// l3.Val=6
	// l3.Next = &l5

	// var l4 ListNode
	// l4.Val=8
	// l4.Next = &l6

	var l3 ListNode
	l3.Val=1
	l3.Next = nil

	var l4 ListNode
	l4.Val=8
	l4.Next = nil


	var l1 ListNode
	l1.Val=5
	l1.Next = nil

	var l2 ListNode
	l2.Val=5
	l2.Next = nil

	addTwoNumbers(&l1,&l2)

}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
 
	var root ListNode
	root.Val = 0
	// out := cacluate(l1,l2,0,0)

	cacluateResult(l1,l2,&root,0)

	// fmt.Println(out)

	var temp = root
	for 1==1 {
		if(&temp != nil){
			fmt.Println(temp.Val)
			if(temp.Next != nil){
				temp = *temp.Next 
			} else {
				break
			}
		} else{
			break
		}
	}

	return &root
}

func cacluateResult(l1 ,l2 ,root  *ListNode, remainder int) {
	
	Var := remainder
	if(l1 != nil && l2 != nil){
		Var = Var + l1.Val + l2.Val + root.Val
	}else if(l1 != nil && l2 == nil){
		Var = Var + l1.Val + root.Val
	}else if(l1 == nil && l2 != nil){
		Var = Var + l2.Val + root.Val
	}

	//计算
	if(Var >= 10) {
		root.Val = Var - int(Var /10) * 10 
		remainder = Var / 10
	} else{
		root.Val =Var
		remainder = 0
	}
	

	var nextNode ListNode
	

	if(l1 != nil && l2 != nil){
		cacluateResult(l1.Next,l2.Next,&nextNode,remainder)
		if(nextNode.Val != -1){
			root.Next = &nextNode
		}
	}else if(l1 != nil && l2 == nil){
		cacluateResult(l1.Next,l2,&nextNode,remainder)
		if(nextNode.Val != -1){
			root.Next = &nextNode
		}
	}else if(l1 == nil && l2 != nil){
		cacluateResult(l1,l2.Next,&nextNode,remainder)
		if(nextNode.Val != -1){
			root.Next = &nextNode
		}
	} else if(root.Val == 0){
		root.Val = -1
		return
	}
}

